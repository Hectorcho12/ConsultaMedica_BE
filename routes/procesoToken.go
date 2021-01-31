package routes

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/Hectorcho12/ConsultaMedica_BE/database"
	"github.com/Hectorcho12/ConsultaMedica_BE/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*UserID Variable global para ID de usuario*/
var UserID string

/*DoctorID variable global para el id del doctor*/
var DoctorID int

/*UserMaster Variable global para validar si el usuario es master*/
var UserMaster bool

/*UserChild Variable global para validar si el usuario es child*/
var UserChild bool

/*UserTipoPlan Variable global para validar el plan del usuario*/
var UserTipoPlan int

/*ProcessToken metodo encargado de procesar el token*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("ClimediKey2020")
	claims := &models.Claim{}
	var usuario models.Usuario

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de Token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	var db string
	properties, err := os.Open("" + os.Getenv("CONFIG") + "/climedi/climedi.properties")
	if err != nil {
		log.Println("Error al leer archivo de configuraciones")
	} else {
		scanner := bufio.NewScanner(properties)
		for scanner.Scan() {
			linea := scanner.Text()
			if strings.HasPrefix(linea, "userdb") {
				db = linea[7:]
			}
		}
	}

	dbref := database.Connect(db)
	defer database.Disconnect(dbref)

	if err == nil {
		exists, _ := usuario.CheckUsuario(claims.ID, dbref)

		if exists {
			UserID = claims.ID
			DoctorID = claims.IDDoc
			UserChild = claims.Child
			UserMaster = claims.Master
			UserTipoPlan = claims.TipoPlan
		}
		return claims, exists, UserID, nil
	}

	return claims, false, string(""), err
}
