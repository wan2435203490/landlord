package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	r "landlord/common/response"
	"landlord/core/logger"
)

type IService interface {
	Get() *Service
}

type Service struct {
	Context *gin.Context
	Orm     *gorm.DB
	Msg     string
	MsgID   string
	Log     *logger.Helper
	Err     error
}

func (s *Service) Get() *Service {
	return s
}

func (s *Service) AddError(err error) error {
	if s.Err == nil {
		s.Err = err
	} else if err != nil {
		s.Err = fmt.Errorf("%v; %w", s.Err, err)
	}
	return s.Err
}

func (s *Service) Error(err error) {
	r.ErrorInternal(err.Error(), s.Context)
}

func (s *Service) ErrorMsg(msg string) {
	r.ErrorInternal(msg, s.Context)
}
