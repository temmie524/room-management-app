package reservation

import (
	"room_app_back/domain/model"
	"room_app_back/pkg"
	"room_app_back/usecase/reservation"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationController struct {
	ru reservation.ReservationUsecase
}

func NewReservationController(ru reservation.ReservationUsecase) *ReservationController {
	return &ReservationController{
		ru: ru,
	}
}

func (rc *ReservationController) Index(c echo.Context) error {
	reservations, err := rc.ru.Reservations()
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, reservations)
}

func (rc *ReservationController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation, err := rc.ru.ReservationById(id)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}

	return c.JSON(200, reservation)
}

func (rc *ReservationController) Create(c echo.Context) error {
	r := model.Reservation{}
	c.Bind(&r)
	reservation, err := rc.ru.Add(r)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, reservation)
}

func (rc *ReservationController) Save(c echo.Context) error {
	r := model.Reservation{}
	c.Bind(&r)
	reservation, err := rc.ru.Update(r)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, reservation)
}

func (rc *ReservationController) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	reservation := model.Reservation{
		ID: uint(id),
	}
	err := rc.ru.DeleteById(reservation)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, "deleted")
}
