package api

import (
	"errors"
	"github.com/amallick86/sanyukta/db"
	"github.com/amallick86/sanyukta/models"
	"github.com/amallick86/sanyukta/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type createAccountRequest struct {
	FirstName string `json:"FirstName" binding:"required,alphanum"`
	LastName  string `json:"LastName" binding:"required,alphanum"`
	Email     string `json:"Email" binding:"required,email"`
	UserName  string `json:"UserName" binding:"required,alphanum"`
	Password  string `json:"password" binding:"required,min=6"`
}

// CreateUser handles request for user creation
// @Summary create a new user
// @Tags AUTH
// @ID CreateUser
// @Accept json
// @Produce json
// @Param data body createAccountRequest true "create user"
// @Success 201 {object} sendOTPResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /auth/signup [post]
func (server *Server) createUser(ctx *gin.Context) {

	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

	}
	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		UserName:  req.UserName,
		Password:  hashedPassword,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			logrus.Error(err)
			ctx.JSON(http.StatusConflict, errorResponse(err))
			return
		}
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	otp, expTime, err := util.CheckOtpValidation(models.EmailOTP{
		ID:         util.CreateNullInt64(false, 0),
		UID:        user.UID,
		FullName:   req.FirstName + " " + req.LastName,
		OTP:        util.CreateNullInt64(false, 0),
		CountOfDay: util.CreateNullInt16(false, 0),
		CreatedAt:  util.CreateNUllTime(false, time.Time{}),
		ExpiredAt:  util.CreateNUllTime(false, time.Time{}),
	})
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
		UserName:           req.FirstName + " " + req.LastName,
		UserEmail:          req.Email,
		OTP:                strconv.FormatInt(otp, 10),
	}
	err = util.SendEmailVerificationCode(param)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("unable to send email")))
	}
	_, err = server.store.AddOtpOfUser(ctx, db.AddOtpOfUserParam{
		UID:        user.UID,
		OTP:        otp,
		CountOfDay: 1,
		ExpiredAt:  expTime,
	})
	ctx.JSON(http.StatusCreated, sendOTPResponse{
		ExpiredAt: expTime,
		Message:   "congratulation your account has been successfully created",
	})
	return
}
