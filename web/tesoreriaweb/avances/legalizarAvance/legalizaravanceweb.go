package legalizaravanceweb

import (
	
	"fmt"
	_"nix/utilidades"
	_"nix/model/tesoreriaModel/avances/legalizarAvance"
	"nix/repository/tesoreriarepository/avances/legalizarAvance"
	"strconv"
	"strings"
	_"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_"encoding/json"
	
)

func Init(router *gin.Engine, middleware *jwt.GinJWTMiddleware) {

	apiLegalizaravance := router.Group("/tesoreria")
	//apiTipoavance.Use(middleware.MiddlewareFunc())

	apiLegalizaravance.GET("/legalizaravance/:opcion/:vigencia", List)
	apiLegalizaravance.GET("/legalizaravance/:opcion/:vigencia/:idSolicitud/:idTipo", FindOne)
	apiLegalizaravance.POST("/legalizaravance/:opcion", Create)
	//apiTipoavance.PUT("/tipoavance/:idtipo", Modify)
	//apiTipoavance.PUT("/tipoavance", Modify)
	//apiTipoavance.DELETE("/tipoavance/:idtipo", Delete)
	apiLegalizaravance.OPTIONS("/legalizaravance/:opcion", Options) 
	legalizaravancerepository.Init()
	fmt.Println("")
}

func List(c *gin.Context) {

    opcion := strings.TrimSpace(c.Params.ByName("opcion"))
	vigencia, _ := strconv.ParseInt(c.Params.ByName("vigencia"), 0, 64)

	switch opcion {
		case "legalizado":
		    		fmt.Println("")
		default:
		    solicitudesavance, msg := legalizaravancerepository.FindAll(vigencia)
			if msg.Code != 0 {c.JSON(200, msg)} else {c.JSON(200, solicitudesavance)
		}

	}

}

func FindOne(c *gin.Context) {

    opcion := strings.TrimSpace(c.Params.ByName("opcion"))
    //vigencia, _ := strconv.ParseInt(c.Params.ByName("vigencia"), 0, 64)
	//solicitud, _ := strconv.ParseInt(c.Params.ByName("idSolicitud"), 0, 64)
	//tipo, _ := strconv.ParseInt(c.Params.ByName("idTipo"), 0, 64)

	switch opcion {
		case "solicitud":
				fmt.Println("")  
		}
}

func Create(c *gin.Context) {
opcion := strings.TrimSpace(c.Params.ByName("opcion"))
switch opcion {
			case "solicitud":
				  fmt.Println("")
    
			}//fin switch

}


/*Options funcion para peticiones de otros servidores */
func Options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
