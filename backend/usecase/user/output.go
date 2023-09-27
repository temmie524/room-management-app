package user

import "backend/domain/model"

type AddOutputs []AddOutput

type AddOutput struct {
	*model.User
}
