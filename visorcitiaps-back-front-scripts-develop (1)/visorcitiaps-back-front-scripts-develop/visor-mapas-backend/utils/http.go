package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
  }

func SendSingleFile(url, filePath, filename, datastore, workspace string) (err error) {

	//Abrir archivo
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return err
	}
	//Leer todo el archivo
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	//Preparar request
	var buffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&buffer)

	//Crea campo (file) 'file' con la ruta del archivo
	//Necesario para cargar un archivo y enviarlo mediante PUT
	formWriter, err := multipartWriter.CreateFormFile("file", filename)
	if err != nil {
		log.Println(err)
		return err
	}

	//Escribe los bytes de la imagen en el campo 'file'
	_, err = formWriter.Write(fileBytes)
	if err != nil {
		log.Println(err)
		return err
	}

	//Agregar campos (text) adicionales junto a su valor
	// err = multipartWriter.WriteField("datastore", datastore)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	
	// err = multipartWriter.WriteField("coordinateSystem", coordinateSystem)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	
	// err = multipartWriter.WriteField("workspace", workspace)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	multipartWriter.Close()

	//Crear request (sin ejecutar)
	req, err := http.NewRequest("PUT", url, &buffer) // debe ser put xd
	if err != nil {
		log.Println(err)
		return err
	}

	//Añadir la información en header request
	req.Header.Set("Content-Type", "application/zip")
	req.SetBasicAuth("admin", "geoserver")
	// req.Header.Add("Authorization","Basic " + basicAuth("admin","geoserver"))

	log.Println(fmt.Sprintf("Variables %s, %s, %s, %s", url, workspace, datastore, filename))
	//Ejecutar request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	//Verificar respuesta
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err = fmt.Errorf("bad status: %s", res.Status)
		log.Println(fmt.Errorf("error: %s", err))
		return err
	}

	log.Println("Archivo enviado exitosamente a " + url)

	return nil
}

func PaginationFindOptions(pageNumber, numberPerPage string) (*options.FindOptions, error) {

	page, err := strconv.ParseInt(pageNumber, 10, 8)
	size, err2 := strconv.ParseInt(numberPerPage, 10, 8)

	if err != nil {
		return nil, err
	}

	if err2 != nil {
		return nil, err2
	}

	var skipNumber int64

	if page > 0 {
		skipNumber = (page - 1) * size
	} else {
		skipNumber = 0
	}

	findOptions := options.Find()
	findOptions.SetLimit(size)
	findOptions.SetSkip(skipNumber)
	//Ordedar de acuerdo a '_id' descendente
	findOptions.Sort = bson.D{primitive.E{Key: "_id", Value: -1}}

	return findOptions, nil
}

func EstimatedTotalPages(numberPerPage string, count, skip int64) string {

	size, _ := strconv.ParseInt(numberPerPage, 10, 8)

	var result int64

	if ((skip + count) % size) > 0 {
		result = ((skip + count) / size) + 1
	} else {
		result = ((skip + count) / size)
	}
	return strconv.FormatInt(result, 10)
}
