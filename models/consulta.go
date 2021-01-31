package models

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

/*Consulta estructura del objecto consulta*/
type Consulta struct {
	tableName         struct{}  `pg:"consulta"`
	ID                int       `pg:"id,pk"`
	Paciente          int       `pg:"paciente,notnull"`
	FchConsulta       time.Time `pg:"fch_consulta,"`
	Razon             string    `pg:"razon"`
	Sintomas          string    `pg:"sintomas"`
	Procedimientos    string    `pg:"procedimientos"`
	Diagnstico        string    `pg:"diagnostico"`
	Notas             string    `pg:"notas"`
	Peso              float32   `pg:"peso"`
	Altura            float32   `pg:"altura"`
	PresionSistolica  int8      `pg:"presion_sistolica"`
	PresionDiastolica int8      `pg:"presion_diastolica"`
	Temperatura       int8      `pg:"temperatura"`
	MasaCorporal      float32   `pg:"masa_corporal"`
	Doctor            int       `pg:"doctor,notnull"`
	GeneroReceta      bool      `pg:"genero_receta"`
	Medicamentos      string    `pg:"medicamentos"`
}

/*Insert estructura para el objeto consulta
Notas

*/
func (CI *Consulta) Insert(db *pg.DB) error {
	insertErr := db.Insert(CI)
	if insertErr != nil {
		log.Printf("Error al ingresar una nueva consulta, ERROR: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Consulta ingresada exitosamente")
	return nil
}

/*Update metodo para actualizar consultas
Notas

*/
func (CI *Consulta) Update(db *pg.DB) error {
	UpdateErr := db.Update(CI)
	if UpdateErr != nil {
		log.Printf("Error al modificar una consulta, ERROR: %v\n", UpdateErr)
		return UpdateErr
	}
	log.Printf("Consulta modificada exitosamente")
	return nil
}

/*GetByPaciente para obtener las conusltas por paciente
Notas

*/
func (CI *Consulta) GetByPaciente(db *pg.DB, IDpaciente string) ([]Consulta, error) {
	var arrayC []Consulta

	selectError := db.Model(&arrayC).Where("paciente = ?", IDpaciente).Select()

	if selectError != nil {
		log.Printf("Error al obtener datos de consulta para paciente "+IDpaciente+", ERROR: %v\n", selectError)
		return arrayC, selectError
	}

	return arrayC, nil
}

/*GetByDoctor para obtener las conusltas por los pacientes que son atendidos por un doctor
Notas

*/
func (CI *Consulta) GetByDoctor(db *pg.DB, IDdoctor string) ([]Consulta, error) {

	var arrayC []Consulta

	selectError := db.Model(&arrayC).Where("doctor = ?", IDdoctor).Select()

	if selectError != nil {
		log.Printf("Error al obtener datos de consulta para doctor "+IDdoctor+", ERROR: %v\n", selectError)
		return arrayC, selectError
	}

	return arrayC, nil
}
