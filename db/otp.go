package db

import (
	"context"
	"github.com/amallick86/sanyukta/models"
	"github.com/google/uuid"
	"time"
)

const getOtpOfUser = `SELECT CONCAT(u.firstname, ' ', u.lastname) as FullName, u.uid,o.id, o.otp, o.countofday AS countofday, o.createdat, o.expiredat
FROM users as u LEFT JOIN   emailotp as o
ON o.UID = u.UID
WHERE u.Email = $1 AND u.isverified = false
order by o.createdat DESC ;`

func (q *Queries) GetOtpOfUser(ctx context.Context, email string) (i models.EmailOTP, err error) {

	row := q.db.QueryRowContext(ctx, getOtpOfUser, email)
	err = row.Scan(
		&i.FullName,
		&i.UID,
		&i.ID,
		&i.OTP,
		&i.CountOfDay,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const addOtpOfUser = `INSERT INTO emailotp(
 uid,otp,countofday,expiredat
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, uid,otp,countofday,createdat,expiredat`

type AddOtpOfUserParam struct {
	UID        uuid.UUID `json:"UID" `
	OTP        int64     `json:"OTP"`
	CountOfDay int16     `json:"CountOfDay"`
	ExpiredAt  time.Time `json:"ExpiredAt" `
}

func (q *Queries) AddOtpOfUser(ctx context.Context, param AddOtpOfUserParam) (i models.EmailOTP, err error) {

	row := q.db.QueryRowContext(ctx, addOtpOfUser, param.UID, param.OTP, param.CountOfDay, param.ExpiredAt)
	err = row.Scan(
		&i.ID,
		&i.UID,
		&i.OTP,
		&i.CountOfDay,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}
