package current

import "github.com/b-pagis/go-public-errors-example/users"

type Repository struct{}

func (r Repository) Find(name string) (users.User, error) {
	switch name {
	case "Maria":
		return users.User{ID: "1", Name: "Maria", AccessLevel: 1}, nil
	case "Nushi":
		return users.User{}, internalError{code: "adminAccessRequired", msg: "need admin access level"}
	default:
		return users.User{}, internalError{code: "notFound", msg: "user not found"}
	}
}
