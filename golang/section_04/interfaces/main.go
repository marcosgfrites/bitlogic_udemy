package main

import (
	"fmt"

	"./common"
)

func HumanosRespirando(hu common.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.Sexo())
}

func main() {
	Pedro := new(common.Hombre)
	HumanosRespirando(Pedro)

	Maria := new(common.Mujer)
	HumanosRespirando(Maria)
}
