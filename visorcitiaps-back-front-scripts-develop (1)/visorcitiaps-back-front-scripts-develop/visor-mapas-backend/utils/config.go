package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Config Configuration

type Database struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type Redis struct {
	Pass string `yaml:"pass"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   string `yaml:"db"`
}

type Server struct {
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	Logfile string `yaml:"logfile"`
	Logpath string `yaml:"logpath"`
}

type Session struct {
	Path      string `yaml:"path"`
	MaxAge    int    `yaml:"maxage"`
	HttpOnly  bool   `yaml:"httponly"`
	SecretKey string `yaml:"secretkey"`
}

type JWT struct {
	Exp       int    `yaml:"exp"`
	SecretKey string `yaml:"secretkey"`
}

type Auth struct {
	Session Session `yaml:"session"`
	JWT     JWT     `yaml:"jwt"`
}

type APIRest struct {
	Page string `yaml:"page"`
	Size string `yaml:"size"`
}

type Email struct {
	Smtp     string `yaml:"smtp"`
	Port     string `yaml:"port"`
	Sender   string `yaml:"sender"`
	Password string `yaml:"password"`
}

type GeoServer struct {
	Host       string `yaml:"host"`
	Layerspath string `yaml:"layerspath"`
	Postpath   string `yaml:"postpath"`
}

type CORS struct {
	Origin string `yaml:"origin"`
}

type Configuration struct {
	Database  Database  `yaml:"database"`
	Redis     Redis     `yaml:"redis"`
	Server    Server    `yaml:"server"`
	Auth      Auth      `yaml:"auth"`
	APIRest   APIRest   `yaml:"apirest"`
	Email     Email     `yaml:"mail"`
	GeoServer GeoServer `yaml:"geoserver"`
	CORS      CORS      `yaml:"cors"`
}

func LoadConfig(filename string) {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error al cargar configuraciones: yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Error al cargar configuraciones: Unmarshal: %v", err)
	}
	log.Println("Configuraciones cargadas exitosamente")
}
