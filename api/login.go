package api

import (
	"database/sql"
	"github.com/amallick86/sanyukta/db"
	"github.com/amallick86/sanyukta/models"
	"github.com/amallick86/sanyukta/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type loginRequest struct {
	UserNameOREmail string `json:"UserNameOREmail" binding:"required"`
	Password        string `json:"Password" binding:"required,min=6"`
}
type createAccountResponse struct {
	UID        uuid.UUID `json:"UID" `
	FirstName  string    `json:"FirstName" `
	LastName   string    `json:"LastName" `
	Email      string    `json:"Email" `
	UserName   string    `json:"UserName" `
	IsVerified bool      `json:"IsVerified" `
	CreatedAt  time.Time `json:"CreatedAt" `
}

func newUserResponse(user models.User) createAccountResponse {
	return createAccountResponse{
		UID:        user.UID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		UserName:   user.UserName,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
	}
}

type loginResponse struct {
	SessionId             uuid.UUID             `json:"session_id"`
	AccessToken           string                `json:"accessToken"`
	AccessTokenExpiresAt  time.Time             `json:"accessTokenExpiresAt"`
	RefreshToken          string                `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time             `json:"refreshTokenExpiresAt"`
	User                  createAccountResponse `json:"user"`
}

// Login handles request for user creation
// @Summary login and generate token with JWT
// @Tags AUTH
// @ID Login
// @Accept json
// @Produce json
// @Param deviceId query string true "DeviceId"
// @Param data body loginRequest true "Login request"
// @Success 200 {object} loginResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /auth/login [post]
func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	var res loginResponse
	var user models.User
	var err error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	deviceId := ctx.DefaultQuery("deviceId", "no id")
	if util.IsValidEmail(req.UserNameOREmail) == true {
		user, err = server.store.GetUserByEmail(ctx, req.UserNameOREmail)
		if err != nil {
			if err == sql.ErrNoRows {
				logrus.Error(err)
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			logrus.Error(err)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		user, err = server.store.GetUserByUserName(ctx, req.UserNameOREmail)
		if err != nil {
			if err == sql.ErrNoRows {
				logrus.Error(err)
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			logrus.Error(err)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken, acessPayload, err := server.tokenMaker.CreateToken(
		user.UID,
		server.config.AccessTokenDuration)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.UID,
		server.config.RefreshTokenDuration)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UID:          user.UID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		DeviceID:     deviceId,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res = loginResponse{
		SessionId:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  acessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, res)
}
