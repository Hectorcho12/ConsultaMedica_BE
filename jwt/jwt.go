package jwt

import (
	"time"

	"github.com/Hectorcho12/ConsultaMedica_BE/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerateJWT metodo encargado de generar Json Web Token*/
func GenerateJWT(m models.Usuario, CodigoDoc int) (string, error) {
	myKey := []byte("ClimediKey2020")

	payload := jwt.MapClaims{
		"ID":       m.ID,
		"IDDoc":    CodigoDoc,
		"Master":   m.Master,
		"Status":   m.Status,
		"Child":    m.Child,
		"TipoPlan": m.TipoPlan,
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
