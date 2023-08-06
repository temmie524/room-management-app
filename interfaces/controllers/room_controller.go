package controllers

import (
	"room_app_back/interfaces/database"
	"room_app_back/usecases"
	"strconv"

	"github.com/labstack/echo"
)

type RoomController struct {
	Interactor usecases.RoomInteractor
}

func NewRoomController(sqlHandler database.SqlHandler) *RoomController {
	return &RoomController{
		Interactor: usecases.RoomInteractor{
			RoomRepository: &database.RoomRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RoomController) Index(c echo.Context) (err error) {
	rooms, err := controller.Interactor.Rooms()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, rooms)
	return
}

func (controller *RoomController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	room, err := controller.Interactor.RoomById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}
