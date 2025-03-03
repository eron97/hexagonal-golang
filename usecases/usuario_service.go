package usecases

import (
	"github.com/eron97/testesGo.git/estudos/hexagonal/domain"
	"github.com/eron97/testesGo.git/estudos/hexagonal/ports"
)

type UsuarioServiceImpl struct {
	repositorio ports.UsuarioRepositorio
}

func NewUsuarioService(repo ports.UsuarioRepositorio) *UsuarioServiceImpl {
	return &UsuarioServiceImpl{repositorio: repo}
}

func (s *UsuarioServiceImpl) CriarUsuario(nome string, email string) error {
	usuario := domain.Usuario{Nome: "Mock Name", Email: "mock@email.com"}
	return s.repositorio.Salvar(usuario)
}
