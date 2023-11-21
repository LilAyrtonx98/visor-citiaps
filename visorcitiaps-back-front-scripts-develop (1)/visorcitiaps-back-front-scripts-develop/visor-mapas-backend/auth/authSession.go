package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
)

//Estructura para obtener el usuario y contraseña del cuerpo de la solicitud rest
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	//key   = []byte(utils.Config.Auth.Session.SecretKey)
	//store = sessions.NewCookieStore(key)
	store *sessions.CookieStore
)

func InitAuthSession() {

	keyString := utils.Config.Auth.Session.SecretKey
	store = sessions.NewCookieStore([]byte(keyString))

	log.Println("Inicio de sesión mediante cookies activado")
}

func LoginSessionHandler(c *gin.Context) {

	//Se crea la cookie de sesión
	session, _ := store.Get(c.Request, "session_cookie")
	session.Options = &sessions.Options{
		/*
			Path: "/api",
			//Domain:   "localhost",
			MaxAge:   300, //En X segundos más a partir de la generación de la cookie
			HttpOnly: true,
		*/
		Path:     utils.Config.Auth.Session.Path,
		MaxAge:   utils.Config.Auth.Session.MaxAge, //En X segundos más a partir de la generación de la cookie
		HttpOnly: utils.Config.Auth.Session.HttpOnly,
	}

	//Obtener credenciales del usuario desde el JSON
	var credentials Credentials
	c.BindJSON(&credentials)

	username := credentials.Username
	password := credentials.Password

	//Verificar que las credenciales no estan vacías
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Parameters can't be empty"})
		return
	}

	//Verificar credenciales
	if ok, user := models.CheckPasswordByUsername(username, password); ok && user != nil {

		// Crear token con información del usuario
		userToken, status := CreateJWT(user)
		if status != 0 {
			c.AbortWithStatusJSON(status, gin.H{"message": "login error"})
			return
		}

		session.Values["user_token"] = userToken // Cookie con JWT del usuario

		//Permisos de usuario serializados
		serialized, _ := json.Marshal(user.Permissions)

		//Header con los permisos del usuario. Navegador lo guarda en una cookie local
		c.Writer.Header().Set("UserPermissions", string(serialized))

		err := session.Save(c.Request, c.Writer)
		if err != nil { //Si no se puede generar la cookie de sesión
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate session cookie"})
		} else { //Si las credenciales coinciden

			//Registrar acción en log de usuario
			models.CreateLog(user.ID.Hex(), "Inicio de sesión", "", "", "", "", false)

			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		}
	} else {
		//Si las credenciales no coinciden
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
	}

}

func LogoutSessionHandler(c *gin.Context) {

	session, _ := store.Get(c.Request, "session_cookie")
	auth, ok := session.Values["user_token"].(string)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid session cookie"})
		return
	} else {
		session.Options = &sessions.Options{
			Path: "/api",
			//Domain:   "localhost",
			MaxAge:   -1, //Si es negativo quiere decir que ya esta expirada y el navegador la borra
			HttpOnly: true,
		}

		// Revoke users authentication
		session.Values["user_token"] = ""
		session.Save(c.Request, c.Writer)

		//Registrar acción en log de usuario
		userClaims, status := ValidateJWT(auth)
		if status != 0 {
			return
		}
		models.CreateLog(userClaims.UserID, "Cierre de sesión", "", "", "", "", false)

		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	}
}

func AddCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func GetUserIDFromCookie(c *gin.Context) (string, bool) {
	//Leer cookie. Si no hay ninguna, retorna una cookie vacía
	session, _ := store.Get(c.Request, "session_cookie")

	//Verificar si el usuario esta autentificado en la cookie
	auth, ok := session.Values["user_token"].(string)
	if !ok || auth == "" {
		return "", false
	}
	userClaims, status := ValidateJWT(auth)
	if status != 0 {
		return "", false
	}
	return userClaims.UserID, true
}
