package main

import (
	"github.com/eron97/testesGo.git/estudos/hexagonal/adapters"
	"github.com/eron97/testesGo.git/estudos/hexagonal/services"
)

func main() {
	repositorio := &adapters.UsuarioRepositorioSQL{}
	servico := services.NewUsuarioService(repositorio)

	servico.CriarUsuario("UserMock", "mock@email.com")
}
