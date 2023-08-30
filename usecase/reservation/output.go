package reservation

import "room_app_back/domain/model"

type AddOutputs []AddOutput

type AddOutput struct {
	*model.Reservation
}
