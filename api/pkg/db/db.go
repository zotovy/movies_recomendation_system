package db

import (
	"app/pkg/config"
	"app/pkg/logger"
	"app/pkg/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type DB = pgx.Conn

func NewDB(cfg *config.Config, logger *logger.Logger) (conn *pgx.Conn, err error) {
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		connStr := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			cfg.PostgreSQL.Host,
			cfg.PostgreSQL.Username,
			cfg.PostgreSQL.Password,
			cfg.PostgreSQL.Database,
			cfg.PostgreSQL.Port,
		)

		connConfig, err := pgx.ParseConfig(connStr)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		connConfig.Tracer = &myQueryTracer{
			log: logger,
		}

		conn, err = pgx.ConnectConfig(ctx, connConfig)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		return nil
	}, 5, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return conn, nil
}

type myQueryTracer struct {
	log *logger.Logger
}

func (tracer *myQueryTracer) TraceQueryStart(
	ctx context.Context,
	_ *pgx.Conn,
	data pgx.TraceQueryStartData) context.Context {
	tracer.log.Info(
		fmt.Sprintf("Executing command  %s, %s", data.SQL, data.Args),
	)

	return ctx
}

func (tracer *myQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}
