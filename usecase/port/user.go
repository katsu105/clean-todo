package port

import (
	"context"

	"github.com/katsu105/clean-todo/entity"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
