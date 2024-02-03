package utils

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"strings"
)

func RunMigrate(dbStr string) error {
	m, err := migrate.New(
		"file://migrations",
		dbStr,
	)
	if err != nil {
		return err
	}
	err = m.Up()
	if strings.Contains(err.Error(), "no change") {
		log.Info("Migrate: Nothing to change")
		return nil
	}
	return err
}
