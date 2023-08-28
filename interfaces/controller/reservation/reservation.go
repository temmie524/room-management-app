package reservation

import (
	"net/http"
	"room_app_back/domain/model"
	"room_app_back/usecase/reservation"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationController struct {
	ru reservation.IReservationUsecase
}

func NewReservationController(ru reservation.IReservationUsecase) *ReservationController {
	return &ReservationController{ru}
}

func (rc *ReservationController) Index(c echo.Context) error {
	reservations, err := rc.ru.Reservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, reservations)
}

func (rc *ReservationController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation, err := rc.ru.ReservationById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, reservation)
}

func (rc *ReservationController) Create(c echo.Context) error {
	r := model.Reservation{}
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	reservation, err := rc.ru.Add(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, reservation)
}

func (rc *ReservationController) Save(c echo.Context) error {
	r := model.Reservation{}
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	reservation, err := rc.ru.Update(&r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, reservation)
}

func (rc *ReservationController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if err := rc.ru.DeleteById(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
