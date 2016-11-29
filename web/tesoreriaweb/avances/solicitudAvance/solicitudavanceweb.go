package solicitudavanceweb

import (
	
	"fmt"
	"nix/utilidades"
	"nix/model/tesoreriaModel/avances/solicitudAvance"
	"nix/repository/tesoreriarepository/avances/solicitudAvance"
	"strconv"
	"strings"
	"time"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_"encoding/json"
	
)

func Init(router *gin.Engine, middleware *jwt.GinJWTMiddleware) {

	apiSolicitudavance := router.Group("/tesoreria")
	//apiTipoavance.Use(middleware.MiddlewareFunc())

	apiSolicitudavance.GET("/solicitudavance/:opcion/:vigencia", List)
	apiSolicitudavance.GET("/solicitudavance/:opcion/:vigencia/:idSolicitud/:idTipo", FindOne)
	apiSolicitudavance.POST("/solicitudavance/:opcion", Create)
	//apiTipoavance.PUT("/tipoavance/:idtipo", Modify)
	//apiTipoavance.PUT("/tipoavance", Modify)
	//apiTipoavance.DELETE("/tipoavance/:idtipo", Delete)
	apiSolicitudavance.OPTIONS("/solicitudavance/:opcion", Options) 
	solicitudavancerepository.Init()
	fmt.Println("")
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
		case "requisitosSolicitudAvance":
		    requisitosAvn, msg := solicitudavancerepository.FindAllReqAvn(solicitud,tipo)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, requisitosAvn) }
		case "solicitudAvanceBeneficiario":
		    beneficiarioAvn, msg := solicitudavancerepository.FindEstadoAvanceBeneficiario(solicitud,tipo)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, beneficiarioAvn) }  
		case "financiaAvance":
		    necesidadAvn, msg := solicitudavancerepository.FindOneFinanciaAvance(vigencia,solicitud)
		    if msg.Code != 0 { c.JSON(200, msg) }  else {c.JSON(200, necesidadAvn) }         
		}
}

func Create(c *gin.Context) {
opcion := strings.TrimSpace(c.Params.ByName("opcion"))
switch opcion {
			case "solicitud":
				  CrearSolicitud(c);

			case "tipoavance":
					CrearTipoAvance(c);	
				
			case "verificaavance":
					CrearVerificarAvance(c);
				
			case "cancelavance":
				    CrearCancelarAvance(c);

			case "necesidadavance":
				    CrearNecesidadAvance(c);

			case "apruebaavance":
				    CrearApruebaAvance(c);
				    
			}//fin switch

}
/*
func Modify(c *gin.Context) {

	var tipoavanceupd tipoavance.Tipoavance
	c.Bind(&tipoavanceupd)
	msg := tipoavancerepository.Update(tipoavanceupd)

	c.JSON(200, msg)

    opcion := strings.TrimSpace(c.Params.ByName("opcion"))
    vigencia, _ := strconv.ParseInt(c.Params.ByName("vigencia"), 0, 64)
	solicitud, _ := strconv.ParseInt(c.Params.ByName("idSolicitud"), 0, 64)
	tipo, _ := strconv.ParseInt(c.Params.ByName("idTipo"), 0, 64)

	switch opcion {
			case "tipoavance":
				 var avanceupd solicitudavance.Solicitud
				 if c.Bind(&avanceupd) == nil {
					 solicitudtipoavanceins := avanceins.Tipoavance
					 solicitudtipoavanceins.IdSolicitud = avanceins.Solicitud.IdSolicitud
					 //fmt.Println("IDS :",solicitudtipoavanceins.IdSolicitud)
					 _, msgTipo := solicitudavancerepository.FindOneSolicitudTipoAvance(solicitudtipoavanceins)
					 if msgTipo.Code != 0 {
							//registra Solicitud
							msgIns := solicitudavancerepository.CreateSolicitudTipo(solicitudtipoavanceins)
							c.JSON(200, msgIns)
							} else {
								c.JSON(200, utilidades.CheckInfo( "El tipo de avance ya existe"))
							}
				}//fin if- tipoavance	
			}//fin switch



}
/*
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


/*Options funcion para registrar nueva solicitud de avance */
func CrearSolicitud(c *gin.Context) {
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
 		}//fin if - opcion solicitud	

}

/*Options funcion para registrar nuevo tipo de avance */
func CrearTipoAvance(c *gin.Context) {
	var avanceins solicitudavance.Solicitud
	 if c.Bind(&avanceins) == nil {
		 solicitudtipoavanceins := avanceins.Tipoavance
		 solicitudtipoavanceins.IdSolicitud = avanceins.Solicitud.IdSolicitud
		 //fmt.Println("IDS :",solicitudtipoavanceins.IdSolicitud)
		 _, msgTipo := solicitudavancerepository.FindOneSolicitudTipoAvance(solicitudtipoavanceins)
		 if msgTipo.Code != 0 {
				//registra Solicitud
				msgIns := solicitudavancerepository.CreateSolicitudTipo(solicitudtipoavanceins)
				c.JSON(200, msgIns)
				} else {
					c.JSON(200, utilidades.CheckInfo( "El tipo de avance ya existe"))
				}
	}//fin if- tipoavance	
}

