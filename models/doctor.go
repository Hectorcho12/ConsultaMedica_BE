package models

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

/*Doctor estructura para el objeto Doctor
Notas

*/
type Doctor struct {
	tableName     struct{}  `pg:"doctor"`
	ID            int       `pg:"id"`
	Nombre        string    `pg:"nombre,notnull"`
	FchNacimiento time.Time `pg:"fch_nacimiento,"`
	Especialidad  int       `pg:"especialidad"`
	Pais          string    `pg:"pais"`
	Ciudad        string    `pg:"ciudad"`
	FchRegistro   time.Time `pg:"fch_registro,"`
	Estado        bool      `pg:"estado,use_zero,notnull"`
	Correo        string    `pg:"correo,notnull"`
}

/*DoctorID estructura para obtener el id del doctor
Notas

*/
type DoctorID struct {
	tableName struct{} `pg:"doctor"`
	ID        int      `pg:"id,notnull"`
}

/*Insert estructura para el objeto Doctor
Notas

*/
func (DI *Doctor) Insert(db *pg.DB) error {
	insertErr := db.Insert(DI)
	if insertErr != nil {
		log.Printf("Error al ingresar un nuevo Doctor, ERROR: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Dcotor ingresado exitosamente")
	return nil
}

/*Update metodo de update para la estructura doctor
Notas

*/
func (DI *Doctor) Update(db *pg.DB) error {
	updateErr := db.Update(DI)
	if updateErr != nil {
		log.Printf("Error al modificar un Doctor, ERROR: %v\n", updateErr)
		return updateErr
	}
	log.Printf("Dcotor modificado exitosamente")
	return nil
}

/*GetID metodo que obtiene el id del doctor en base al correo
Notas

*/
func (DI *DoctorID) GetID(email string, db *pg.DB) (int, error) {

	SelectError := db.Model(DI).Column("id").Where("correo = ?", email).Select()

	if SelectError != nil {
		log.Printf("Error al consultar id doctor, ERROR: %v\n", SelectError)
		return 0, SelectError
	}

	return DI.ID, SelectError
}

/*GetByID metodo que obtiene la informacion del doctor por su ID*/
func (DI *Doctor) GetByID(db *pg.DB, ID string) (DIreturn *Doctor, errorReturn error) {
	selectError := db.Model(DI).Where("id = ?", ID).Select()

	if selectError != nil {
		log.Printf("Error al obtener datos del paciente "+ID+", ERROR: %v\n", selectError)
		return DI, selectError
	}

	return DI, nil
}
