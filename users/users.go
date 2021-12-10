// users package represents business domain in our example

package users

type User struct {
	ID          string
	Name        string
	AccessLevel int32
}

type finder interface {
	Find(name string) (User, error)
}

type UserFinder struct {
	DB       finder
	LegacyDB finder
}

// Find checks if current user has required access (>3) level and then searches
// various databases to find the user.
func (u UserFinder) Find(currentUser User, searchableUserName string) (User, error) {

	// Our busines scenario says tht we must check current user access level
	// then search for user in the database. If user was not found in the
	// existing database, we then must check in legacy database.

	if currentUser.AccessLevel <= 2 {
		return User{}, newError("accessRestricted", "need higher access level")
	}

	user, err := u.DB.Find(searchableUserName)

	if err == nil {
		return user, err
	}

	type notFounder interface {
		NotFound() bool
	}

	if e, ok := err.(notFounder); !ok {
		return User{}, err
	} else {
		if !e.NotFound() {
			return User{}, err
		}
	}

	user, err = u.LegacyDB.Find(searchableUserName)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
