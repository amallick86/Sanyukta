package db

import (
	"context"
	"github.com/amallick86/sanyukta/models"
	"github.com/google/uuid"
)

const getUserById = ` SELECT uid, firstname,lastname,email,username,isverified,subscriptionid, createdat FROM users  WHERE uid = $1`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (i models.User, err error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	err = row.Scan(
		&i.UID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.UserName,
		&i.IsVerified,
		&i.SubscriptionID,
		&i.CreatedAt,
	)

	return i, err
}

const createUser = `
INSERT INTO users(
 firstname,lastname,email,username,password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING uid, firstname,lastname,email,username, isverified, createdat
`

type CreateUserParams struct {
	FirstName string `json:"FirstName" `
	LastName  string `json:"LastName" `
	Email     string `json:"Email" `
	UserName  string `json:"UserName" `
	Password  string `json:"password" `
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (i models.User, err error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.FirstName, arg.LastName, arg.Email,
		arg.UserName, arg.Password)
	err = row.Scan(
		&i.UID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.UserName,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUserName = `
SELECT uid, username, password, firstname,lastname,email,isverified, createdat  FROM users
WHERE username = $1 
`

func (q *Queries) GetUserByUserName(ctx context.Context, userName string) (i models.User, err error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, userName)
	err = row.Scan(
		&i.UID,
		&i.UserName,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `
SELECT uid, email, password, firstname,lastname,isverified, createdat FROM users
WHERE email = $1 
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (i models.User, err error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	err = row.Scan(
		&i.UID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}

const makeUserVerify = `
Update users SET isverified = $1 WHERE email = $2
RETURNING uid, firstname,lastname,email,username, isverified, createdat
`

func (q *Queries) MakeUserVerify(ctx context.Context, email string) (i models.User, err error) {
	row := q.db.QueryRowContext(ctx, makeUserVerify, true, email)
	err = row.Scan(
		&i.UID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.UserName,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}
