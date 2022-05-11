package main

import (
	"fmt"
	"log"
	"net/http"
	"nilai-mahasiswa/controllers"
	"nilai-mahasiswa/middlewares"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// // this code for testing database in the beginning
	// db, e := config.MySQL()

	// if e != nil {
	// 	log.Fatal(e)
	// }

	// eb := db.Ping()
	// if eb != nil {
	// 	panic(eb.Error())
	// }

	// fmt.Println("success")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	router := httprouter.New()

	router.GET("/mahasiswa", controllers.GetMahasiswa)
	router.POST("/mahasiswa", middlewares.BasicAuth(controllers.InsertMahasiswa))
	router.PUT("/mahasiswa/:id", middlewares.BasicAuth(controllers.UpdateMahasiswa))
	router.DELETE("/mahasiswa/:id", middlewares.BasicAuth(controllers.DeleteMahasiswa))

	fmt.Println("Server is Running at Port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
