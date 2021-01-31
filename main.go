package main

import (
	"github.com/Hectorcho12/ConsultaMedica_BE/handlers"
)

//"github.com/Hectorcho12/ConsultaMedica_BE/handlers"

func main() {

	handlers.Manejadores()

}

/*InPaciente bla bla bla
func InPaciente(dbref *pg.DB) {

	newPaciente := &models.Paciente{
		Nombres:       "Hector David Rodriguez Aguilar",
		Genero:        "F",
		FchNacimiento: time.Now(),
		Telefono:      "87961905",
		Correo:        "hectorcho12@gmail.com",
		Notas:         "PRUEBA",
	}

	newPaciente.Insert(dbref)

}
*/