/*Options funcion para registrar la verificacion de avance */
func CrearVerificarAvance(c *gin.Context) {
	var verificains []solicitudavance.RequisitoSolicitudavance
	 if c.Bind(&verificains) == nil {
	 	//fmt.Println("IDS :",verificains)
	 	fecha_registro:= time.Now().Format("2006-01-02 15:04:05")
	 	for i := 0; i < len(verificains); i += 1 {
		    registro := verificains [i]
		    registro.FechaRegistroReq =fecha_registro;
		    msgIns := solicitudavancerepository.CreateVerificaSolicitud(registro)
			c.JSON(200, msgIns)
			}

		var IdEst int64	
		var consultaEstado solicitudavance.Estados	
		consultaEstado.Nombre = "verificado"
		consultaEstado.Proceso = "avances"
		resEstado, msgEst := solicitudavancerepository.FindOneEstado(consultaEstado)
		if msgEst.Code != 0 {
			IdEst=0
		} else {
			//c.JSON(200, resEstado)
			IdEst=resEstado.IdEstado
		}
		//validacion y registro de estados de la solicitud de avance
		var estadosolicitudavanceins solicitudavance.Estadosolicitudavance
		estadosolicitudavanceins.IdSolicitud = verificains[0].IdSolicitud
		estadosolicitudavanceins.IdEstado=IdEst
		estadosolicitudavanceins.Observaciones="Verificados los requisitos de la solicitud de avance"
		estadosolicitudavanceins.Usuario = verificains[0].Usuario
		_, msgEstAv := solicitudavancerepository.FindOneEstadoAvance(estadosolicitudavanceins)
		if msgEstAv.Code != 0 {
			//registra estados
				msgIns := solicitudavancerepository.CreateEstadoSolicitud(estadosolicitudavanceins)
				c.JSON(200, msgIns)
			} else {
				c.JSON(200, utilidades.CheckInfo( "El Estado de la solicitud ya existe"))
			}	
		c.JSON(200, "Se registro la verficacion de la solicitud")	
	}//fin if- verificaavance
}

/*Options funcion para registrar la cancelacion de una solicitud de  avance */
func CrearCancelarAvance(c *gin.Context) {
	var avanceins solicitudavance.Estadosolicitudavance
	 if c.Bind(&avanceins) == nil {
	 	//fmt.Println("IDS :",avanceins)
	 	var IdEst int64	
		var consultaEstado solicitudavance.Estados	
		consultaEstado.Nombre = "cancelado"
		consultaEstado.Proceso = "avances"
		resEstado, msgEst := solicitudavancerepository.FindOneEstado(consultaEstado)
		if msgEst.Code != 0 {
			IdEst=0
		} else {
			//c.JSON(200, resEstado)
			IdEst=resEstado.IdEstado
		}
		//validacion y registro de estados de la solicitud de avance
		var estadosolicitudavanceins solicitudavance.Estadosolicitudavance
		estadosolicitudavanceins.IdSolicitud = avanceins.IdSolicitud
		estadosolicitudavanceins.IdEstado=IdEst
		estadosolicitudavanceins.Observaciones=avanceins.Observaciones
		estadosolicitudavanceins.Usuario = avanceins.Usuario
		//fmt.Println("CAN :",estadosolicitudavanceins)
		_, msgEstAv := solicitudavancerepository.FindOneEstadoAvance(estadosolicitudavanceins)
		if msgEstAv.Code != 0 {
			//registra estados
				msgIns := solicitudavancerepository.CreateEstadoSolicitud(estadosolicitudavanceins)
				c.JSON(200, msgIns)
			} else {
				c.JSON(200, utilidades.CheckInfo( "El Estado de la solicitud ya existe"))
			}	
		c.JSON(200, "Se registro la cancelacion de la solicitud")
	}//fin if- cancelavance
}

/*Options funcion para registrar la necesidad de una solicitud de  avance */
func CrearNecesidadAvance(c *gin.Context) {
	var necesidadins solicitudavance.Financiacionavance

	if c.Bind(&necesidadins) == nil {
		msgIns := solicitudavancerepository.CreateNecesidadAvance(necesidadins)
		c.JSON(200, msgIns)
		
	}//fin if- cancelavance
}

/*Options funcion para registrar la aprobación de una solicitud de  avance */
func CrearApruebaAvance(c *gin.Context) {

	var apruebains solicitudavance.CertificaSolicitud
	if c.Bind(&apruebains) == nil {
	 	//fmt.Println("IDS :",apruebains)
	 	financiaupd := apruebains.Presupuesto
	 	msg := solicitudavancerepository.UpdateFinancia(financiaupd)
		
		//registra estado	 
	 	var IdEst int64	
		var consultaEstado solicitudavance.Estados	
		consultaEstado.Nombre = "aprobado"
		consultaEstado.Proceso = "avances"
		resEstado, msgEst := solicitudavancerepository.FindOneEstado(consultaEstado)
		if msgEst.Code != 0 {
			IdEst=0
		} else {
			//c.JSON(200, resEstado)
			IdEst=resEstado.IdEstado
		}
		avanceins := apruebains.Estadosolicitud
		//validacion y registro de estados de la solicitud de avance
		var estadosolicitudavanceins solicitudavance.Estadosolicitudavance
		estadosolicitudavanceins.IdSolicitud = avanceins.IdSolicitud
		estadosolicitudavanceins.IdEstado=IdEst
		estadosolicitudavanceins.Observaciones=avanceins.Observaciones
		estadosolicitudavanceins.Usuario = avanceins.Usuario
		//fmt.Println("CAN :",estadosolicitudavanceins)
		_, msgEstAv := solicitudavancerepository.FindOneEstadoAvance(estadosolicitudavanceins)
		if msgEstAv.Code != 0 {
			//registra estados
				//msgIns := solicitudavancerepository.CreateEstadoSolicitud(estadosolicitudavanceins)
				solicitudavancerepository.CreateEstadoSolicitud(estadosolicitudavanceins)
				//c.JSON(200, msgIns)
			} else {
				//c.JSON(200, utilidades.CheckInfo( "El Estado de la solicitud ya existe"))
			}	

		c.JSON(200, msg)
	}//fin if- apruebaavance
}