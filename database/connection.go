package database

import (
	"bufio"
	"log"
	"os"
	"strings"

	pg "github.com/go-pg/pg"
)

/*Connect Clase conexion a base de datos
- Devuelve un objeto de referencia sobre la conexion
user string, pass string, host string, dbname string
*/
func Connect(dbname string) *pg.DB {

	var _user string
	var _pass string
	var _addr string

	properties, err := os.Open("" + os.Getenv("CONFIG") + "/climedi/climedi.properties")
	if err != nil {
		log.Println("Error al leer archivo de configuraciones")
	} else {
		scanner := bufio.NewScanner(properties)
		for scanner.Scan() {
			linea := scanner.Text()
			if strings.HasPrefix(linea, "loginuser") {
				_user = linea[10:]
			}
			if strings.HasPrefix(linea, "loginpassword") {
				_pass = linea[14:]
			}
			if strings.HasPrefix(linea, "host") {
				_addr = linea[5:]
			}
		}
	}
	properties.Close()

	opts := &pg.Options{
		User:     _user,
		Password: _pass,
		Addr:     _addr,
		Database: dbname}

	db := pg.Connect(opts)

	if db == nil {
		log.Printf("Error al conectar a base de datos")
		os.Exit(100)
	}
	log.Printf("Conectado")
	return db
}

/*Disconnect Clase encargada de desconectar la base de datos
- db *pg.DB: Objeto de referencia creado por el metodo Connect
*/
func Disconnect(db *pg.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("Error al cerrar conexion a base de datos %v\n", err)
	}
	log.Println("Conexion a base de datos cerrada")
}

/*CheckConnecion revisa si la conexion sigue activa*/
func CheckConnecion(db *pg.DB) (bool, error) {
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Printf("Conexion a base de datos invalida, Revisar: %v\n", err)
		return false, err
	}
	log.Println("Conexion a base de datos valida")
	return true, err

}
