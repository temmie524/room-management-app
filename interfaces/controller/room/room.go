package room

import (
	"room_app_back/usecase/room"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomController struct {
	ru room.IRoomUsecase
}

func NewRoomController(ru room.IRoomUsecase) *RoomController {
	return &RoomController{ru}
}

func (rc *RoomController) Index(c echo.Context) error {
	rooms, err := rc.ru.Rooms()
	if err != nil {

		return c.JSON(500, err)
	}
	return c.JSON(200, rooms)
}

func (rc *RoomController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	room, err := rc.ru.RoomById(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, room)
}
