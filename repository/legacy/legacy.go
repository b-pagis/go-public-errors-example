package legacy

import (
	"errors"

	"github.com/b-pagis/go-public-errors-example/users"
)

type Repository struct{}

func (r Repository) Find(name string) (users.User, error) {
	switch name {
	case "Mohammed":
		return users.User{ID: "3", Name: "Mohammed", AccessLevel: 3}, nil
	case "Jose":
		return users.User{}, errors.New("some kind of third party internal error")
	default:
		return users.User{}, internalError{code: "notFound", msg: "user not found"}
	}
}
