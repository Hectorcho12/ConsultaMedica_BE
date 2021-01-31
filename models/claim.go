package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/*Claim structura para los claims de jwt*/
type Claim struct {
	ID       string `json:"ID"`
	IDDoc    int    `json:"IDDoc"`
	Master   bool   `json:"Master"`
	Status   bool   `json:"Status"`
	TipoPlan int    `json:"Tipoplan"`
	Child    bool   `json:"Child"`
	jwt.StandardClaims
}
