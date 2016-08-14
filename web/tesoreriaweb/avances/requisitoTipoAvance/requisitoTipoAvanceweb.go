package requisitotipoavanceweb

import (
	"fmt"
	"nix/model/tesoreriaModel/avances/requisitoTipoAvance"
	"nix/repository/tesoreriarepository/avances/requisitoTipoAvance"
	"strconv"
	"strings"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, middleware *jwt.GinJWTMiddleware) {

	apiRequisitoAvance := router.Group("/tesoreria")
	//apiRequisito.Use(middleware.MiddlewareFunc())

	apiRequisitoAvance.GET("/requisitoAvance/:opcion/:idtipo", List)
	apiRequisitoAvance.GET("/requisitoAvance/:opcion/:idtipo/:idreq", FindOne)
	apiRequisitoAvance.POST("/requisitoAvance", Create)
	//apiRequisitoAvance.PUT("/Requisito/:idtipo", Modify)
	apiRequisitoAvance.PUT("/requisitoAvance", Modify)
	//apiRequisitoAvance.DELETE("/requisito/:idreq", Delete)
	apiRequisitoAvance.OPTIONS("/requisitoAvance", Options) 
	requisitotipoavancerepository.Init()

}

func List(c *gin.Context) {

    opcion := strings.TrimSpace(c.Params.ByName("opcion"))
	idtipo, _ := strconv.ParseInt(c.Params.ByName("idtipo"), 0, 64)

	switch opcion {
	case "selreq":
	    requisitos, msg := requisitotipoavancerepository.FindSelected(idtipo)
	    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, requisitos) }
	default:
	    requisitosAvance, msg := requisitotipoavancerepository.FindAll(idtipo)
	    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, requisitosAvance) }
	}

	
	

}

func FindOne(c *gin.Context) {

	tipoavanceid, _ := strconv.ParseInt(c.Params.ByName("idtipo"), 0, 64)
	requisitoid, _ := strconv.ParseInt(c.Params.ByName("idreq"), 0, 64)
	fmt.Println("IDW :",tipoavanceid,requisitoid)
	requisitoAvance, msg := requisitotipoavancerepository.FindOne(tipoavanceid,requisitoid)
	if msg.Code != 0 {
		c.JSON(200, msg)
	} else {
		c.JSON(200, requisitoAvance)
	}
}

func Create(c *gin.Context) {
	var requisitoavanceins requisitotipoavance.RequisitoTipoavance
	c.Bind(&requisitoavanceins)
	msg := requisitotipoavancerepository.Create(requisitoavanceins)
	c.JSON(200, msg)

}
func Modify(c *gin.Context) {
	var requisitoavanceupd requisitotipoavance.RequisitoTipoavance
	c.Bind(&requisitoavanceupd)
	msg := requisitotipoavancerepository.Update(requisitoavanceupd)
	c.JSON(200, msg)

}
/*
func Delete(c *gin.Context) {
	requisitoid, _ := strconv.ParseInt(c.Params.ByName("idreq"), 0, 64)
	msg := requisitorepository.Delete(requisitoid)
	c.JSON(200, msg)
}
*/
/*Options funcion para peticiones de otros servidores */
func Options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

