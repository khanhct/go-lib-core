package service

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	Orm    *gorm.DB
	Logger *logrus.Logger
	Error  error
}

func (db *Service) AddError(err error) error {
	if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}

	return db.Error
}
