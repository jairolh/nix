package utilidades

import (
    "encoding/json"
    "os"
)

type Configuration struct{
	//DB
	Host string
	DataBaseName string
	User string
	Password string
}

var Parametros Configuration

func Init()  Configuration {

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	Parametros = Configuration{}
	
	// se adiciona ya que no lee el archivo json
	Parametros.Host = "localhost" 
    Parametros.DataBaseName = "nix" 
    Parametros.User = "nixsgf" 
    Parametros.Password = "4dm1n=NIX2016"
  
	err := decoder.Decode(&Parametros)
	CheckErr(err,"Error cargando parametros de configuracion")
	return Parametros

}
