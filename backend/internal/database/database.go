package database

import "context"

type Database interface {
	Connect(ctx context.Context) error
	Query(ctx context.Context, query string, args ...interface{}) (interface{}, error)
	Execute(ctx context.Context, statement string, args ...interface{}) (int64, error)
	Close() error
}

