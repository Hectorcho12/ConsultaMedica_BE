package models

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

/*HistorialPacienteJSON modelo objeto JSON enviado para creacion de paciente e historial */
type HistorialPacienteJSON struct {
	Identidad      string    `json:"identidad,notnull"`
	Nombres        string    `json:"nombres,notnull"`
	Genero         string    `json:"genero,notnull"`
	FchNacimiento  time.Time `json:"fch_nacimiento"`
	Telefono       string    `json:"telefono"`
	Correo         string    `json:"correo"`
	Direccion      string    `json:"direccion"`
	NotasP         string    `json:"notasp"`
	Antecedentes   string    `json:"antecedentes"`
	RecetasActivas string    `json:"recetas_activas"`
	AntFamiliares  string    `json:"antecedentes_familiares"`
	NotasH         string    `json:"notash"`
}

/*HistorialPaciente modelo objeto PG-DB enviado para creacion de paciente e historial */
type HistorialPaciente struct {
	Identidad      string    `pg:"identidad,notnull"`
	Nombres        string    `pg:"nombres,notnull"`
	Genero         string    `pg:"genero,notnull"`
	FchNacimiento  time.Time `pg:"fch_nacimiento"`
	Telefono       string    `pg:"telefono"`
	Correo         string    `pg:"correo"`
	Direccion      string    `pg:"direccion"`
	NotasP         string    `pg:"notasp"`
	Antecedentes   string    `pg:"antecedentes"`
	RecetasActivas string    `pg:"recetas_activas"`
	AntFamiliares  string    `pg:"antecedentes_familiares"`
	NotasH         string    `pg:"notash"`
}

/*Paciente modelo tabla pacientes */
type Paciente struct {
	tableName     struct{}  `pg:"paciente"`
	ID            int       `pg:"id"`
	Identidad     string    `pg:"identidad,notnull"`
	Nombres       string    `pg:"nombres,notnull"`
	Genero        string    `pg:"genero,notnull"`
	FchNacimiento time.Time `pg:"fch_nacimiento"`
	Telefono      string    `pg:"telefono"`
	Correo        string    `pg:"correo"`
	Direccion     string    `pg:"direccion"`
	Notas         string    `pg:"notas"`
}

/*Historial modelo tabla Historial */
type Historial struct {
	tableName      struct{} `pg:"historial_medico"`
	Paciente       int      `pg:"paciente,pk,notnull"`
	Antecedentes   string   `pg:"antecedentes"`
	RecetasActivas string   `pg:"recetas_activas"`
	AntFamiliares  string   `pg:"antecedentes_familiares"`
	Notas          string   `pg:"notas"`
}

/*RelPacienteDoctor modelo tabla relacion paciente-doctor */
type RelPacienteDoctor struct {
	tableName struct{} `pg:"rel_paciente_doctor"`
	Doctor    int      `pg:"doctor"`
	Paciente  int      `pg:"paciente"`
}

/*Insert funcion encargada de insertar paciente*/
func (PI *Paciente) Insert(db *pg.DB) (int, error) {
	_, insertErr := db.Model(PI).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error al ingresar un nuevo paciente, ERROR: %v\n", insertErr)
		return PI.ID, insertErr
	}
	log.Printf("Paciente ingresado exitosamente")

	return PI.ID, nil
}

