package repository

import (
	"github.com/tama-jp/basee_web/internal/frameworks/config"
	db "github.com/tama-jp/basee_web/internal/frameworks/database"
	"github.com/tama-jp/basee_web/internal/frameworks/logger"
	"github.com/tama-jp/basee_web/internal/usecases/port"
)

type loggerRepository struct {
	conf *config.Config
	db   *db.DataBase
	log  *logger.LogBase
}

func NewLoggerRepository(conf *config.Config, db *db.DataBase, log *logger.LogBase) port.LoggerPort {
	return &loggerRepository{conf, db, log}
}

func (r *loggerRepository) PrintInfo(num string, group string, message string) {
	r.log.PrintInfo(num, group, message)
}

func (r *loggerRepository) PrintError(num string, group string, message string) {
	r.log.PrintError(num, group, message)
}

func (r *loggerRepository) PrintDebug(num string, group string, message string) {
	r.log.PrintDebug(num, group, message)
}
