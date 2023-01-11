package db

import (
	"context"
	"github.com/amallick86/sanyukta/models"
	"github.com/google/uuid"
	"time"
)

// save user session in database
const createSession = `
INSERT INTO sessions (
	id,
	uid,
	refreshtoken,
	useragent,
	clientip,
	isblocked,
	expiresat ,
	deviceid
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING  id, uid, refreshtoken, useragent, clientip, isblocked, expiresat , deviceid, createdat
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"ID" `
	UID          uuid.UUID `json:"UID" `
	RefreshToken string    `json:"RefreshToken"`
	UserAgent    string    `json:"UserAgent" `
	ClientIP     string    `json:"ClientIP" `
	IsBlocked    bool      `json:"IsBlocked" `
	DeviceID     string    `json:"DeviceID"`
	ExpiresAt    time.Time `json:"ExpiresAt" `
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (i models.Sessions, err error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIP,
		arg.IsBlocked,
		arg.ExpiresAt,
		arg.DeviceID,
	)
	err = row.Scan(
		&i.ID,
		&i.UID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIP,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.DeviceID,
		&i.CreatedAt,
	)
	return i, err
}

// Get user session

const getSession = `
SELECT  id, uid, refreshtoken, useragent, clientip, isblocked, expiresat , deviceid, createdat FROM sessions
WHERE id = $1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (i models.Sessions, err error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	err = row.Scan(
		&i.ID,
		&i.UID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIP,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.DeviceID,
		&i.CreatedAt,
	)
	return i, err
}
