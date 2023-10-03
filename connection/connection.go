package connection

import (
	"database/sql"
	"fmt"
	"log"
	"net/rpc"

	"github.com/PyMarcus/rpc_chat/repository"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectionSQL() (*sql.DB, error) {
	path := "sqlite.db"
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Fatal(err)
	}
	//setupDb(db)
	// r := []*models.Room{{Name: "Jogos"}, { Name: "Noticias"}, {Name: "Dinheiro"},
	//  {Name: "T.I"}, { Name: "Filmes"}, {Name: "Musicas"}, {Name: "Namoro"},
	// {Name: "Amizade"}, { Name: "Ciencia"}, { Name: "Viagem"}, { Name: "Livros"}}
	// repository.NewRepository(db).InsertRoomsIntoDB(r)
	return db, nil
}

// setupDb setting the database
func setupDb(sqlDB *sql.DB) {
	r := repository.NewRepository(sqlDB)

	err := r.Migrate()

	if err != nil {
		log.Fatal(err)
	}
}

func ServerRPCConnection(ip, port string) *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%s", ip, port))

	if err != nil {
		panic(err)
	}
	return client
}
