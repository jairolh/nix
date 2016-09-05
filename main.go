package main

import (
	"nix/utilidades"

	"nix/web/security"
	"nix/web/tesoreriaweb/avances/tipoAvance"
	"nix/web/tesoreriaweb/avances/requisitos"
	"nix/web/tesoreriaweb/avances/requisitoTipoAvance"
	"nix/web/tesoreriaweb/avances/solicitudAvance"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {

	//carga inicial de parametros de configuracion
	utilidades.Init()

	//Inicializacion de framework Gin-gonic
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())

	//inicializacion de middleware de AUTH
	security.Init(r)

	//inicializacion de todos los routes del aplicativo
	tipoavanceweb.Init(r, security.MiddlewareAUTH)
	requisitosweb.Init(r, security.MiddlewareAUTH)
	requisitotipoavanceweb.Init(r, security.MiddlewareAUTH)
	solicitudavanceweb.Init(r, security.MiddlewareAUTH)



	//inicia el servidor
	endless.ListenAndServe(":8088", r)
}
/*
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}*/


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}

/*Options funcion para peticiones de otros servidores */
func Options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
