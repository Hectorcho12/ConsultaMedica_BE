package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Hectorcho12/ConsultaMedica_BE/middleware"
	"github.com/Hectorcho12/ConsultaMedica_BE/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores para backend*/
func Manejadores() {

	router := mux.NewRouter()

	/*Rutas pacientes*/
	router.HandleFunc("/registroPaciente", middleware.CheckJWT(routes.RegistroPaciente)).Methods("POST")
	router.HandleFunc("/verPaciente", middleware.CheckJWT(routes.VerPaciente)).Methods("GET")
	router.HandleFunc("/modificarPaciente", middleware.CheckJWT(routes.ModificarPaciente)).Methods("PUT")
	/*Rutas Doctores*/
	router.HandleFunc("/registroUsuario", routes.RegistroUsuario).Methods("POST")
	router.HandleFunc("/registroDoctor", routes.RegistroDoctor).Methods("POST")
	router.HandleFunc("/verDoctor", middleware.CheckJWT(routes.VerDoctor)).Methods("GET")
	router.HandleFunc("/modificarDoctor", middleware.CheckJWT(routes.ModificarDoctor)).Methods("PUT")
	/*Rutas login*/
	router.HandleFunc("/login", routes.Login).Methods("POST")
	/*Rutas consulta*/
	router.HandleFunc("/creaConsulta", middleware.CheckJWT(routes.CreaConsulta)).Methods("POST")
	router.HandleFunc("/verConsulta", middleware.CheckJWT(routes.VerConsulta)).Methods("GET")
	router.HandleFunc("/verPacienteData", middleware.CheckJWT(routes.VerPacienteData)).Methods("GET")
	router.HandleFunc("/modificarConsulta", middleware.CheckJWT(routes.ModificarConsulta)).Methods("PUT")
	/*Rutas Historial*/
	router.HandleFunc("/modificarHistorial", middleware.CheckJWT(routes.ModificarHistorial)).Methods("PUT")
	/*Receta medica*/
	//generaReceta

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
