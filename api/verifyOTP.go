package api

import (
	"errors"
	"github.com/amallick86/sanyukta/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type verifyOTPRequest struct {
	Email string `json:"Email" binding:"required,email"`
	Otp   int64  `json:"Otp" binding:"required"`
}

// Verify otp
// @Summary verify otp
// @Tags AUTH
// @ID VerifyOTP
// @Accept json
// @Produce json
// @Param data body verifyOTPRequest true "email"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /auth/verification/verify-otp [post]
func (server *Server) verifyOTP(ctx *gin.Context) {
	var req verifyOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userOtp, err := server.store.GetOtpOfUser(ctx, req.Email)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("your email is already verified")))
		return
	}
	exp, err := util.VerifyOTPExpiration(userOtp.ExpiredAt.Time)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	otp, err := util.VerifyOTP(req.Otp, userOtp.OTP.Int64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if exp == true && otp == true {
		_, err = server.store.MakeUserVerify(ctx, req.Email)
		if err != nil {
			logrus.Error(err)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, successResponse("congratulation your email has been verified "))
		return
	}
	ctx.JSON(http.StatusOK, errorResponse(errors.New("enter valid input")))
	return
}
