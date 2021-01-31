package database

import "golang.org/x/crypto/bcrypt"

/*Encriptar Funcion encargada de encriptar una password*/
func Encriptar(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
