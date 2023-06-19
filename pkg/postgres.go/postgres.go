package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/cucumberjaye/GophKeeper/internal/app/migration"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dbDriver = "pgx"
)

// New - создает подключение к Postgres и добавляет миграции.
func New(dbCfg *configs.DBConnConfig) (*pgxpool.Pool, error) {
	dataBase, err := sql.Open(dbDriver,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.User,
			dbCfg.Password,
			dbCfg.DBName,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("open database failed with error: %w", err)
	}
	goose.SetBaseFS(migration.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("set goose dialect failed with error: %w", err)
	}

	if err := goose.Up(dataBase, "migrations"); err != nil {
		return nil, fmt.Errorf("migrations up failed with error: %w", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%s",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
		dbCfg.PoolMaxConns,
	)

	repCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse data base config failed with error: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), repCfg)
	if err != nil {
		return nil, fmt.Errorf("create conn pool failed with error: %w", err)
	}
	return pool, nil
}
