package common

type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}

type Animal interface {
	Respirar()
	Comer()
	EsCarnivoro() bool
}

type Vegetal interface {
	ClasificacionVegetal() string
}

type Hombre struct {
	Edad       int
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
}

type Mujer struct {
	Edad       int
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
}

func (h *Hombre) Respirar() {
	h.Respirando = true
}

func (h *Hombre) Comer() {
	h.Comiendo = true
}

func (h *Hombre) Pensar() {
	h.Pensando = true
}

func (h *Hombre) Sexo() string {
	return "Hombre"
}

func (m *Mujer) Respirar() {
	m.Respirando = true
}

func (m *Mujer) Comer() {
	m.Comiendo = true
}

func (m *Mujer) Pensar() {
	m.Pensando = true
}

func (m *Mujer) Sexo() string {
	return "Mujer"
}
