// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/db.proto

package rocksdb

import (
	"context"
	"fmt"
	"io"

	"github.com/go-redis/redis/v9"
	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/db/grpc/pb"
	"github.com/pingcap/go-ycsb/pkg/util"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// properties
const (
	grpcAddr = "grpc.addr"
)

type grpcCreator struct{}

type grpcDB struct {
	p *properties.Properties

	conn   *grpc.ClientConn
	client pb.DBClient

	r       *util.RowCodec
	bufPool *util.BufPool
}

func (c grpcCreator) Create(p *properties.Properties) (ycsb.DB, error) {
	addr := p.GetString(grpcAddr, "127.0.0.1:8000")

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewDBClient(conn)

	return &grpcDB{
		p:       p,
		conn:    conn,
		client:  client,
		r:       util.NewRowCodec(p),
		bufPool: util.NewBufPool(),
	}, nil
}

func (db *grpcDB) Close() error {
	db.conn.Close()
	return nil
}

func (db *grpcDB) InitThread(ctx context.Context, _ int, _ int) context.Context {
	return ctx
}

func (db *grpcDB) CleanupThread(_ context.Context) {
}

func (db *grpcDB) getRowKey(table string, key string) string {
	return fmt.Sprintf("%s:%s", table, key)
}

func (db *grpcDB) Read(ctx context.Context, table string, key string, fields []string) (map[string][]byte, error) {
	val, err := db.client.Read(ctx, &pb.Key{Key: db.getRowKey(table, key)})
	if err != nil {
		return nil, err
	}

	value := val.Val
	if value == nil {
		return nil, redis.Nil
	}

	return db.r.Decode(value, fields)
}

func (db *grpcDB) Scan(ctx context.Context, table string, startKey string, count int, fields []string) ([]map[string][]byte, error) {
	res := make([]map[string][]byte, count)

	rowStartKey := db.getRowKey(table, startKey)

	stream, err := db.client.Scan(ctx, &pb.KeyScan{StartKey: rowStartKey, Count: int32(count)})
	if err != nil {
		return nil, err
	}

	for i := 0; i < count; i++ {
		val, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		m, err := db.r.Decode(val.Val, fields)
		if err != nil {
			return nil, err
		}

		res[i] = m
	}

	return res, nil
}

func (db *grpcDB) Update(ctx context.Context, table string, key string, values map[string][]byte) error {
	m, err := db.Read(ctx, table, key, nil)
	if err != nil {
		return err
	}

	for field, value := range values {
		m[field] = value
	}

	buf := db.bufPool.Get()
	defer db.bufPool.Put(buf)

	buf, err = db.r.Encode(buf, m)
	if err != nil {
		return err
	}

	rowKey := db.getRowKey(table, key)

	_, err = db.client.Put(ctx, &pb.KeyVal{Key: rowKey, Val: buf})
	return err
}

func (db *grpcDB) Insert(ctx context.Context, table string, key string, values map[string][]byte) error {
	rowKey := db.getRowKey(table, key)

	buf := db.bufPool.Get()
	defer db.bufPool.Put(buf)

	buf, err := db.r.Encode(buf, values)
	if err != nil {
		return err
	}

	_, err = db.client.Put(ctx, &pb.KeyVal{Key: rowKey, Val: buf})
	return err
}

func (db *grpcDB) Delete(ctx context.Context, table string, key string) error {
	rowKey := db.getRowKey(table, key)

	_, err := db.client.Delete(ctx, &pb.Key{Key: rowKey})
	return err
}

func init() {
	ycsb.RegisterDBCreator("grpc", grpcCreator{})
}
