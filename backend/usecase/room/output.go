package room

import "backend/domain/model"

type AddOutputs []AddOutput

type AddOutput struct {
	*model.Room
}
