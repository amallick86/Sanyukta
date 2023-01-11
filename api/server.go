package api

import (
	"fmt"
	db "github.com/amallick86/sanyukta/db"
	"github.com/amallick86/sanyukta/docs"
	"github.com/amallick86/sanyukta/token"
	"github.com/amallick86/sanyukta/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

// @Security bearerAuth
func (server *Server) setupRouter() {
	//mode of gin dev mode or release mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Sanyukta"
	docs.SwaggerInfo.Description = "Sanyukta API'S"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//cors middleware
	router.Use(cors.Default())

	hello := router.Group("/hello")
	{
		hello.GET("/", server.Hello)
	}
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/signup", server.createUser)
		authRouter.POST("/login", server.login)
		tokenRouter := authRouter.Group("/refresh-token")
		{
			tokenRouter.POST("/", server.renewAccessToken)
		}
		verificationRouter := authRouter.Group("/verification")
		{
			verificationRouter.POST("/send-otp", server.sendOTP)
			verificationRouter.POST("verify-otp", server.verifyOTP)
		}
	}

	//swager route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	//server route
	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type Err struct {
	Error string `json:"error"`
}

// error response function
func errorResponse(err error) Err {

	return Err{Error: err.Error()}
}

type SuccessResponse struct {
	Message string `json:"message" `
}

// success response function
func successResponse(msg string) SuccessResponse {

	return SuccessResponse{Message: msg}
}
