package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
)

/*RegistroPaciente funcion para registrar un paciente */
func RegistroPaciente(w http.ResponseWriter, r *http.Request) {

	//user := r.Header.Get("user")

	var hp models.HistorialPacienteJSON
	var h models.Historial
	var m models.Paciente

	err := json.NewDecoder(r.Body).Decode(&hp)

	if err != nil {
		http.Error(w, "Error al recivir data "+err.Error(), 400)
		return
	}

	if len(hp.Nombres) == 0 {
		http.Error(w, "Los nombres del paciente son requeridos", 400)
		return
	}

	if len(hp.Genero) == 0 {
		http.Error(w, "EL genero es requerido", 400)
		return
	}

	//Parse valores de tabla paciente a objeto paciente
	m.Identidad = hp.Identidad
	m.Nombres = hp.Nombres
	m.Genero = hp.Genero
	m.FchNacimiento = hp.FchNacimiento
	m.Telefono = hp.Telefono
	m.Correo = hp.Correo
	m.Direccion = hp.Direccion
	m.Notas = hp.NotasP

	//dbref := database.Connect(user)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	estatus, errdb := database.CheckConnecion(dbref)
	if estatus == false {
		http.Error(w, "Error al conetar con base de datos"+errdb.Error(), 400)
		return
	}

	idPaciente, insertErr := m.Insert(dbref)
	if insertErr != nil {
		http.Error(w, "Error al Insertar registro de paciente"+insertErr.Error(), 400)
		return
	}

	//Parse valores de tabla historial a objeto historial
	h.Paciente = idPaciente
	h.Antecedentes = hp.Antecedentes
	h.RecetasActivas = hp.RecetasActivas
	h.AntFamiliares = hp.AntFamiliares
	h.Notas = hp.NotasH

	insertErr = h.InsertHist(dbref, DoctorID, idPaciente)
	if insertErr != nil {
		http.Error(w, "Error al Insertar el historial de paciente"+insertErr.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
