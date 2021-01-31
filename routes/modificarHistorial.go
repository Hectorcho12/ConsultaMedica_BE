package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*ModificarHistorial metodo para modificar el historial de X paciente*/
func ModificarHistorial(w http.ResponseWriter, r *http.Request) {

	var m models.Historial
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "Error al recivir data "+err.Error(), 400)
		return
	}

	if m.Paciente == 0 {
		http.Error(w, "El paciente es una campo obligatorio", 400)
		return
	}

	//dbref := database.Connect(UserID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	errorUpdate := m.UpdateHist(dbref)

	if errorUpdate != nil {
		http.Error(w, "Error al modificar datos del paciente. ERROR: "+errorUpdate.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
