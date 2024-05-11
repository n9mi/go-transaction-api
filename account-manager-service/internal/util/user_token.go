package util

import (
	"account-manager-service/internal/delivery/http/exception"
	"account-manager-service/internal/model"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Returns jwt token that contains AuthData, its expiration time in unix time format, and error while generating
func GenerateAuthToken(key string, expMinutes int, authData *model.AuthData) (string, int64, error) {
	timeDuration := time.Duration(expMinutes) * time.Minute
	expAtUnix := time.Now().Add(timeDuration).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  expAtUnix,
		"data": *authData,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", 0, exception.NewHttpError(http.StatusInternalServerError, "failed to generate auth token")
	}

	return s, expAtUnix, nil
}

func GenerateAccessToken(viperCfg *viper.Viper, authData *model.AuthData) (string, int64, error) {
	key := viperCfg.GetString("ACCESS_TOKEN_KEY")
	expMinutes := viperCfg.GetInt("ACCESS_TOKEN_EXPIRE_MINUTES")

	return GenerateAuthToken(key, expMinutes, authData)
}

// Returns parsed AuthData from jwt token
func ParseAuthToken(key string, token string) (*model.AuthData, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, exception.NewHttpError(http.StatusBadRequest, "invalid token")
	}

	// Map claims to AuthData
	claims := t.Claims.(jwt.MapClaims)
	authDataMap, ok := claims["data"].(map[string]interface{})
	if !ok {
		return nil, exception.NewHttpError(http.StatusBadRequest, "invalid token")
	}
	var authData model.AuthData
	authData.UserID, ok = authDataMap["UserID"].(string)
	if !ok {
		return nil, exception.NewHttpError(http.StatusBadRequest, "invalid token")
	}
	authData.Email, ok = authDataMap["Email"].(string)
	if !ok {
		return nil, exception.NewHttpError(http.StatusBadRequest, "invalid token")
	}

	return &authData, nil
}

func ParseAccessToken(viperCfg *viper.Viper, token string) (*model.AuthData, error) {
	key := viperCfg.GetString("ACCESS_TOKEN_KEY")

	return ParseAuthToken(key, token)
}

// Access token stored in redis, with redis key that match this format
func GetAccessTokenRedisKey(userID string) string {
	return fmt.Sprintf("accessKey:%s", userID)
}

// Extract AuthData from token to get UserID, use UserID to get redis key,
// then retrieve and check if given token from request and stored user's redis token are same
// Returns AuthData if given token is valid, returns error if token is invalid
func VerifyAccessToken(ctx context.Context, viperCfg *viper.Viper, redisClient *redis.Client, log *logrus.Logger,
	tokenFromRequest string) (*model.AuthData, error) {
	tokenData, err := ParseAccessToken(viperCfg, tokenFromRequest)
	if err != nil {
		log.Warnf("[%d] failed to parse token %+v", http.StatusBadRequest, err)
		return nil, err
	}

	tokenFromRedis, err := redisClient.Get(ctx, GetAccessTokenRedisKey(tokenData.UserID)).Result()
	if err != nil {
		log.Warnf("[%d] failed to get access token from redis with given key : %+v", http.StatusBadRequest, err)
		return nil, exception.NewHttpError(http.StatusBadRequest, "invalid token")
	}

	if tokenFromRedis != tokenFromRequest {
		log.Warnf("[%d] invalid token from request", http.StatusBadRequest)
		return nil, exception.NewHttpError(http.StatusBadRequest, "invalid token")
	}

	return tokenData, nil
}
