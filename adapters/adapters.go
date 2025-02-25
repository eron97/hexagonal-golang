package adapters

import (
	"fmt"

	"github.com/eron97/testesGo.git/estudos/hexagonal/domain"
)

type UsuarioRepositorioSQL struct{}

func (r *UsuarioRepositorioSQL) Salvar(usuario domain.Usuario) error {
	fmt.Printf("Salvando usuário %s no banco SQL\n", usuario.Nome)
	return nil
}
