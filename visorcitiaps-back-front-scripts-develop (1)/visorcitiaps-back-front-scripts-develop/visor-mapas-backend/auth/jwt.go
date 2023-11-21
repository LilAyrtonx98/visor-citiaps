package auth

import (
	"log"
	"net/http"
	"time"

	"gitlab.com/cesar.kreep/visor-mapas-backend/utils"
	"github.com/dgrijalva/jwt-go"
)

//Llave secreta del JWT
var jwtKey []byte

//Estructura del JWT. Campos personalizados + estándar
type Claims struct {
	UserID               string `json:"user_id"`
	PermissionUsers      bool   `json:"permission_users"`
	PermissionLayers     bool   `json:"permission_layers"`
	PermissionGeo        bool   `json:"permission_geo"`
	PermissionMaps       bool   `json:"permission_maps"`
	PermissionVisor      bool   `json:"permission_visor"`
	PermissionAnnotation bool   `json:"permission_annotation"`
	jwt.StandardClaims
}

//Ejecutar en el main
func InitJWTSession() {
	jwtKey = []byte(utils.Config.Auth.JWT.SecretKey)

	log.Println("Inicio de sesión mediante JWT activado")
}

func CreateJWT(user *User) (string, int) {

	// Declare the expiration time of the token
	//Convierte un int en time
	t := time.Duration(utils.Config.Auth.JWT.Exp)
	expirationTime := time.Now().Add(time.Minute * t)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserID:               user.ID.Hex(),
		PermissionUsers:      user.Permissions.Users,
		PermissionLayers:     user.Permissions.Layers,
		PermissionGeo:        user.Permissions.Geo,
		PermissionMaps:       user.Permissions.Maps,
		PermissionVisor:      user.Permissions.Visor,
		PermissionAnnotation: user.Permissions.Annotation,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	//Si hay error, se cancela
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", http.StatusInternalServerError
	}

	return tokenString, 0
}

func ValidateJWT(tokenString string) (*Claims, int) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		return claims, http.StatusUnauthorized
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, http.StatusUnauthorized
		}
		return claims, http.StatusBadRequest
	}

	return claims, 0 //0 indica que el token está bien
}

func RefreshJWT(tokenString string) (string, int) {
	claims, status := ValidateJWT(tokenString)
	if status != 0 {
		return "", status
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return "", http.StatusBadRequest
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", http.StatusInternalServerError
	}

	return newTokenString, 0
}

func AutoRefreshJWT(tokenString string) string {
	claims, status := ValidateJWT(tokenString)
	if status != 0 {
		return ""
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry.
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return tokenString
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(time.Duration(utils.Config.Auth.JWT.Exp) * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return tokenString
	}

	return newTokenString
}
