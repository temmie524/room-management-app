package room

import (
	"room_app_back/pkg"
	"room_app_back/usecase/room"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomController struct {
	ru room.RoomUsecase
}

func NewRoomController(ru room.RoomUsecase) *RoomController {
	return &RoomController{
		ru: ru,
	}
}

func (rc *RoomController) Index(c echo.Context) error {
	rooms, err := rc.ru.Rooms()
	if err != nil {

		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, rooms)
}

func (rc *RoomController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	room, err := rc.ru.RoomById(id)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, room)
}
