package services

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
	usuario := domain.Usuario{Nome: nome, Email: email}
	return s.repositorio.Salvar(usuario)
}
