package routes

import (
	"encoding/json"

	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*ModificarDoctor metodo para modificar el doctor*/
func ModificarDoctor(w http.ResponseWriter, r *http.Request) {

	var m models.Doctor
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "Error al recivir data "+err.Error(), 400)
		return
	}

	if m.ID == 0 {
		http.Error(w, "ID es una campo obligatorio", 400)
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

	errorUpdate := m.Update(dbref)

	if errorUpdate != nil {
		http.Error(w, "Error al modificar datos del doctor. ERROR: "+errorUpdate.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
