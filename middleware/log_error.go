package middleware

import (
	"fmt"
	"net/http"
)

func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Fprintln(w, err)
			}
		}()

		h.ServeHTTP(w, r)
	})
}