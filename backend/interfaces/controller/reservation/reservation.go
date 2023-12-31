package reservation

import (
	"backend/usecase/reservation"
	"net/http"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	reservation, err := rc.ru.ReservationById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, reservation)
}

func (rc *ReservationController) Create(c echo.Context) error {
	var r reservation.AddInput
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	jwt := cookie.Value
	userId, err := rc.ru.JwtToUserId(jwt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	reservation, err := rc.ru.Add(&r, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, reservation)
}

func (rc *ReservationController) Save(c echo.Context) error {
	var r reservation.AddInput
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
		return c.JSON(http.StatusInternalServerError, err)
	}
	if err := rc.ru.DeleteById(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
