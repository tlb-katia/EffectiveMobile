package postgres

import (
	"EffectiveMobile_Project/config"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"log/slog"
)

type Storage struct {
	Db      *sql.DB
	Driver  database.Driver
	Log     *slog.Logger
	Builder squirrel.StatementBuilderType
}

func NewStorage(conf *config.Config, log *slog.Logger) (*Storage, error) {
	const op = "storage.postgres"

	conStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		conf.PGHost, conf.PGPort, conf.PGUser, conf.PGName, "disable", conf.PGPassword)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("%s %s", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s %s", op, err)
	}

	driver, _ := postgres.WithInstance(db, &postgres.Config{})

	return &Storage{
		Db:      db,
		Driver:  driver,
		Log:     log,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}, nil
}
