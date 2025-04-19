package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {

	// Define your routes here
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.POST("/usuarios", func(c *gin.Context) {
		var nuevoUsuario Usuario
		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el JSON"})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Apellido == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre y apellido son obligatorios"})
			return
		}

		usuarios = append(usuarios, nuevoUsuario)
		c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente", "usuario": nuevoUsuario})

	})

	r.GET("/usuarios", func(c *gin.Context) {
		if len(usuarios) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron usuarios"})
			return
		}

		c.JSON(http.StatusOK, usuarios)
	})

}
