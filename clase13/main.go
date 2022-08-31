package main

// Vamos a crear una aplicación web con el framework Gin que tenga un endpoint /productos que responda con una lista de productos.
// Crear una estructura Productos con los valores:
// Id
// Nombre
// Precio
// Stock
// Codigo
// Publicado
// FechaDeCreacion
// Crear un archivo productos.json con al menos seis productos (deben seguir la misma estructura ya mencionada).
// Leer el archivo productos.json.
//
// Imprimir la lista de productos a través del endpoint en formato JSON. El endpoint deberá ser de método GET.

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

func main() {
	var prod []Product
	data, _ := os.ReadFile("./productos.json")
	if err := json.Unmarshal(data, &prod); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/productos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, prod)
	})
	router.Run()
}
