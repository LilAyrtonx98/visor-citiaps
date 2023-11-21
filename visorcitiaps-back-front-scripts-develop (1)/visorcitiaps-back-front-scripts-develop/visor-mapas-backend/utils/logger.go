package utils

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func InitLogger() {

	dateTime := time.Now()
	dateTimeFormat := dateTime.Format("2006-01-02_15-04-05")
	logFileName := dateTimeFormat + "_" + Config.Server.Logfile
	logPath := Config.Server.Logpath

	// Verificar si existe el directorio y crearlo si no existe
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err = os.MkdirAll(logPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	fullpath := logPath + logFileName

	log.Println("De ahora en adelante logs se encuentran en", fullpath)
	logs, err := os.OpenFile(logPath+logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logs)
	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	log.Println("Aplicación iniciada")
}
