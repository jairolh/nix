package legalizaravancerepository

import(

_"fmt"
_"time"
"gopkg.in/gorp.v1"
"nix/repository"
"nix/model"
"nix/model/tesoreriaModel/avances/legalizarAvance"
_"nix/model/tesoreriaModel/avances/requisitoTipoAvance"
"nix/utilidades"
)

var connectionDB *gorp.DbMap

func Init() {
	connectionDB = repository.GetConnectionDB()
}


/*
func CreateEstadoSolicitud(estadosolicitudavanceIns legalizaravance.Estadosolicitudavance) model.MessageReturn{
	estadosolicitudavanceIns.FechaRegistro= time.Now().Format("2006-01-02 15:04:05")
	//err := connectionDB.Insert(&tipoavance)
	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.estado_avance"
    consultaIns = consultaIns+"(id_estado, id_solicitud, fecha_registro, observaciones, usuario, estado) "
    consultaIns = consultaIns+" VALUES ($1,$2,$3,$4,$5,'A') ";
	//resultado, err := connectionDB.Exec(consultaIns, tipoavanceIns.Referencia, tipoavanceIns.Nombre, tipoavanceIns.Descripcion,tipoavanceIns.FechaRegistro)
	_, err := connectionDB.Exec(consultaIns, estadosolicitudavanceIns.IdEstado, estadosolicitudavanceIns.IdSolicitud, estadosolicitudavanceIns.FechaRegistro, estadosolicitudavanceIns.Observaciones, estadosolicitudavanceIns.Usuario)
	msg := utilidades.CheckErr(err, "Error Insertando el estado de la solicitud")
	if msg.Code == 0{
		return utilidades.CheckInfo(" Se registro el estado de la solicitud de avance exitosamente.")
		}
	return msg
}

func CreateVerificaSolicitud(verificaavanceIns legalizaravance.RequisitoSolicitudavance) model.MessageReturn{
	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.solicitud_requisito_tipo_avance "
    consultaIns = consultaIns+" (id_tipo, id_req, id_solicitud, valido, observaciones, fecha_registro,"
    consultaIns = consultaIns+"  documento, estado, ubicacion_doc) "
    consultaIns = consultaIns+" VALUES ($1,$2,$3,$4,$5,$6,$7,'A',$8) ";
	_, err := connectionDB.Exec(consultaIns,verificaavanceIns.IdTipo,verificaavanceIns.IdReq,verificaavanceIns.IdSolicitud, verificaavanceIns.Valido,
    verificaavanceIns.Observaciones, verificaavanceIns.FechaRegistroReq,verificaavanceIns.Documento,verificaavanceIns.UbicacionDoc)
	msg := utilidades.CheckErr(err, "Error Insertando la verificacion de avance")
	if msg.Code == 0{
		return utilidades.CheckInfo("Se registro la verificacion de requisito exitosamente.")
	 }

	return msg
}





func FindOneBeneficiario(beneficiario legalizaravance.Beneficiario) (legalizaravance.Beneficiario, model.MessageReturn) {	
	idBene:=beneficiario.IdBeneficiario
	doc:=beneficiario.Documento
	var beneficiarioavance legalizaravance.Beneficiario
    var consulta string
	consulta = "SELECT bene.id_beneficiario, bene.nombres, bene.apellidos, bene.tipo_documento, bene.documento,"
    consulta = consulta+" bene.lugar_documento, bene.direccion, bene.correo, bene.telefono, bene.celular, bene.estado"
    consulta = consulta+" FROM tesoreria.beneficiario bene"
    consulta = consulta+" WHERE "
    consulta = consulta+" bene.id_beneficiario=$1 "
    consulta = consulta+" AND bene.documento=$2"
    consulta = consulta+" ORDER BY bene.documento,bene.nombres, bene.apellidos"
   //fmt.Println("dat :",consulta,idtipo,idreq)
	err := connectionDB.SelectOne(&beneficiarioavance, consulta, idBene,doc)
	msg := utilidades.CheckErr(err, "Error consultando el beneficiario")
	return beneficiarioavance, msg
}

func FindOneSolicitudSec(solicitud legalizaravance.Solicitudavance) (legalizaravance.Solicitudavance, model.MessageReturn) {	
    var solicitudavance legalizaravance.Solicitudavance
	var consulta string
	consulta = "SELECT sol.id_solicitud, sol.id_beneficiario, sol.vigencia, sol.consecutivo, sol.objetivo, "
    consulta = consulta+" sol.justificacion, sol.valor_total, sol.codigo_dependencia, sol.dependencia, "
    consulta = consulta+" sol.codigo_facultad, sol.facultad, sol.codigo_proyecto_curricular, sol.proyecto_curricular, "
    consulta = consulta+" sol.codigo_convenio, sol.convenio, sol.codigo_proyecto_inv, sol.proyecto_inv, sol.estado"
    consulta = consulta+" FROM tesoreria.solicitud_avance sol"
    consulta = consulta+" WHERE sol.vigencia=$1 "
    consulta = consulta+" AND sol.consecutivo=$2 "
    consulta = consulta+" ORDER BY sol.vigencia, sol.consecutivo "
    //fmt.Println("datSQL :",consulta,solicitud.Vigencia,solicitud.Consecutivo)
	err := connectionDB.SelectOne(&solicitudavance, consulta, solicitud.Vigencia,solicitud.Consecutivo)
	msg := utilidades.CheckErr(err, "No existe la solicitud en la DB")
	return solicitudavance, msg
}

func FindOneSolicitudTipoAvance(tipoavance legalizaravance.Solicitudtipoavance) (legalizaravance.Solicitudtipoavance, model.MessageReturn) {	
	var solicitudtipoavance legalizaravance.Solicitudtipoavance
	var consulta string
	consulta = "SELECT tipo.id_solicitud, tipo.id_tipo, tipo.descripcion, tipo.valor, tipo.estado"
    consulta = consulta+" FROM tesoreria.solicitud_tipo_avance tipo"
    consulta = consulta+" WHERE tipo.id_solicitud=$1 "
    consulta = consulta+" AND tipo.id_tipo=$2 "
    consulta = consulta+" ORDER BY tipo.id_solicitud "
    //fmt.Println("datSQL :",consulta,tipoavance.IdSolicitud,tipoavance.IdTipo)
	err := connectionDB.SelectOne(&solicitudtipoavance, consulta, tipoavance.IdSolicitud, tipoavance.IdTipo)
	msg := utilidades.CheckErr(err, "No existe el tipo de avance de la solicitud")
	return solicitudtipoavance, msg
}


func FindOneEstadoAvance(estadosolicitud legalizaravance.Estadosolicitudavance) (legalizaravance.Estadosolicitudavance, model.MessageReturn) {	
	var estadosolicitudes legalizaravance.Estadosolicitudavance
	var consulta string
  	consulta = " SELECT est_av.id_estado, est_av.id_solicitud, est_av.fecha_registro, est_av.observaciones, est_av.usuario, est_av.estado, est.nombre nombre_estado"
    consulta = consulta+" FROM tesoreria.estado_avance est_av"
    consulta = consulta+" INNER JOIN tesoreria.estados est ON est.id_estado=est_av.id_estado "
	consulta = consulta+" WHERE est_av.id_solicitud=$1 "
    consulta = consulta+" AND est_av.id_estado=$2 "
    consulta = consulta+" AND est_av.estado='A' "
    consulta = consulta+" ORDER BY est_av.id_solicitud "    
    //fmt.Println("datSQL :",consulta, estadosolicitud.IdSolicitud, estadosolicitud.IdEstado)
	err := connectionDB.SelectOne(&estadosolicitudes, consulta, estadosolicitud.IdSolicitud, estadosolicitud.IdEstado)
	msg := utilidades.CheckErr(err, "No existe el Estado de la la solicitud")
	return estadosolicitudes, msg
}

func FindEstadoAvanceBeneficiario(solicitud int64, beneficiario int64) ([]legalizaravance.Estadosolicitudavance, model.MessageReturn) {	
	var estadosolicitudes []legalizaravance.Estadosolicitudavance
	var consulta string
  	consulta = " SELECT DISTINCT est_av.id_estado, est_av.id_solicitud, est_av.fecha_registro, est_av.observaciones,  "
	consulta = consulta+" est_av.usuario, est_av.estado, est.nombre nombre_estado "
	consulta = consulta+" FROM tesoreria.solicitud_avance sol   "
	consulta = consulta+" INNER JOIN tesoreria.estado_avance est_av ON est_av.id_solicitud=sol.id_solicitud  "
	consulta = consulta+" INNER JOIN tesoreria.estados est ON est.id_estado=est_av.id_estado AND fecha_registro=  "
	consulta = consulta+" (SELECT MAX(fecha_registro) FROM tesoreria.estado_avance WHERE id_solicitud=est_av.id_solicitud)   "
	consulta = consulta+" WHERE sol.id_beneficiario=$2 "
	consulta = consulta+" AND est_av.id_solicitud NOT IN ($1)  "
	consulta = consulta+" AND est_av.estado='A'  "
	consulta = consulta+" ORDER BY est_av.id_solicitud " 
	//fmt.Println("datSQL :",consulta, solicitud , beneficiario)
    _, err := connectionDB.Select(&estadosolicitudes, consulta, solicitud , beneficiario)
	msg := utilidades.CheckErr(err, "No exite registro de avances para el beneficiario")
	return estadosolicitudes, msg
}

func FindOneEstado(estado legalizaravance.Estados) (legalizaravance.Estados, model.MessageReturn) {	
	var estados legalizaravance.Estados
	var consulta string
	consulta = " SELECT id_estado, nombre, descripcion, estado, proceso"
    consulta = consulta+" FROM tesoreria.estados "
	consulta = consulta+" WHERE trim(lower(nombre)) LIKE ($1) "
    consulta = consulta+" AND trim(lower(proceso)) LIKE ($2) "
    consulta = consulta+" AND estado='A' "
    consulta = consulta+" ORDER BY id_estado "
    //fmt.Println("datSQL :",consulta,estado.Nombre, estado.Proceso)
	err := connectionDB.SelectOne(&estados, consulta, estado.Nombre, estado.Proceso)
	msg := utilidades.CheckErr(err, "No existe el estado")
	return estados, msg
}


func FindOne(vigencia int64,solicitud int64) ([]legalizaravance.SolicitudGeneral, model.MessageReturn) {
	var solicitudavance []legalizaravance.SolicitudGeneral
	var consulta string
	consulta = "SELECT sol.id_solicitud, sol.id_beneficiario, sol.vigencia, sol.consecutivo, sol.objetivo, sol.justificacion, "
    consulta = consulta+" sol.valor_total, sol.codigo_dependencia, sol.dependencia, sol.codigo_facultad, sol.facultad, "
    consulta = consulta+" sol.codigo_proyecto_curricular, sol.proyecto_curricular, sol.codigo_convenio, sol.convenio, sol.codigo_proyecto_inv, sol.proyecto_inv, "
    consulta = consulta+" est_av.id_estado, est_av.fecha_registro, est.nombre estado_actual, "
    consulta = consulta+" bene.id_beneficiario, bene.nombres, bene.apellidos, bene.tipo_documento, bene.documento,  "
    consulta = consulta+" bene.lugar_documento, bene.direccion, bene.correo, bene.telefono, bene.celular "
    consulta = consulta+" FROM tesoreria.solicitud_avance sol  "
    consulta = consulta+" INNER JOIN tesoreria.beneficiario bene ON bene.id_beneficiario=sol.id_beneficiario  "
    consulta = consulta+" INNER JOIN tesoreria.estado_avance est_av ON est_av.id_solicitud=sol.id_solicitud  "
    consulta = consulta+" INNER JOIN tesoreria.estados est ON est.id_estado=est_av.id_estado AND fecha_registro= "
    consulta = consulta+" (SELECT MAX(fecha_registro) FROM tesoreria.estado_avance WHERE id_solicitud=est_av.id_solicitud)  "
    consulta = consulta+" WHERE  sol.vigencia=$1 "
    consulta = consulta+" AND    sol.consecutivo=$2 "
    //fmt.Println("dat :",consulta,solicitud)
	_, err := connectionDB.Select(&solicitudavance, consulta,vigencia,solicitud)

	msg := utilidades.CheckErr(err, "Error consultando la solicitud de avance")
	return solicitudavance, msg

}


func FindOneFinanciaAvance(vigencia int64,solicitud int64) ([]legalizaravance.Financiacionavance, model.MessageReturn) {
	var financiaavance []legalizaravance.Financiacionavance
	var consulta string
	consulta = "SELECT id_solicitud, interno_rubro, nombre_rubro, vigencia, unidad_ejecutora, "
    consulta = consulta+"necesidad, fecha_necesidad, valor_necesidad, objeto, "
    consulta = consulta+" (CASE WHEN disponibilidad is null THEN '0' else disponibilidad END ) disponibilidad, "
    consulta = consulta+" (CASE WHEN fecha_disp is null THEN 'ND' else fecha_disp END ) fecha_disp, "
    consulta = consulta+" (CASE WHEN valor_disp is null THEN '0' else valor_disp END ) valor_disp, "
	consulta = consulta+" (CASE WHEN registro is null THEN '0' else registro END ) registro, "
	consulta = consulta+" (CASE WHEN fecha_registro is null THEN 'ND' else fecha_registro END ) fecha_registro, "
    consulta = consulta+" (CASE WHEN valor_registro is null THEN '0' else valor_registro END ) valor_registro, "
    consulta = consulta+" (CASE WHEN compromiso is null THEN '0' else compromiso END ) compromiso, "
    consulta = consulta+" (CASE WHEN orden_pago is null THEN '0' else orden_pago END ) orden_pago, "
  	consulta = consulta+" (CASE WHEN fecha_orden is null THEN 'ND' else fecha_orden END ) fecha_orden, "
    consulta = consulta+" (CASE WHEN valor_orden is null THEN '0' else valor_orden END ) valor_orden "
    consulta = consulta+"FROM tesoreria.financiacion_avance "
    consulta = consulta+"WHERE  vigencia=$1 "
    consulta = consulta+" AND  id_solicitud=$2 "
    //fmt.Println("dat :",consulta,vigencia,solicitud)
	_, err := connectionDB.Select(&financiaavance, consulta,vigencia,solicitud)

	msg := utilidades.CheckErr(err, "Error consultando la financiacion del avance")
	return financiaavance, msg

}*/

