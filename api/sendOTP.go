package api

import (
	"errors"
	"github.com/amallick86/sanyukta/db"
	"github.com/amallick86/sanyukta/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type sendOTPRequest struct {
	Email string `json:"Email" binding:"required,email"`
}
type sendOTPResponse struct {
	ExpiredAt time.Time `json:"ExpiredAt" `
	Message   string    `json:"Message"`
}

// Send otp
// @Summary send otp
// @Tags AUTH
// @ID SendOTP
// @Accept json
// @Produce json
// @Param data body sendOTPRequest true "send otp"
// @Success 200 {object} sendOTPResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /auth/verification/send-otp [post]
func (server *Server) sendOTP(ctx *gin.Context) {
	var req sendOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userOtp, err := server.store.GetOtpOfUser(ctx, req.Email)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("you are unable to use this service")))
		return
	}
	dayCount := util.CheckCreatedDateOfOTP(userOtp.CreatedAt.Time, userOtp.CountOfDay.Int16)
	userOtp.CountOfDay.Int16 = dayCount
	otp, expTime, err := util.CheckOtpValidation(userOtp)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	param := util.SendEmailVerificationCodeParam{
		SenderEmail:        server.config.SenderEmail,
		SendInBlueEmail:    server.config.SendInBlueEmail,
		SendInBluePassword: server.config.SendInBluePassword,
		SendInBlueAddress:  server.config.ServerAddress,
		SMTPHost:           server.config.SMTPHost,
		UserName:           userOtp.FullName,
		UserEmail:          req.Email,
		OTP:                strconv.FormatInt(otp, 10),
	}
	err = util.SendEmailVerificationCode(param)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("unable to send email")))
	}
	_, err = server.store.AddOtpOfUser(ctx, db.AddOtpOfUserParam{
		UID:        userOtp.UID,
		OTP:        otp,
		CountOfDay: userOtp.CountOfDay.Int16 + 1,
		ExpiredAt:  expTime,
	})
	ctx.JSON(http.StatusOK, sendOTPResponse{
		ExpiredAt: expTime,
		Message:   "otp has been send to your email",
	})
}
