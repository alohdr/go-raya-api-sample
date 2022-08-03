package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-rest-api/config"
	"go-rest-api/db"
	"go-rest-api/redis"
	"go-rest-api/utils"
	"os"
	"time"
)

type (
	Service struct {
		repository Repository
	}
)

func NewService() Service {
	return Service{
		repository: Repository{
			db:  db.Postgres(),
			rdb: redis.NewRedisClient(os.Getenv("REDIS_HOST"), ""),
		},
	}
}

func (s Service) Login(ctx context.Context, param LoginParams) (*LoginResponse, error) {
	cfg := config.Get()
	user, err := s.repository.GetUserByIdentity(ctx, param.Identity)
	if err != nil {
		return nil, err
	}

	if user.UserPass != param.Password {
		return nil, err
	}

	claims := &utils.Claims{
		Data: param.Identity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.ExpiredDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(cfg.Secret)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		ID:       user.UserID,
		Username: user.UserName,
		Email:    user.UserEmail,
		Token:    tokenString,
	}, nil

}
