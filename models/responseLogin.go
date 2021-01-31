package models

/*ResponseLogin modelo usado para la repsuesta del token*/
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
