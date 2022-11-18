package service

import (
	"fmt"

	"git.code.oa.com/ad/go-project-base/logger"
	"gorm.io/gorm"
)

// Service ...
type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logger.Helper
	Error error
}

// AddError ...
func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}