func FindAll(vigencia int64) ([]legalizaravance.SolicitudGeneral, model.MessageReturn) {
	var solicitudesavance []legalizaravance.SolicitudGeneral
	var consulta string
	consulta = "SELECT sol.id_solicitud, sol.id_beneficiario, sol.vigencia, sol.consecutivo, sol.objetivo, "
    consulta = consulta+" sol.justificacion, sol.valor_total, sol.codigo_dependencia, sol.dependencia, "
    consulta = consulta+" sol.codigo_facultad, sol.facultad, sol.codigo_proyecto_curricular, sol.proyecto_curricular, "
    consulta = consulta+" sol.codigo_convenio, sol.convenio, sol.codigo_proyecto_inv, sol.proyecto_inv, "
    consulta = consulta+" est_av.id_estado, est_av.fecha_registro fecha_estado, est.nombre estado_actual, "
    consulta = consulta+" bene.id_beneficiario, bene.nombres, bene.apellidos, bene.tipo_documento, "
    consulta = consulta+" bene.documento, bene.correo, bene.telefono, bene.celular, "
    consulta = consulta+" fin_av.interno_rubro, fin_av.nombre_rubro, fin_av.unidad_ejecutora, "
    consulta = consulta+" fin_av.necesidad, fin_av.fecha_necesidad, fin_av.valor_necesidad, fin_av.objeto, fin_av.disponibilidad, "
    consulta = consulta+" fin_av.fecha_disp, fin_av.valor_disp, fin_av.registro, fin_av.fecha_registro, fin_av.valor_registro, "
    consulta = consulta+" fin_av.compromiso, fin_av.orden_pago, fin_av.fecha_orden, fin_av.valor_orden, fin_av.fecha_certifica_giro "
    consulta = consulta+" FROM tesoreria.solicitud_avance sol"
    consulta = consulta+" INNER JOIN tesoreria.beneficiario bene ON bene.id_beneficiario=sol.id_beneficiario"
    consulta = consulta+" INNER JOIN tesoreria.estado_avance est_av ON est_av.id_solicitud=sol.id_solicitud"
    consulta = consulta+" INNER JOIN tesoreria.estados est ON est.id_estado=est_av.id_estado AND fecha_registro=(SELECT MAX(fecha_registro) FROM tesoreria.estado_avance WHERE id_solicitud=est_av.id_solicitud)"
    consulta = consulta+" INNER JOIN tesoreria.financiacion_avance fin_av ON fin_av.id_solicitud=sol.id_solicitud AND fin_av.vigencia=sol.vigencia  "    
    consulta = consulta+" WHERE "
    consulta = consulta+" sol.vigencia=$1 AND UPPER(est.nombre)='GIRADO' "
    consulta = consulta+" ORDER BY sol.vigencia DESC, sol.consecutivo::int DESC"
    //fmt.Println("dat :",consulta,vigencia)
	_, err := connectionDB.Select(&solicitudesavance, consulta,vigencia)
	msg := utilidades.CheckErr(err, "Error consultando las solicitudes de avance")
	return solicitudesavance, msg

}
/*
func FindAllTipo(solicitud  int64) ([]legalizaravance.Solicitudtipoavance, model.MessageReturn) {
	var tipoSolicitud []legalizaravance.Solicitudtipoavance
	var consulta string
	consulta = "SELECT tav.referencia,tav.nombre,sol.id_tipo, sol.id_solicitud, sol.descripcion, sol.valor, sol.estado "
    consulta = consulta+" FROM tesoreria.solicitud_tipo_avance sol "
    consulta = consulta+" INNER JOIN tesoreria.tipo_avance tav ON tav.id=sol.id_tipo "
    consulta = consulta+" WHERE sol.id_solicitud=$1 "
    //fmt.Println("dat :",consulta,solicitud)
	_, err := connectionDB.Select(&tipoSolicitud, consulta,solicitud )
	msg := utilidades.CheckErr(err, "Error consultando la solicitud de avance")
	return  tipoSolicitud, msg

}

func FindAllReq(tipo int64) ([]requisitotipoavance.Requisito, model.MessageReturn) {
	var requisitosSolicitud []requisitotipoavance.Requisito
	var consulta string
	consulta = "SELECT req.id, req.referencia, req.nombre, req.descripcion, req.etapa, req.fecha_registro, "
    consulta = consulta+" req.estado "
    consulta = consulta+" FROM tesoreria.requisito_avance req "
    consulta = consulta+" INNER JOIN tesoreria.requisito_tipo_avance reqav ON reqav.id_req=req.id and reqav.estado='A' "
    consulta = consulta+" WHERE req.estado='A'  "
    consulta = consulta+" and reqav.id_tipo IN ($1) "
    consulta = consulta+" ORDER BY req.etapa DESC, req.referencia ASC "
    //fmt.Println("dat :",consulta,tipo)
	_, err := connectionDB.Select(&requisitosSolicitud , consulta,tipo)
	msg := utilidades.CheckErr(err, "Error consultando los requisitos para la solicitud de avance")
	return requisitosSolicitud, msg

}

func FindAllReqAvn(solicitud int64,tipo int64) ([]legalizaravance.RequisitoSolicitudavance, model.MessageReturn) {
	var requisitosSolicitud []legalizaravance.RequisitoSolicitudavance
	var consulta string
	consulta = " SELECT  req.referencia referenciareq, req.nombre nombrereq, req.descripcion descripcionreq, req.etapa etapareq, req.fecha_registro,  req.estado, "
    consulta = consulta+" reqav.id_tipo, reqav.id_req, $1 id_solicitud, "
    consulta = consulta+" (CASE WHEN rqsav.valido IS NULL THEN '' ELSE rqsav.valido END ) valido, "
    consulta = consulta+" (CASE WHEN rqsav.observaciones IS NULL THEN '' ELSE rqsav.observaciones END ) observacionesreqav, "
    consulta = consulta+" (CASE WHEN rqsav.fecha_registro IS NULL THEN '2000-01-01' ELSE rqsav.fecha_registro END ) fecha_registro_reqav, "
    consulta = consulta+" (CASE WHEN rqsav.documento IS NULL THEN '' ELSE rqsav.documento END ) documento, "
    consulta = consulta+" (CASE WHEN rqsav.estado IS NULL THEN '' ELSE rqsav.estado END )estado_reqav, "
    consulta = consulta+" (CASE WHEN rqsav.ubicacion_doc IS NULL THEN '' ELSE rqsav.ubicacion_doc END ) ubicacion_doc, "
    consulta = consulta+" 'system' usuario "
    consulta = consulta+" FROM tesoreria.requisito_avance req   "
    consulta = consulta+" INNER JOIN tesoreria.requisito_tipo_avance reqav ON reqav.id_req=req.id and reqav.estado='A'  "
    consulta = consulta+" LEFT OUTER JOIN tesoreria.solicitud_requisito_tipo_avance rqsav ON rqsav.id_solicitud=$1 AND rqsav.id_tipo=reqav.id_tipo AND rqsav.id_req=reqav.id_req and rqsav.estado='A' "
    consulta = consulta+" WHERE req.estado='A' and reqav.id_tipo IN ($2)  ORDER BY req.etapa DESC, req.referencia ASC  "
    //fmt.Println("dat :",consulta,solicitud,tipo)
	_, err := connectionDB.Select(&requisitosSolicitud , consulta, solicitud, tipo)
	msg := utilidades.CheckErr(err, "Error consultando los requisitos para la solicitud de avance")
	return requisitosSolicitud, msg

}*/