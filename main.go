package main

import (
	"fmt"
	"log"
	"net/http"
	"nilai-mahasiswa/controllers"
	"nilai-mahasiswa/middlewares"

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

	router := httprouter.New()

	router.GET("/mahasiswa", controllers.GetMahasiswa)
	router.POST("/mahasiswa", middlewares.BasicAuth(controllers.InsertMahasiswa))
	router.PUT("/mahasiswa/:id", middlewares.BasicAuth(controllers.UpdateMahasiswa))
	router.DELETE("/mahasiswa/:id", middlewares.BasicAuth(controllers.DeleteMahasiswa))

	fmt.Println("Server is Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
