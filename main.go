package main

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/handlers"
)

func main() {
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

	log.Printf("Listening on %v\n", "localhost:8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatalln(err.Error())
}

// func initStorage(db *sql.DB) {
// 	err := db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to Database!")
// }
