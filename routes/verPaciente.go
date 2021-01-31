package routes

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*VerPaciente ruta para visualizar los datos del paciente*/
func VerPaciente(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	pacienteID := r.URL.Query().Get("pacienteid")

	nombres := r.URL.Query().Get("nombres")
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	ofset, _ := strconv.Atoi(r.URL.Query().Get("Ofset"))

	var m models.Paciente
	var arraym []models.Paciente

	//dbref := database.Connect(UserID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	if len(ID) != 0 {

		m, err := m.GetByID(dbref, ID)

		if err != nil {
			http.Error(w, "Error al consultar paciente "+err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(m)
		return

	}

	if len(pacienteID) != 0 {

		m, err := m.GetByIDPacient(dbref, pacienteID)

		if err != nil {
			http.Error(w, "Error al consultar paciente "+err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(m)
		return
	}

	if len(nombres) != 0 {

		var arrayNombres []models.Paciente

		arrayNombres, err := m.GetByName(dbref, nombres, DoctorID)

		if err != nil {
			http.Error(w, "Error al consultar paciente "+err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(arrayNombres)
		return

	}

	var errorarray error
	arraym, errorarray = m.GetAll(dbref, DoctorID, limit, ofset)

	if errorarray != nil {
		http.Error(w, "Error al consultar paciente "+errorarray.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(arraym)

}
