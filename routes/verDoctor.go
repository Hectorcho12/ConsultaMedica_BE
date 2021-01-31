package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*VerDoctor metodo para ver los datos del doctor*/
func VerDoctor(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var m *models.Doctor

	//dbref := database.Connect(UserID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	m, errorSelect := m.GetByID(dbref, ID)

	if errorSelect != nil {
		http.Error(w, "Error al modificar datos del paciente. ERROR: "+errorSelect.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
	return

}
