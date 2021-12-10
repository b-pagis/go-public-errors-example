package main

import (
	"fmt"
	"strings"

	"github.com/b-pagis/go-public-errors-example/repository/current"
	"github.com/b-pagis/go-public-errors-example/repository/legacy"
	"github.com/b-pagis/go-public-errors-example/users"
)

type userService interface {
	Find(current users.User, searchableUserName string) (users.User, error)
}

type scenario struct {
	UserService userService
}

func main() {
	usersService := users.UserFinder{
		DB:       current.Repository{},
		LegacyDB: legacy.Repository{},
	}

	scenarios := scenario{
		UserService: usersService,
	}

	scenarios.firstExample()
	scenarios.secondExample()
}

func (s scenario) firstExample() {
	fmt.Print("\n***************\n* Scenario #1 *\n***************\n")

	name := "Maria"
	currentUser := users.User{ID: "2", Name: "Nushi", AccessLevel: 2}

	fmt.Printf("action: \t searching for: %s\n", name)

	_, err := s.UserService.Find(currentUser, name)
	if err != nil {
		fmt.Printf("outcome:\t failed to get user %s. Error: %v\n", name, err.Error())
	}
}

func (s scenario) secondExample() {
	fmt.Print("\n***************\n* Scenario #2 *\n***************\n")

	usersToFind := []string{"Maria", "Nushi", "Mohammed", "Jose", "Wei"}
	currentUser := users.User{ID: "3", Name: "Mohammed", AccessLevel: 3}

	for _, name := range usersToFind {
		fmt.Printf("\naction: \t searching for: %v\n", name)

		user, err := s.UserService.Find(currentUser, name)

		if err != nil {
			// It is not the best practice to handle error in such manner
			// and it should be avoided if possible. It would be better to
			// handle the error where it was received and return new internal
			// error. Sadly, it is not always possible to do things the way
			// we want.
			if strings.Contains(err.Error(), "third party") {

				fmt.Println("outcome:\t issue with third party, please contact support@thirdparty")

				continue
			}

			fmt.Printf("outcome:\t failed to get user: %v. Error: %v\n", name, err.Error())

			continue
		}

		fmt.Printf("outcome:\t found user: %+v\n", user)
	}
}
