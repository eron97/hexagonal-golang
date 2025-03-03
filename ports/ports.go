package ports

import "github.com/eron97/testesGo.git/estudos/hexagonal/domain"

// Define as interfaces que serão utilizadas
type UsuarioService interface {
	CriarUsuario(nome string, email string) error
}

type UsuarioRepositorio interface {
	Salvar(usuario domain.Usuario) error
}
