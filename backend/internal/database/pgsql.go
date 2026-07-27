package database

import (
	"context"
	"database/sql"
	"fmt"
	"magic-bullet/backend/internal/config"

	_ "github.com/lib/pq"
)

type PostgresConnector struct {
	config *config.Config
	db *sql.DB
	connectionString string
}

func NewPGSQLConnector(config *config.Config) (*PostgresConnector, error) {
	if config == nil {
		return nil, fmt.Errorf("config parameter connot be nil")
	}

	return &PostgresConnector{
		config: config,
		connectionString: fmt.Sprintf(
				"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName,
			),
	}, nil
}

func (pg *PostgresConnector) Connect(ctx context.Context) error {
	var err error

	pg.db, err = sql.Open("postgres", pg.connectionString)
	if err != nil {
		return fmt.Errorf("Failed to open connection: %s", err)
	}
	
	err = pg.db.PingContext(ctx) 
	if err != nil {
		return fmt.Errorf("Failed to ping database: %s", err)
	}

	return nil
}

func (pg *PostgresConnector) Query(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	if pg.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	rows, err := pg.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Error on execute query: %s", err)
	}
	defer rows.Close()

	var results []map[string]interface{}

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("Error on get columns: %s", err)
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuesPtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuesPtrs[i] = &values[i]
		}

		if err := rows.Scan(valuesPtrs...); err != nil {
			return nil, fmt.Errorf("Failed to scan row: %s", err)
		}

		rowMap := make(map[string]interface{})

		for i, colName := range columns {
			val := values[i]

			b, ok := val.([]byte)
			if ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}

		}

		results = append(results, rowMap)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return results, nil
}

func (pg *PostgresConnector) Execute(ctx context.Context, statement string, args ...interface{}) (int64, error) {
	if pg.db == nil {
		return 0, fmt.Errorf("database connection is not initialized")
	}

	result, err := pg.db.ExecContext(ctx, statement, args...)
	if err != nil {
		return 0, fmt.Errorf("Cannot be able to execute: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error getting rows affected: %v", err)
	}

	return rowsAffected, nil
}

func (pg *PostgresConnector) Close() error {
	if pg.db != nil {
		err := pg.db.Close()
		if err != nil {
			return fmt.Errorf("Error closing database: %s", err)
		}

		pg.db = nil
	}

	return nil
}
