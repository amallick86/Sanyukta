package util

import (
	"errors"
	"github.com/amallick86/sanyukta/models"
	"time"
)

func CheckCreatedDateOfOTP(createdAt time.Time, count int16) (dayCount int16) {

	now := time.Now()
	h := now.Hour()
	m := now.Minute()
	s := now.Second()
	midNight := now.Add(-time.Hour*time.Duration(h) +
		-time.Minute*time.Duration(m) +
		-time.Second*time.Duration(s))
	if createdAt.After(midNight) {
		return count
	}
	return dayCount
}
func CheckOtpValidation(userOtp models.EmailOTP) (int64, time.Time, error) {
	if userOtp.ID.Int64 != 0 {
		if time.Now().Before(userOtp.ExpiredAt.Time) {
			return 0, userOtp.ExpiredAt.Time, errors.New("old otp hasn't expired ")
		} else {
			otp := RandomOTP()
			expirationTime, err := SetExpirationOfOTP(userOtp.CountOfDay.Int16)
			return otp, expirationTime, err
		}
	} else {
		otp := RandomOTP()
		expirationTime, err := SetExpirationOfOTP(userOtp.CountOfDay.Int16)
		return otp, expirationTime, err
	}

}
func SetExpirationOfOTP(count int16) (t time.Time, err error) {
	if count < 3 {
		switch count {
		case 0:
			return time.Now().Add(time.Minute * 2), err
		case 1:
			return time.Now().Add(time.Minute * 4), err
		default:
			return time.Now().Add(time.Minute * 8), err
		}
	}
	err = errors.New("number of otp has excide for today")
	return t, err
}

func VerifyOTP(reqOtp, dbOtp int64) (b bool, err error) {
	if reqOtp == dbOtp {
		return true, err
	}
	return false, errors.New("please enter valid otp")
}
func VerifyOTPExpiration(dbOtpExp time.Time) (b bool, err error) {
	if time.Now().Before(dbOtpExp) {
		return true, err
	}
	return false, errors.New("otp has been expired")
}
