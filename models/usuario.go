package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/go-pg/pg"
)

/*Usuario structura del usuario en base de datos
Notas
- Tipo plan tipo integer posee los siguientes valores
	* 1 = plan basico
	* 2 = plan plus
*/
type Usuario struct {
	tableName  struct{}  `pg:"users"`
	ID         string    `pg:"id,pk,notnull"`
	Password   string    `pg:"password,notnull"`
	Master     bool      `pg:"master,use_zero,notnull"`
	Status     bool      `pg:"status,use_zero,notnull"`
	FchIngreso time.Time `pg:"fch_ingreso,notnull"`
	Child      bool      `pg:"child,use_zero,notnull"`
	TipoPlan   int       `pg:"tipo_plan,use_zero,notnull"`
}

/*Insert funcion encargada de insertar Usuario*/
func (UI *Usuario) Insert(db *pg.DB) error {

	var errorpass error
	UI.Password, errorpass = database.Encriptar(UI.Password)
	if errorpass != nil {
		log.Printf("Error al encriptar password, ERROR: %v\n", errorpass)
	}

	insertErr := db.Insert(UI)
	if insertErr != nil {
		log.Printf("Error al ingresar un nuevo Usuario, ERROR: %v\n", insertErr)
		return insertErr
	}

	log.Printf("Usuario ingresado exitosamente")
	return nil
}

/*CheckUsuario valida si el usuario existe*/
func (UI *Usuario) CheckUsuario(usuario string, db *pg.DB) (exists bool, err error) {

	var count int
	_, errorSelect := db.Model((*Usuario)(nil)).QueryOne(pg.Scan(&count), `SELECT COUNT(*) FROM USERS WHERE ID = ?`, usuario)

	//Usuario si existe
	if count > 0 {
		return true, errorSelect
	}
	//Usuario no existe
	return false, errorSelect

}

/*Login funcion encargada de realizar login a la aplicacion*/
func (UI *Usuario) Login(email string, pass string, db *pg.DB) (r bool, doc int) {

	exists, errorExists := UI.CheckUsuario(email, db)
	if errorExists != nil {
		log.Printf("Error al consultar usuario. Error: %v\n", errorExists)
		return false, 0
	}

	if !exists {
		log.Println("Usuario no encontrado")
		return false, 0
	}

	err := db.Model(UI).Column("id", "password", "master", "status", "child", "tipo_plan").Where("id = ?", email).Select()

	if err != nil {
		log.Printf("Error al consultar usuario. Error: %v\n", err)
		return false, 0
	}

	//dbref := database.Connect(UI.ID)
	dbref := database.Connect("climediprueba")
	defer database.Disconnect(dbref)

	var m DoctorID
	codigoDoc, errorGetID := m.GetID(UI.ID, dbref)

	if errorGetID != nil {
		log.Printf("Error al consultar codigo de doctor (Metodo GetID). Error: %v\n", errorGetID)
		return false, 0
	}

	passwordClient := []byte(pass)
	passwordDB := []byte(UI.Password)

	errDecrypt := bcrypt.CompareHashAndPassword(passwordDB, passwordClient)

	if errDecrypt != nil {
		return false, 0
	}

	return true, codigoDoc

}
