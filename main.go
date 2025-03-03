package main

import (
	"fmt"

	"github.com/eron97/testesGo.git/estudos/hexagonal/adapters"
	"github.com/eron97/testesGo.git/estudos/hexagonal/services"
)

func main() {
	// Configuração MySQL
	mysqlRepo := adapters.NewMySQLRepository("mysql://localhost:3306/db")
	servicoMySQL := services.NewUsuarioService(mysqlRepo)

	// Criando usuário usando MySQL
	fmt.Println("=== Criando usuário via MySQL ===")
	servicoMySQL.CriarUsuario("João", "joao@email.com")

	// Configuração SQL Server
	sqlServerRepo := adapters.NewSQLServerRepository("sqlserver://localhost:1433/db")
	servicoSQLServer := services.NewUsuarioService(sqlServerRepo)

	// Criando usuário usando SQL Server
	fmt.Println("\n=== Criando usuário via SQL Server ===")
	servicoSQLServer.CriarUsuario("Maria", "maria@email.com")
}
