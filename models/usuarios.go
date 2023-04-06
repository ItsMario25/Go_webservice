package models

type Usuario struct {
	ID       int
	Nombre   string
	Email    string
	Password string
}

type logg struct {
	Email    string
	Password string
}
