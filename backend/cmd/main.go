package main

import (
	"log"

	"github.com/kamdyns/movie-chat/db"
	"github.com/kamdyns/movie-chat/internal/user"
	"github.com/kamdyns/movie-chat/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userServ := user.NewService(userRep)
	userHandler := user.NewHandler(userServ)

	router.InitRouter(userHandler)
	router.Start("localhost:8080", router.InitRouter(userHandler))
}
