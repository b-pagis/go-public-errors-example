package http

import (
	"log"
	"net/http"
)

func HandleError(handle func(string, string) ([]byte, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUserID := r.URL.Query().Get("currentUserID")
		searchableUserName := r.URL.Query().Get("name")
		
		response, err := handle(currentUserID, searchableUserName)

		if err != nil {
			log.Println(err)
			w.Header().Add("Content-Type", "application/problem+json")

			type handlerErr interface {
				Public() bool
				HTTPStatusCode() int
				Code() string
				Error() string
			}

			if e, ok := err.(handlerErr); ok && e.Public() {
				w.WriteHeader(e.HTTPStatusCode())
				w.Write([]byte(`{"code":"` + e.Code() + `", "message":"` + err.Error() + `"}`))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code":"internalError"}`))

			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(response)

		log.Printf("Success: %+v\n", string(response))
	}
}
