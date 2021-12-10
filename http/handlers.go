package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/b-pagis/go-public-errors-example/users"
)

type finder interface {
	Find(currentUser users.User, searchableUserName string) (users.User, error)
}

type userResponse struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	AccessLevel int32  `json:"accessLevel,omitempty"`
}

type Handlers struct {
	UserFinder finder
}

func (h Handlers) FindOnlyInternalError(w http.ResponseWriter, r *http.Request) {
	currentUser, err := getUserSession(r.URL.Query().Get("currentUserID"))

	if err != nil {
		log.Println(err)

		w.Header().Add("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code":"internalError"}`))

		return
	}

	searchableUserName := r.URL.Query().Get("name")

	user, err := h.UserFinder.Find(currentUser, searchableUserName)
	if err != nil {
		log.Println(err)

		w.Header().Add("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code":"internalError"}`))

		return
	}

	respBytes, err := json.Marshal(userResponse{ID: user.ID, Name: user.Name, AccessLevel: user.AccessLevel})
	if err != nil {
		log.Println(err)

		w.Header().Add("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code":"internalError"}`))

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(respBytes)

	log.Printf("Success: %+v\n", string(respBytes))
}

func (h Handlers) FindPublicErrors(w http.ResponseWriter, r *http.Request) {
	publicErrs := publicErrors{
		"notFound":        http.StatusNotFound,
		"sessionNotFound": http.StatusForbidden,
	}

	currentUser, err := getUserSession(r.URL.Query().Get("currentUserID"))

	if err != nil {
		log.Println(err)
		if publicErrs.isPublic(err) {
			w.Header().Add("Content-Type", "application/problem+json")
			w.WriteHeader(publicErrs.HTTPStatusCode(err))
			w.Write([]byte(`{"code":"` + publicErrs.Code(err) + `", "message":"` + err.Error() + `"}`))

			return
		} else {
			w.Header().Add("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code":"internalError"}`))

			return
		}
	}

	searchableUserName := r.URL.Query().Get("name")

	user, err := h.UserFinder.Find(currentUser, searchableUserName)

	if err != nil {
		log.Println(err)
		if publicErrs.isPublic(err) {
			w.Header().Add("Content-Type", "application/problem+json")
			w.WriteHeader(publicErrs.HTTPStatusCode(err))
			w.Write([]byte(`{"code":"` + publicErrs.Code(err) + `", "message":"` + err.Error() + `"}`))

			return
		} else {
			w.Header().Add("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code":"internalError"}`))

			return
		}
	}

	respBytes, err := json.Marshal(userResponse{ID: user.ID, Name: user.Name, AccessLevel: user.AccessLevel})

	if err != nil {
		log.Println(err)
		if publicErrs.isPublic(err) {
			w.Header().Add("Content-Type", "application/problem+json")
			w.WriteHeader(publicErrs.HTTPStatusCode(err))
			w.Write([]byte(`{"code":"` + publicErrs.Code(err) + `", "message":"` + err.Error() + `"}`))

			return
		} else {
			w.Header().Add("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code":"internalError"}`))

			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(respBytes)
	log.Printf("Success: %+v\n", string(respBytes))
}

func (h Handlers) FindPublicErrorsForMid(currentUserID string, searchableUserName string) ([]byte, error) {
	publicErrs := publicErrors{
		"notFound":            http.StatusNotFound,
		"sessionNotFound":     http.StatusForbidden,
		"adminAccessRequired": http.StatusForbidden,
	}

	currentUser, err := getUserSession(currentUserID)

	if err != nil {
		return nil, handlerError{allowedList: publicErrs, originalErr: err}
	}

	user, err := h.UserFinder.Find(currentUser, searchableUserName)

	if err != nil {
		return nil, handlerError{allowedList: publicErrs, originalErr: err}
	}

	respBytes, err := json.Marshal(userResponse{ID: user.ID, Name: user.Name, AccessLevel: user.AccessLevel})

	if err != nil {
		return nil, handlerError{allowedList: publicErrs, originalErr: err}
	}

	return respBytes, nil
}

func getUserSession(id string) (users.User, error) {
	switch id {
	case "1":
		return users.User{ID: "1", Name: "Maria", AccessLevel: 1}, nil
	case "4":
		return users.User{ID: "4", Name: "Jose", AccessLevel: 4}, nil
	default:
		return users.User{}, sessionError{code: "sessionNotFound", msg: "session not found or expired"}
	}
}
