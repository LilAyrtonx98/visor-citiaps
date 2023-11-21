package auth

import (
	"net/http"
	"strings"

	"github.com/citiaps/visor-mapas-backend/models"
	"github.com/citiaps/visor-mapas-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type User = models.User

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := utils.Config.CORS.Origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", domain) // Aqui debería ir el dominio del frontend para que se pueda acceder desde cualquier lugar y no tener problemas de cors
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With, access-control-allow-origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Totalpages, UserPermissions, Token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthSessionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Leer cookie. Si no hay ninguna, retorna una cookie vacía
		session, _ := store.Get(c.Request, "session_cookie")

		//Verificar si el usuario esta autentificado en la cookie y tiene token
		if auth, ok := session.Values["user_token"].(string); !ok || auth == "" {
			//Si la cookie no es válida, no deja continuar
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "This is secret and you can't see it because your cookie is not valid"})
			return
		} else if userClaims, status := ValidateJWT(auth); status != 0 {
			// Si el token del usuario dentro de la cookie no es válido, no se deja continuar
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "This is secret and you can't see it because the token in your cookie is not valid"})
			return
		} else if ok := CheckSessionPermission(permission, userClaims); !ok {
			//Si el usuario no tiene permisos necesarios, no deja continuar
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "This is secret and you can't see it because you don't have permission"})
			return
		} else {
			//Refrescar cookie
			session.Options = &sessions.Options{
				Path:     utils.Config.Auth.Session.Path,
				MaxAge:   utils.Config.Auth.Session.MaxAge, //En X segundos más a partir de la generación de la cookie
				HttpOnly: utils.Config.Auth.Session.HttpOnly,
			}

			// Refrescar token
			userToken := AutoRefreshJWT(auth)
			session.Values["user_token"] = userToken

			_ = session.Save(c.Request, c.Writer)

			//Continuar
			c.Next()
		}
	}
}

func CheckSessionPermission(permission string, claims *Claims) bool {
	switch permission {
	case "users":
		return claims.PermissionUsers
	case "layers":
		return claims.PermissionLayers
	case "geo":
		return claims.PermissionGeo
	case "maps":
		return claims.PermissionMaps
	case "visor":
		return claims.PermissionVisor
	case "annotation":
		return claims.PermissionAnnotation
	case "admin": // Solo necesita tener algún permiso de administración
		return (claims.PermissionUsers || claims.PermissionLayers || claims.PermissionGeo || claims.PermissionMaps)
	case "none": // No necesita un permiso específico
		return true
	default:
		return false
	}
}

//Solo verifica si hay token. Si no tiene token, no se termina la solicitud
func AuthJWTMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		splitToken := strings.Split(tokenString, "Bearer ")

		//Si hay algún valor luego de 'Bearer ', se verifica el token
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
			claims, status := ValidateJWT(tokenString)
			if status != 0 {
				//Token no válido
				c.AbortWithStatusJSON(status, gin.H{"message": "This is secret and you can't see it because your token is not valid"})
				return
			} else if ok := CheckSessionPermission(permission, claims); !ok {
				//Si el usuario no tiene permisos necesarios, no deja continuar
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "This is secret and you can't see it because you don't have permission"})
				return
			}
			//Token valido. Continuar
			c.Next()
		} else {
			//No hay token
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "authorization error"})
			return
		}
	}
}
