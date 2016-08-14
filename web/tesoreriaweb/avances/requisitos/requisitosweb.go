package requisitosweb

import (
	"fmt"
	"nix/model/tesoreriaModel/avances/requisitos"
	"nix/repository/tesoreriarepository/avances/requisitos"
	"strconv"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, middleware *jwt.GinJWTMiddleware) {

	apiRequisito := router.Group("/tesoreria")
	//apiRequisito.Use(middleware.MiddlewareFunc())

	apiRequisito.GET("/requisito", List)
	apiRequisito.GET("/requisito/:idreq", FindOne)
	apiRequisito.POST("/requisito", Create)
	//apiRequisito.PUT("/Requisito/:idtipo", Modify)
	apiRequisito.PUT("/requisito", Modify)
	apiRequisito.DELETE("/requisito/:idreq", Delete)
	apiRequisito.OPTIONS("/requisito", Options) 
	requisitorepository.Init()

}

func List(c *gin.Context) {
	requisitos, msg := requisitorepository.FindAll()
	if msg.Code != 0 {
		c.JSON(200, msg)
	} else {
		c.JSON(200, requisitos)
	}

}

func FindOne(c *gin.Context) {
	requisitoid, _ := strconv.ParseInt(c.Params.ByName("idreq"), 0, 64)
	fmt.Println("IDW :",requisitoid)
	requisito, msg := requisitorepository.FindOne(requisitoid)
	if msg.Code != 0 {
		c.JSON(200, msg)
	} else {
		c.JSON(200, requisito)
	}
}

func Create(c *gin.Context) {
	var requisitoins requisito.Requisito
	c.Bind(&requisitoins)
	msg := requisitorepository.Create(requisitoins)
	c.JSON(200, msg)

}

func Modify(c *gin.Context) {
	var requisitoupd requisito.Requisito
	c.Bind(&requisitoupd)
	msg := requisitorepository.Update(requisitoupd)
	c.JSON(200, msg)
}

func Delete(c *gin.Context) {
	requisitoid, _ := strconv.ParseInt(c.Params.ByName("idreq"), 0, 64)
	msg := requisitorepository.Delete(requisitoid)
	c.JSON(200, msg)
}

/*Options funcion para peticiones de otros servidores */
func Options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

