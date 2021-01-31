package routes

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/jwt"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*Login funcion encargada de validar el ingreso a la app*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var m models.Usuario

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(m.ID) == 0 {
		http.Error(w, "Email es un campo requerido ", 400)
		return
	}

	if len(m.Password) == 0 {
		http.Error(w, "Password es un campo requerido ", 400)
		return
	}

	var db string
	properties, err := os.Open("" + os.Getenv("CONFIG") + "/climedi/climedi.properties")
	if err != nil {
		log.Println("Error al leer archivo de configuraciones")
	} else {
		scanner := bufio.NewScanner(properties)
		for scanner.Scan() {
			linea := scanner.Text()
			if strings.HasPrefix(linea, "userdb") {
				db = linea[7:]
			}
		}
	}

	dbref := database.Connect(db)
	defer database.Disconnect(dbref)

	exists, codigoDoc := m.Login(m.ID, m.Password, dbref)

	if exists == false {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtkey, err := jwt.GenerateJWT(m, codigoDoc)

	if err != nil {
		http.Error(w, "Error al generar autenticacion "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtkey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
