package service

import (
	"account-manager-service/internal/delivery/http/exception"
	"account-manager-service/internal/entity"
	"account-manager-service/internal/model"
	"account-manager-service/internal/repository"
	"account-manager-service/internal/util"
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	ViperConfig    *viper.Viper
	Validate       *validator.Validate
	DB             *gorm.DB
	RedisClient    *redis.Client
	Log            *logrus.Logger
	UserRepository *repository.UserRepository
}

func NewAuthService(viperCfg *viper.Viper, validate *validator.Validate, db *gorm.DB, redisClient *redis.Client, log *logrus.Logger,
	userRepository *repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		ViperConfig:    viperCfg,
		Validate:       validate,
		DB:             db,
		RedisClient:    redisClient,
		Log:            log,
		UserRepository: userRepository,
	}
}

func (s *AuthServiceImpl) SignUp(ctx context.Context, request *model.SignUpRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("[%d] invalid request body : %+v", http.StatusBadRequest, err)
		return err
	}

	tx := s.DB.WithContext(ctx).Begin()

	exist, err := s.UserRepository.ExistByEmail(tx, request.Email)
	if err != nil {
		s.Log.Warnf("[%d] transaction failed : %+v", http.StatusInternalServerError, err)
		return err
	}
	if exist {
		s.Log.Warnf("[%d] duplicate email : %+v", http.StatusConflict, err)
		return exception.NewHttpError(http.StatusConflict, "email already exists")
	}

	exist, err = s.UserRepository.ExistByPhoneNumber(tx, request.PhoneNumber)
	if err != nil {
		s.Log.Warnf("[%d] transaction failed : %+v", http.StatusInternalServerError, err)
		return err
	}
	if exist {
		s.Log.Warnf("[%d] duplicate phone number : %+v", http.StatusConflict, err)
		return exception.NewHttpError(http.StatusConflict, "phone number already exists")
	}

	newPassword, err := util.GenerateFromPassword(request.Password)
	if err != nil {
		s.Log.Warnf("[%d] failed to generate password : %+v", http.StatusInternalServerError, err)
		return err
	}
	newUser := &entity.User{
		ID:          uuid.NewString(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Password:    newPassword,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
	}
	if err := s.UserRepository.Repository.Create(tx, newUser); err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			s.Log.Warnf("[%d] failed to rollback : %+v", http.StatusInternalServerError, err)
			return err
		}

		s.Log.Warnf("[%d] transaction failed : %+v", http.StatusInternalServerError, err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("[%d] failed to commit : %+v", http.StatusInternalServerError, err)
		return err
	}

	return nil
}
