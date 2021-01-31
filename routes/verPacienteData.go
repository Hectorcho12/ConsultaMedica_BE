package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*VerPacienteData metodo para ver los datos del paciente*/
func VerPacienteData(w http.ResponseWriter, r *http.Request) {

	paciente := r.URL.Query().Get("paciente")

	var m *models.Paciente
	var resultm models.HistorialPaciente

	//dbref := database.Connect(UserID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	resultm, errorSelect := m.GetPacienteData(dbref, paciente)

	if errorSelect != nil {
		http.Error(w, "Error al obtener datos de consultas. ERROR: "+errorSelect.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultm)
	return

}
