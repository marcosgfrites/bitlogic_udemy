package main

import (
	"fmt"
	"time"

	us "./user"
)

type Persona struct {
	us.Usuario
}

func main() {
	user := new(Persona)
	user.AltaUsuario(1, "marcosgfd", time.Now(), true)
	fmt.Println(user)
}
