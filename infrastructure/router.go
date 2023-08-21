package infrastructure

import (
	rere "room_app_back/infrastructure/repository/reservation"
	rore "room_app_back/infrastructure/repository/room"
	usre "room_app_back/infrastructure/repository/user"

	reco "room_app_back/interfaces/controller/reservation"
	roco "room_app_back/interfaces/controller/room"
	usco "room_app_back/interfaces/controller/user"

	reus "room_app_back/usecase/reservation"
	rous "room_app_back/usecase/room"
	usus "room_app_back/usecase/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	sqlhandler := NewSqlHandler()

	userRepository := usre.NewUserRepository(sqlhandler)
	userUsecase := usus.NewUserUsecase(*userRepository)
	userController := usco.NewUserController(*userUsecase)

	reservationRepository := rere.NewReservationRepository(sqlhandler)
	reserationUsecase := reus.NewReservationUsecase(*reservationRepository)
	reservationController := reco.NewReservationController(*reserationUsecase)

	roomRepository := rore.NewRoomRepository(sqlhandler)
	roomUsecase := rous.NewRoomUsecase(*roomRepository)
	roomController := roco.NewRoomController(*roomUsecase)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	//User Routes
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/new", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	//Reservation Routes
	e.GET("/reservations", func(c echo.Context) error { return reservationController.Index(c) })
	e.GET("/reservations/:id", func(c echo.Context) error { return reservationController.Show(c) })
	e.POST("/reservations/new", func(c echo.Context) error { return reservationController.Create(c) })
	e.PUT("/reservations/:id", func(c echo.Context) error { return reservationController.Save(c) })
	e.DELETE("/reservations/:id", func(c echo.Context) error { return reservationController.Delete(c) })

	//Room Routes
	e.GET("/rooms", func(c echo.Context) error { return roomController.Index(c) })
	e.GET("/rooms/:id", func(c echo.Context) error { return roomController.Show(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
