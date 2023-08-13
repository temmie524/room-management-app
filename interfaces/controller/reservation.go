package controller

import (
	"room_app_back/domain/model"
	"room_app_back/pkg"
	"room_app_back/usecases/interactors"
	"strconv"

	"github.com/labstack/echo"
)

type ReservationController struct {
	ri interactors.ReservationInteractor
}

func NewReservationController(ri interactors.ReservationInteractor) *ReservationController {
	return &ReservationController{
		ri: ri,
	}
}

func (rc *ReservationController) Index(c echo.Context) error {
	reservations, err := rc.ri.Reservations()
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, reservations)
}

func (rc *ReservationController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation, err := rc.ri.ReservationById(id)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}

	return c.JSON(200, reservation)
}

func (rc *ReservationController) Create(c echo.Context) error {
	r := model.Reservation{}
	c.Bind(&r)
	reservation, err := rc.ri.Add(r)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, reservation)
}

func (rc *ReservationController) Save(c echo.Context) error {
	r := model.Reservation{}
	c.Bind(&r)
	reservation, err := rc.ri.Update(r)
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
	err := rc.ri.DeleteById(reservation)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, "deleted")
}
