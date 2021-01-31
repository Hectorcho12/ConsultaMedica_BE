package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*VerConsulta metodo para ver consulta a paciente(s)*/
func VerConsulta(w http.ResponseWriter, r *http.Request) {

	paciente := r.URL.Query().Get("paciente")

	var m models.Consulta

	//dbref := database.Connect(UserID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	if len(paciente) != 0 {

		arrayC, errorSelect := m.GetByPaciente(dbref, paciente)

		if errorSelect != nil {
			http.Error(w, "Error al obtener datos de consultas. ERROR: "+errorSelect.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(arrayC)
		return
	}

	arrayC, errorSelect := m.GetByDoctor(dbref, strconv.Itoa(DoctorID))

	if errorSelect != nil {
		http.Error(w, "Error al obtener datos de consultas. ERROR: "+errorSelect.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(arrayC)
	return

}
