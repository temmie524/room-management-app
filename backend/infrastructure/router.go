package infrastructure

import (
	"backend/config"
	rere "backend/infrastructure/repository/reservation"
	rore "backend/infrastructure/repository/room"
	usre "backend/infrastructure/repository/user"
	"net/http"
	"os"

	reco "backend/interfaces/controller/reservation"
	roco "backend/interfaces/controller/room"
	usco "backend/interfaces/controller/user"

	reus "backend/usecase/reservation"
	rous "backend/usecase/room"
	usus "backend/usecase/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	db := NewDB()
	cnf := config.NewAppConfig()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", cnf.FeUrl},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   cnf.ApiDomain,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, // TODO: 他の場合はNoneモードで！
		//CookieSameSite: http.SameSiteDefaultMode, // TODO: POSTMANの動作確認時はDefaultMode
	}))

	userRepository := usre.NewUserRepository(db)
	userUsecase := usus.NewUserUsecase(userRepository, cnf)
	userController := usco.NewUserController(userUsecase, cnf)

	reservationRepository := rere.NewReservationRepository(db)
	reservationUsecase := reus.NewReservationUsecase(reservationRepository, cnf)
	reservationController := reco.NewReservationController(reservationUsecase)

	roomRepository := rore.NewRoomRepository(db)
	roomUsecase := rous.NewRoomUsecase(roomRepository)
	roomController := roco.NewRoomController(roomUsecase)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//User Routes
	e.GET("/users", userController.Index)
	e.GET("/users/:id", userController.Show)
	e.POST("/new", userController.Create)
	e.POST("/signup", userController.SignUp)
	e.POST("/login", userController.LogIn)
	e.POST("/logout", userController.LogOut)
	e.GET("/csrf", userController.CsrfToken)
	e.PUT("/users/:id", userController.Save)
	e.DELETE("/users/:id", userController.Delete)

	//Reservation Routes
	rs := e.Group("/reservations")
	rs.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	rs.GET("", reservationController.Index)
	rs.GET("/:id", reservationController.Show)
	rs.POST("/new", reservationController.Create)
	rs.PUT("/:id", reservationController.Save)
	rs.DELETE("/:id", reservationController.Delete)

	rm := e.Group("/rooms")
	rm.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	//Room Routes
	rm.GET("", roomController.Index)
	rm.GET("/:id", roomController.Show)

	e.Logger.Fatal(e.Start(":8080"))
}
