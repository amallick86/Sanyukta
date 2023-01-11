package db

import (
	"context"
	"github.com/amallick86/sanyukta/models"
	"github.com/google/uuid"
)

type Querier interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (i models.User, err error)
	CreateUser(ctx context.Context, arg CreateUserParams) (i models.User, err error)
	GetUserByUserName(ctx context.Context, username string) (i models.User, err error)
	GetUserByEmail(ctx context.Context, email string) (i models.User, err error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (i models.Sessions, err error)
	GetSession(ctx context.Context, id uuid.UUID) (i models.Sessions, err error)
	GetOtpOfUser(ctx context.Context, email string) (i models.EmailOTP, err error)
	AddOtpOfUser(ctx context.Context, param AddOtpOfUserParam) (i models.EmailOTP, err error)
	MakeUserVerify(ctx context.Context, email string) (i models.User, err error)
}

var _ Querier = (*Queries)(nil)
