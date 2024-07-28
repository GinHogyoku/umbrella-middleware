package mw

import (
	"log"
	"net/http"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("User-Role")
		if strings.EqualFold(val, roleAdmin) {
			log.Println("log button user detected")
		}

		next.ServeHTTP(w, r)
	})
}
