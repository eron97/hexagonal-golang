package domain

type Usuario struct {
	Nome  string
	Email string
}

type UsuarioRepositorio interface {
	Salvar(usuario Usuario) error
}
