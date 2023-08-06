package controllers

import (
	"room_app_back/domain/model"
	"room_app_back/interfaces/database"
	"room_app_back/usecases"
	"strconv"

	"github.com/labstack/echo"
)

type ReservationController struct {
	Interactor usecases.ReservationInteractor
}

func NewReservationController(sqlHandler database.SqlHandler) *ReservationController {
	return &ReservationController{
		Interactor: usecases.ReservationInteractor{
			ReservationRepository: &database.ReservationRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ReservationController) Index(c echo.Context) (err error) {
	reservations, err := controller.Interactor.Reservations()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, reservations)
	return
}

func (controller *ReservationController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation, err := controller.Interactor.ReservationById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, reservation)
	return
}

func (controller *ReservationController) Create(c echo.Context) (err error) {
	r := model.Reservation{}
	c.Bind(&r)
	reservation, err := controller.Interactor.Add(r)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, reservation)
	return
}

func (controller *ReservationController) Save(c echo.Context) (err error) {
	r := model.Reservation{}
	c.Bind(&r)
	reservation, err := controller.Interactor.Update(r)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, reservation)
	return
}

func (controller *ReservationController) Delete(c echo.Context) (err error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	reservation := model.Reservation{
		ID: uint(id),
	}
	err = controller.Interactor.DeleteById(reservation)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, "deleted")
	return
}
