package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/citiaps/visor-mapas-backend/models"
)

func LoginJWTHandler(c *gin.Context) {
	var credentials Credentials
	c.BindJSON(&credentials)

	username := credentials.Username
	password := credentials.Password

	//Si coinciden, generar token
	if ok, user := models.CheckPasswordByUsername(username, password); ok && user != nil {

		// Crear token con información del usuario
		tokenString, status := CreateJWT(user)
		if status != 0 {
			c.AbortWithStatusJSON(status, gin.H{"message": "login error"})
			return
		}

		//Permisos de usuario serializados
		serialized, _ := json.Marshal(user.Permissions)

		//Header con los permisos del usuario. Navegador lo guarda en una cookie local
		c.Writer.Header().Set("UserPermissions", string(serialized))

		//Si todo está bien, entregar token en el response header
		c.Writer.Header().Set("Token", tokenString)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	} else { // NO coinciden, detener
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "wrong credentials"})
		return
	}
}

func LogoutJWTHandler(c *gin.Context) {

}