/*InsertHist funcion encargada de insertar paciente*/
func (HI *Historial) InsertHist(db *pg.DB, DoctorID int, PacienteID int) error {
	insertErr := db.Insert(HI)
	if insertErr != nil {
		log.Printf("Error al ingresar el Historial del paciente, ERROR: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Historial ingresado exitosamente")

	var relPacienteDoctor RelPacienteDoctor

	relPacienteDoctor.Paciente = PacienteID
	relPacienteDoctor.Doctor = DoctorID

	_, insertErrRel := db.Model(&relPacienteDoctor).Insert()

	if insertErrRel != nil {
		log.Printf("Error al ingresar un nuevo paciente, ERROR: %v\n", insertErrRel)
		return insertErrRel
	}
	return nil
}

/*Update funcion encargada de modificar paciente*/
func (PI *Paciente) Update(db *pg.DB) error {
	updateErr := db.Update(PI)
	if updateErr != nil {
		log.Printf("Error al modificar un paciente, ERROR: %v\n", updateErr)
		return updateErr
	}
	log.Printf("Paciente modificado exitosamente")
	return nil
}

/*UpdateHist funcion encargada de modificar paciente*/
func (HI *Historial) UpdateHist(db *pg.DB) error {
	updateErr := db.Update(HI)
	if updateErr != nil {
		log.Printf("Error al modificar un Historial, ERROR: %v\n", updateErr)
		return updateErr
	}
	log.Printf("Historial modificado exitosamente")
	return nil
}

/*GetByID funcion encargada de obtener los datos del paciente por su clave primaria*/
func (PI *Paciente) GetByID(db *pg.DB, ID string) (PIreturn *Paciente, errorReturn error) {
	selectError := db.Model(PI).Column("id",
		"identidad",
		"nombres",
		"genero",
		"fch_nacimiento",
		"telefono",
		"correo",
		"direccion",
		"notas").Where("id = ?", ID).Select()

	if selectError != nil {
		log.Printf("Error al obtener datos del paciente "+ID+", ERROR: %v\n", selectError)
		return PI, selectError
	}

	return PI, nil

}

/*GetByIDPacient funcion encargada de obtener los datos del paciente por su identificacion*/
func (PI *Paciente) GetByIDPacient(db *pg.DB, pacienteID string) (PIreturn *Paciente, errorReturn error) {
	selectError := db.Model(PI).Where("identidad = ?", pacienteID).Select()

	if selectError != nil {
		log.Printf("Error al obtener datos del paciente "+pacienteID+", ERROR: %v\n", selectError)
		return PI, selectError
	}

	return PI, nil
}

/*GetByName funcion encargada de obtener los datos del paciente por su Nombre*/
func (PI *Paciente) GetByName(db *pg.DB, pacienteNombre string, DoctorID int) (PIreturn []Paciente, errorReturn error) {

	var arrayP []Paciente
	//	selectError := db.Model(&arrayP).Where("nombres LIKE ?", "%"+pacienteNombre+"%").Where().Select()
	_, selectError := db.Query(&arrayP, `
	select 
	a.id,
	a.identidad,
	a.nombres,
	a.genero,
	a.fch_nacimiento,
	a.telefono,
	a.correo,
	a.direccion,
	a.notas
	From historial_medico b inner join paciente a on a.id = b.paciente 
	inner join rel_paciente_doctor c on a.id = c.paciente
	inner join doctor d on c.doctor = d.id
	where d.id = ?
	and a.nombres iLIKE ? 
	order by a.id ASC`, DoctorID, "%"+pacienteNombre+"%")

	if selectError != nil {
		log.Printf("Error al obtener datos del paciente "+pacienteNombre+", ERROR: %v\n", selectError)
		return arrayP, selectError
	}

	return arrayP, nil
}

/*GetAll funcion encargada de obtener todos los pacientes*/
func (PI *Paciente) GetAll(db *pg.DB, DoctorID int, Limit int, Ofset int) (PIreturn []Paciente, errorReturn error) {
	var arrayP []Paciente

	_, selectError := db.Query(&arrayP, `
	select 
	a.id,
	a.identidad,
	a.nombres,
	a.genero,
	a.fch_nacimiento,
	a.telefono
	From historial_medico b inner join paciente a on a.id = b.paciente 
	inner join rel_paciente_doctor c on a.id = c.paciente
	inner join doctor d on c.doctor = d.id
	where d.id = ?
	order by a.id ASC
	limit ? OFFSET ?`, DoctorID, Limit, Ofset)

	if selectError != nil {
		log.Printf("Error al obtener la lista de pacientes, ERROR: %v\n", selectError)
		return arrayP, selectError
	}

	return arrayP, nil
}

/*GetPacienteData funcion encargada de obtener la data del paciente junto a su historial medico*/
func (PI *Paciente) GetPacienteData(db *pg.DB, idPaciente string) (PIreturn HistorialPaciente, errorReturn error) {

	var PHI HistorialPaciente
	_, err := db.QueryOne(&PHI, `
				SELECT p.identidad, 
				P.nombres,
				P.genero,
				P.fch_nacimiento, 
				P.telefono,
				P.correo,
				P.direccion,
				P.notas as notasp,
				a.antecedentes, 
				a.recetas_activas, 
				a.antecedentes_familiares, 
				a.notas as notash
				FROM paciente p inner join historial_medico a ON p.id = a.paciente
				WHERE p.id = ? 
			`, idPaciente)

	log.Printf(": %v\n", idPaciente)
	log.Printf(": %v\n", PHI)

	if err == nil {
		log.Printf("Error al obtener los datos del paciente, ERROR: %v\n", err)
		return PHI, err
	}

	return PHI, nil
}
