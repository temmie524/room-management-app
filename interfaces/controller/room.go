package controller

import (
	"room_app_back/pkg"
	"room_app_back/usecases/interactors"
	"strconv"

	"github.com/labstack/echo"
)

type RoomController struct {
	ri interactors.RoomInteractor
}

func NewRoomController(ri interactors.RoomInteractor) *RoomController {
	return &RoomController{
		ri: ri,
	}
}

func (rc *RoomController) Index(c echo.Context) error {
	rooms, err := rc.ri.Rooms()
	if err != nil {

		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, rooms)
}

func (rc *RoomController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	room, err := rc.ri.RoomById(id)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, room)
}
