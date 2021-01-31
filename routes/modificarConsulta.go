package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*ModificarConsulta metodo para modificar consulta a paciente*/
func ModificarConsulta(w http.ResponseWriter, r *http.Request) {
	var m models.Consulta
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "Error al recivir data "+err.Error(), 400)
		return
	}

	if m.Paciente == 0 {
		http.Error(w, "Error Campos incompletos: Campo Paciente vacio", 400)
		return
	}

	if m.FchConsulta.IsZero() {
		http.Error(w, "Error Campos incompletos: Campo fecha vacio o con formato inadecuado", 400)
		return
	}

	m.Doctor = DoctorID

	//dbref := database.Connect(UserID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	status, errdb := database.CheckConnecion(dbref)
	if status == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	errorConsulta := m.Update(dbref)

	if errorConsulta != nil {
		http.Error(w, "Error al modificar consulta. ERROR: "+errorConsulta.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
