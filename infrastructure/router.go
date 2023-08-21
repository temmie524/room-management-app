package infrastructure

import (
	"room_app_back/domain/repository"
	"room_app_back/interfaces/controller"
	"room_app_back/usecases/interactors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	sqlhandler := NewSqlHandler()

	userRepository := repository.NewUserRepository(sqlhandler)
	userInteractor := interactors.NewUserInteractor(userRepository)
	userController := controller.NewUserController(*userInteractor)

	reservationRepository := repository.NewReservationRepository(sqlhandler)
	reserationInteractor := interactors.NewReservationInteractor(reservationRepository)
	reservationController := controller.NewReservationController(*reserationInteractor)

	roomRepository := repository.NewRoomRepository(sqlhandler)
	roomInteractor := interactors.NewRoomInteractor(roomRepository)
	roomController := controller.NewRoomController(*roomInteractor)

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
