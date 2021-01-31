package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*RegistroDoctor ruta para ingresar un nuevo doctor*/
func RegistroDoctor(w http.ResponseWriter, r *http.Request) {
	var m models.Doctor
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "Error al recivir data "+err.Error(), 400)
		return
	}

	if len(m.Nombre) == 0 {
		http.Error(w, "El nombre es obligatorio ", 400)
		return
	}

	if len(m.Correo) == 0 {
		http.Error(w, "El correo es obligatorio ", 400)
		return
	}

	dbref := database.Connect("climediprueba")

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	/*Registra el usuario de acceso*/
	insertErr := m.Insert(dbref)
	if insertErr != nil {
		http.Error(w, "Error al Insertar registro de Doctor"+insertErr.Error(), 400)
		return
	}

	database.Disconnect(dbref)

	w.WriteHeader(http.StatusCreated)

}
