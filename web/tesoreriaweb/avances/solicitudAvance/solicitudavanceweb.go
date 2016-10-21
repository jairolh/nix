package solicitudavanceweb

import (
	
	_"fmt"
	"nix/utilidades"
	"nix/model/tesoreriaModel/avances/solicitudAvance"
	"nix/repository/tesoreriarepository/avances/solicitudAvance"
	"strconv"
	"strings"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_"encoding/json"
	
)

func Init(router *gin.Engine, middleware *jwt.GinJWTMiddleware) {

	apiSolicitudavance := router.Group("/tesoreria")
	//apiTipoavance.Use(middleware.MiddlewareFunc())

	apiSolicitudavance.GET("/solicitudavance/:opcion/:vigencia", List)
	apiSolicitudavance.GET("/solicitudavance/:opcion/:vigencia/:idSolicitud/:idTipo", FindOne)
	apiSolicitudavance.POST("/solicitudavance", Create)
	//apiTipoavance.PUT("/tipoavance/:idtipo", Modify)
	//apiTipoavance.PUT("/tipoavance", Modify)
	//apiTipoavance.DELETE("/tipoavance/:idtipo", Delete)
	apiSolicitudavance.OPTIONS("/solicitudavance", Options) 
	solicitudavancerepository.Init()

}

func List(c *gin.Context) {

    opcion := strings.TrimSpace(c.Params.ByName("opcion"))
	vigencia, _ := strconv.ParseInt(c.Params.ByName("vigencia"), 0, 64)

	switch opcion {
		case "secuencia":
		    secuencia, msg := solicitudavancerepository.FindOneSecuencia(vigencia)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, secuencia) }
		default:
		    solicitudesavance, msg := solicitudavancerepository.FindAll(vigencia)
			if msg.Code != 0 {c.JSON(200, msg)} else {c.JSON(200, solicitudesavance)
		}

	}

}

func FindOne(c *gin.Context) {

    opcion := strings.TrimSpace(c.Params.ByName("opcion"))
    vigencia, _ := strconv.ParseInt(c.Params.ByName("vigencia"), 0, 64)
	solicitud, _ := strconv.ParseInt(c.Params.ByName("idSolicitud"), 0, 64)
	tipo, _ := strconv.ParseInt(c.Params.ByName("idTipo"), 0, 64)

	switch opcion {
		case "solicitud":
		    avance, msg := solicitudavancerepository.FindOne(vigencia,solicitud)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, avance) }
		case "tiposAvance":
		    tipos, msg := solicitudavancerepository.FindAllTipo(solicitud)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, tipos) }
		case "requisitosTiposAvance":
		    requisitos, msg := solicitudavancerepository.FindAllReq(tipo)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, requisitos) }
		}
}

func Create(c *gin.Context) {

	var avanceins solicitudavance.Solicitud
//	fmt.Println("dat  :",avanceins)
	 if c.Bind(&avanceins) == nil {
	 		//validacion y registro de beneficiario
	 		beneficiarioavanceins := avanceins.Beneficiario
			beneficiarioAvance, msgBene := solicitudavancerepository.FindOneBeneficiario(beneficiarioavanceins)
			var IdBene int64
			if msgBene.Code != 0 {
				//registra beneficiarios
				msgIns := solicitudavancerepository.CreateBeneficiarioAvance(beneficiarioavanceins)
				c.JSON(200, msgIns)
				IdBene=beneficiarioavanceins.IdBeneficiario
				} else {
				c.JSON(200, utilidades.CheckInfo( "El Beneficiario ya esta registrado"))	
				IdBene=beneficiarioAvance.IdBeneficiario
				}
	 		//validacion y registro de solicitud
			solicitudavanceins := avanceins.Solicitud
			solicitudAvance, msgSol := solicitudavancerepository.FindOneSolicitudSec(solicitudavanceins)
			var IdSol int64
			IdSol=0
			if msgSol.Code != 0 {	//registra Solicitud
				solicitudavanceins.IdBeneficiario=IdBene
				msgIns := solicitudavancerepository.CreateSolicitud(solicitudavanceins)
				c.JSON(200, msgIns)
				solicitud, msgSolin := solicitudavancerepository.FindOneSolicitudSec(solicitudavanceins)
				IdSol=solicitud.IdSolicitud
				if msgSolin.Code != 0 {
					IdSol=solicitud.IdSolicitud
					} else { c.JSON(200, msgSolin)
					}
				} else {
				c.JSON(200, utilidades.CheckInfo( "El consecutivo de la solicitud de avance ya existe para la vigencia "))
				IdSol=solicitudAvance.IdSolicitud
				}
			//fmt.Println("IDS :",IdSol)
			//valida que el numero de solicitud exista
			if IdSol > 0 {
		 		//validacion y registro de tipo de avance
				solicitudtipoavanceins := avanceins.Tipoavance
				solicitudtipoavanceins.IdSolicitud = IdSol
				_, msgTipo := solicitudavancerepository.FindOneSolicitudTipoAvance(solicitudtipoavanceins)
				if msgTipo.Code != 0 {
					//registra Solicitud
					msgIns := solicitudavancerepository.CreateSolicitudTipo(solicitudtipoavanceins)
					c.JSON(200, msgIns)
					} else {
						c.JSON(200, utilidades.CheckInfo( "El tipo de avance ya existe"))
					}
				var IdEst int64	
				var consultaEstado solicitudavance.Estados	
				consultaEstado.Nombre = "registrado"
				consultaEstado.Proceso = "avances"
				resEstado, msgEst := solicitudavancerepository.FindOneEstado(consultaEstado)

				if msgEst.Code != 0 {
					IdEst=0
				} else {
					//c.JSON(200, resEstado)
					IdEst=resEstado.IdEstado
				}
				//validacion y registro de estados de la solicitud de avance
				estadosolicitudavanceins := avanceins.Estadosolicitud
				estadosolicitudavanceins.IdSolicitud = IdSol
				estadosolicitudavanceins.IdEstado=IdEst
				estadosolicitudavanceins.Observaciones="Registro inicial de la Solicitud de Avance"
				//fmt.Println("IDW :",estadosolicitudavanceins)
				_, msgEstAv := solicitudavancerepository.FindOneEstadoAvance(estadosolicitudavanceins)
				if msgEstAv.Code != 0 {
					//registra estados
						msgIns := solicitudavancerepository.CreateEstadoSolicitud(estadosolicitudavanceins)
						c.JSON(200, msgIns)
					} else {
						c.JSON(200, utilidades.CheckInfo( "El Estado de la solicitud ya existe"))
					}	
				}
				c.JSON(200, "Se registro la solicitud")
	 		}else {
	 		c.JSON(400, "NO se pudo rescatar datos") 
	 		}
}

/*
func Modify(c *gin.Context) {

	var tipoavanceupd tipoavance.Tipoavance
	c.Bind(&tipoavanceupd)
	msg := tipoavancerepository.Update(tipoavanceupd)

	c.JSON(200, msg)
}

func Delete(c *gin.Context) {
	tipoavanceid, _ := strconv.ParseInt(c.Params.ByName("idtipo"), 0, 64)
	msg := tipoavancerepository.Delete(tipoavanceid)
	c.JSON(200, msg)
}*/

/*Options funcion para peticiones de otros servidores */
func Options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

