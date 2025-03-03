package domain

// Define a estrutura central do negócio
type Usuario struct {
	Nome  string
	Email string
}

type UsuarioRepositorio interface {
	Salvar(usuario Usuario) error
}
