package middleware

import (
	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/routes"
)

/*CheckJWT funcion encargada de validar el JWT*/
func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token. Error: "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
