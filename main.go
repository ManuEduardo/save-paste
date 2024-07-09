package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ManuEduardo/save-paste/src/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	portServer := os.Getenv("PORT")

	// cfg := mysql.Config{
	// 	User:                 store.Envs.DBUser,
	// 	Passwd:               store.Envs.DBPassword,
	// 	Addr:                 store.Envs.DBAddress,
	// 	DBName:               store.Envs.DBName,
	// 	Net:                  "tcp",
	// 	AllowNativePasswords: true,
	// 	ParseTime:            true,
	// }

	// db, err := store.NewMySQLStorage(cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// store := store.NewStore(db)

	// initStorage(db)

	router := http.NewServeMux()

	// serve files in public
	fileServer := http.FileServer(http.Dir("src/public"))
	router.Handle("GET /public/", http.StripPrefix("/public/", fileServer))

	router.HandleFunc("GET /login", handlers.HandleLogin)
	router.HandleFunc("GET /", handlers.HandleBase)
	// router.HandleFunc("/cars", handler.HandleListCars)
	// router.HandleFunc("/cars", handler.HandleAddCar)
	// router.HandleFunc("/cars/{id}", handler.HandleDeleteCar)
	// router.HandleFunc("/cars/search", handler.HandleSearchCar)

	log.Printf("Listening on %v\n", fmt.Sprintf("localhost:%v", portServer))
	err = http.ListenAndServe(fmt.Sprintf(":%v", portServer), router)
	log.Fatalln(err.Error())
}

// func initStorage(db *sql.DB) {
// 	err := db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to Database!")
// }
