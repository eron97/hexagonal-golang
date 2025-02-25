package ports

import "github.com/eron97/testesGo.git/estudos/hexagonal/domain"

type UsuarioService interface {
	CriarUsuario(nome string, email string) error
}

type UsuarioRepositorio interface {
	Salvar(usuario domain.Usuario) error
}
