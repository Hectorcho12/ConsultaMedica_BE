package routes

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*RegistroUsuario ruta para el registro de un nuevo usuario*/
func RegistroUsuario(w http.ResponseWriter, r *http.Request) {

	var m models.Usuario
	err := json.NewDecoder(r.Body).Decode(&m)

	//Inicio Validaciones//
	if err != nil {
		http.Error(w, "Error al recivir data "+err.Error(), 400)
		return
	}

	if len(m.ID) == 0 {
		http.Error(w, "Correo electronico obligatorio", 400)
		return
	}

	if len(m.Password) == 0 {
		http.Error(w, "Password obligatorio", 400)
		return
	}

	if m.FchIngreso.IsZero() {
		http.Error(w, "Fecha de ingreso obligatoria", 400)
		return
	}
	//Final Validaciones//

	//Validacion base de datos y verificacion de si usuario existe//

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

	/////////////////////////////

	dbref := database.Connect(db)
	defer database.Disconnect(dbref)

	//Valida si el usuario ya existe
	exists, errorcheck := m.CheckUsuario(m.ID, dbref)

	if errorcheck != nil {
		http.Error(w, "Error al consultar si el usuario existe "+errorcheck.Error(), 400)
		return
	}

	if exists {
		http.Error(w, "Error al ingresar usuario. El usuario ya existe", 400)
		return
	}
	/////////////////////////////

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	/*Registra el usuario de acceso*/
	insertErr := m.Insert(dbref)
	if insertErr != nil {
		http.Error(w, "Error al Insertar registro de Ususario"+insertErr.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
