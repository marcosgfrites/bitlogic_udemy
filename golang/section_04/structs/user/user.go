package user

import (
	"time"
)

type Usuario struct {
	Id        int
	Username  string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, username string, fechaAlta time.Time, status bool) {
	this.Id = id
	this.Username = username
	this.FechaAlta = fechaAlta
	this.Status = status
}
