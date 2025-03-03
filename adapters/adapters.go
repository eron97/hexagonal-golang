package adapters

import (
	"fmt"

	"github.com/eron97/testesGo.git/estudos/hexagonal/domain"
)

// Implementação MySQL
type UsuarioRepositorioMySQL struct {
	// Em um caso real, teríamos a conexão do MySQL aqui
	connectionString string
}

func NewMySQLRepository(connString string) *UsuarioRepositorioMySQL {
	return &UsuarioRepositorioMySQL{
		connectionString: connString,
	}
}

func (r *UsuarioRepositorioMySQL) Salvar(usuario domain.Usuario) error {
	// Simulando inserção no MySQL
	fmt.Printf("[MySQL] Executando query: INSERT INTO usuarios (nome, email) VALUES ('%s', '%s')\n",
		usuario.Nome, usuario.Email)
	return nil
}

// Implementação SQLServer
type UsuarioRepositorioSQLServer struct {
	// Em um caso real, teríamos a conexão do SQL Server aqui
	connectionString string
}

func NewSQLServerRepository(connString string) *UsuarioRepositorioSQLServer {
	return &UsuarioRepositorioSQLServer{
		connectionString: connString,
	}
}

func (r *UsuarioRepositorioSQLServer) Salvar(usuario domain.Usuario) error {
	// Simulando inserção no SQL Server
	fmt.Printf("[SQLServer] Executando query: INSERT INTO dbo.usuarios (nome, email) VALUES ('%s', '%s')\n",
		usuario.Nome, usuario.Email)
	return nil
}
