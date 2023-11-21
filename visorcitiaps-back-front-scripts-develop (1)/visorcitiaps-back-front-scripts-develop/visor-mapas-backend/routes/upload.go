package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/citiaps/visor-mapas-backend/utils"
)

const uploadPath = "/tmp/visor-tmp/"

//Cliente sube un archivo al backend mediante método POST y backend lo recibe
func ReceiveSingleFile() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Recibir valores del body request
		datastore := c.PostForm("datastore")
		//coordinateSystem := c.PostForm("coordinateSystem")
		workspace := c.PostForm("workspace")
		filename := c.PostForm("filename")

		//Verificar existencia del directorio
		_, err := os.Stat(uploadPath)
		if os.IsNotExist(err) {
			os.MkdirAll(uploadPath, os.ModePerm)
		}

		//Leer archivo
		file, err := c.FormFile("file")
		if err != nil {
			//c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ANo se puede subir el archivo"})
			return
		}

		//Crear ruta + nombre archivo
		//filePath := uploadPath + filepath.Base(file.Filename)
		filePath := uploadPath + filepath.Base(filename)

		//Guardar archivo
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			//c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "BNo se puede subir el archivo"})
			return
		} else {
			//log.Println(fmt.Sprintf("Archivo %s subido exitosamente en %s", file.Filename, uploadPath))
			log.Println(fmt.Sprintf("Archivo %s subido exitosamente en %s", filename, uploadPath))

			//Subir aquí a GeoServer
			geoserverUrl := utils.Config.GeoServer.Host + utils.Config.GeoServer.Postpath
			log.Println(fmt.Sprintf("URL de subida %s", geoserverUrl))
			err := utils.SendSingleFile(geoserverUrl, filePath, filename, datastore, workspace)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "CNo se puede enviar el archivo a GeoServer"})
				return
			}

			log.Println(fmt.Sprintf("Archivo en geoserver"))
			c.JSON(http.StatusCreated, gin.H{
				"message": "Archivo subido exitosamente",
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Archivo subido exitosamente",
		})

	}

}

//Cliente sube un archivo al backend mediante método POST y backend lo recibe
func ReceiveSingleFileDummy() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Verificar existencia del directorio
		_, err := os.Stat("/tmp/visor-tmp2/")
		if os.IsNotExist(err) {
			os.MkdirAll("/tmp/visor-tmp2/", os.ModePerm)
		}

		//Leer archivo
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "No se puede subir el archivo"})
			return
		}

		//Crear ruta + nombre archivo
		filePath := "/tmp/visor-tmp2/" + filepath.Base(file.Filename)

		//Guardar archivo
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			//c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "No se puede subir el archivo"})
			return
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Archivo subido exitosamente",
			})
			return

		}

	}

}
