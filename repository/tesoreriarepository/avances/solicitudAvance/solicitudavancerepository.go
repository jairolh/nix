package solicitudavancerepository

import(

_"fmt"
"time"

"gopkg.in/gorp.v1"
"nix/repository"
"nix/model"
"nix/model/tesoreriaModel/avances/solicitudAvance"
"nix/utilidades"
)

var connectionDB *gorp.DbMap

func Init() {
	connectionDB = repository.GetConnectionDB()
}

func CreateSolicitud(solicitudavanceIns solicitudavance.Solicitudavance) model.MessageReturn{

	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.solicitud_avance"
    consultaIns = consultaIns+" (id_solicitud, id_beneficiario, vigencia, consecutivo, "
    consultaIns = consultaIns+" objetivo, justificacion, valor_total, codigo_dependencia, "
    consultaIns = consultaIns+" dependencia, codigo_facultad, facultad, codigo_proyecto_curricular, proyecto_curricular,"
    consultaIns = consultaIns+" codigo_convenio, convenio, codigo_proyecto_inv, proyecto_inv, estado)"
    consultaIns = consultaIns+" VALUES (DEFAULT,$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,'A')  RETURNING id_solicitud";

	_, err := connectionDB.Exec(consultaIns,
	solicitudavanceIns.IdBeneficiario, solicitudavanceIns.Vigencia, solicitudavanceIns.Consecutivo, solicitudavanceIns.Objetivo, solicitudavanceIns.Justificacion,
    solicitudavanceIns.ValorTotal, solicitudavanceIns.CodigoDependencia, solicitudavanceIns.Dependencia, solicitudavanceIns.CodigoFacultad,
    solicitudavanceIns.Facultad, solicitudavanceIns.CodigoProyectoCur, solicitudavanceIns.ProyectoCurricular, solicitudavanceIns.CodigoConvenio,
    solicitudavanceIns.Convenio, solicitudavanceIns.CodigoProyectoInv, solicitudavanceIns.ProyectoInv)
		
	msg := utilidades.CheckErr(err, "Error Insertando la solicitud de avance")

	if msg.Code == 0{
		return utilidades.CheckInfo(" Se creo la solicitud "+solicitudavanceIns.Vigencia+"-"+solicitudavanceIns.Consecutivo+" exitosamente.")
	 }
	
	return msg
}

func CreateSolicitudTipo(solicitudtipoavanceIns solicitudavance.Solicitudtipoavance) model.MessageReturn{

	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.solicitud_tipo_avance"
    consultaIns = consultaIns+"( id_solicitud, id_tipo, descripcion, valor, estado) "
    consultaIns = consultaIns+" VALUES ($1,$2,$3,$4,'A') ";
	//fmt.Printf("con :",consultaIns)
	_, err := connectionDB.Exec(consultaIns, solicitudtipoavanceIns.IdSolicitud, solicitudtipoavanceIns.IdTipo, solicitudtipoavanceIns.Descripcion, solicitudtipoavanceIns.Valor)
	msg := utilidades.CheckErr(err, "Error Insertando el Tipo de avance a la solicitud")
	if msg.Code == 0{
		return utilidades.CheckInfo(" Se registro el tipo de avance a la solicitud exitosamente.")
		}
	return msg
}

func CreateEstadoSolicitud(estadosolicitudavanceIns solicitudavance.Estadosolicitudavance) model.MessageReturn{

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


func CreateBeneficiarioAvance(beneficiarioavanceIns solicitudavance.Beneficiario) model.MessageReturn{
	//err := connectionDB.Insert(&tipoavance)
	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.beneficiario"
    consultaIns = consultaIns+"(id_beneficiario, nombres, apellidos, tipo_documento, documento, "
    consultaIns = consultaIns+" lugar_documento, direccion, correo, telefono, celular, estado)"
    consultaIns = consultaIns+" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,'A') ";
	//resultado, err := connectionDB.Exec(consultaIns, tipoavanceIns.Referencia, tipoavanceIns.Nombre, tipoavanceIns.Descripcion,tipoavanceIns.FechaRegistro)
	_, err := connectionDB.Exec(consultaIns,beneficiarioavanceIns.IdBeneficiario,beneficiarioavanceIns.Nombre ,beneficiarioavanceIns.Apellido ,beneficiarioavanceIns.TipoDocumento ,beneficiarioavanceIns.Documento ,beneficiarioavanceIns.LugarDocumento ,beneficiarioavanceIns.Direccion,beneficiarioavanceIns.Correo ,beneficiarioavanceIns.Telefono ,beneficiarioavanceIns.Celular)
	msg := utilidades.CheckErr(err, "Error Insertando el beneficiario del avance")
	if msg.Code == 0{
		return utilidades.CheckInfo(" Se registro el beneficiario de la solicitud de avance exitosamente.")
		}
	return msg
}
/*
func Update(tipoavanceUpd tipoavance.Tipoavance) model.MessageReturn {
	//_, err := connectionDB.Update(&tipoavance)
	var consultaUpd string 
    consultaUpd = "UPDATE tesoreria.tipo_avance SET referencia=$1, nombre=$2, descripcion=$3, estado=$4 WHERE id=$5"
    _, err := connectionDB.Exec(consultaUpd, tipoavanceUpd.Referencia, tipoavanceUpd.Nombre, tipoavanceUpd.Descripcion,tipoavanceUpd.Estado,tipoavanceUpd.IdTipo)
	msg := utilidades.CheckErr(err, "Error Actualizando el tipo de avance")

	if msg.Code == 0{
		return utilidades.CheckInfo(" Se Actualizo el tipo de avance exitosamente.")
	}

	return msg
}




func Delete(id int64) model.MessageReturn{

	//consulta primero para ver si el registro existe
	var tipoavanceDel tipoavance.Tipoavance
	var consulta,consultaDel string
	consulta = "SELECT * FROM tesoreria.tipo_avance tav where tav.id=$1"
	err := connectionDB.SelectOne(&tipoavanceDel, consulta, id)
    msg := utilidades.CheckErr(err, "No se encontro registro Pre eliminacion")

    if msg.Code != 0{
		return msg
	}

    //se elimina registro
    consultaDel = "DELETE FROM tesoreria.tipo_avance tav where tav.id=$1"
    //fmt.Println("DEL :",consultaDel)
	_, err = connectionDB.Exec(consultaDel, id)
	//_, err = connectionDB.Delete(&tipoavanceDel)
    msg = utilidades.CheckErr(err, "Error Eliminando registro")

    if msg.Code != 0{
		return msg
	}else{
		return utilidades.CheckInfo(" Se Elimino el tipo de avance exitosamente.")
	}

}

func FindOne( id int64) (tipoavance.Tipoavance, model.MessageReturn) {
	var tipoavance tipoavance.Tipoavance
    var consulta string
	consulta = "SELECT * FROM tesoreria.tipo_avance tav where tav.id=$1"
	err := connectionDB.SelectOne(&tipoavance, consulta, id)
	msg := utilidades.CheckErr(err, "Error consultando el tipo de avance por ID")

	return tipoavance, msg
}
*/

func FindAll(vigencia int64) ([]solicitudavance.SolicitudGeneral, model.MessageReturn) {
	var solicitudesavance []solicitudavance.SolicitudGeneral
	var consulta string
	
	consulta = "SELECT sol.id_solicitud, sol.id_beneficiario, sol.vigencia, sol.consecutivo, sol.objetivo, "
    consulta = consulta+" sol.justificacion, sol.valor_total, sol.codigo_dependencia, sol.dependencia, "
    consulta = consulta+" sol.codigo_facultad, sol.facultad, sol.codigo_proyecto_curricular, sol.proyecto_curricular, "
    consulta = consulta+" sol.codigo_convenio, sol.convenio, sol.codigo_proyecto_inv, sol.proyecto_inv,"
    consulta = consulta+" est_av.id_estado, est_av.fecha_registro, est.nombre estado_actual,"
    consulta = consulta+" bene.id_beneficiario, bene.nombres, bene.apellidos, bene.tipo_documento, bene.documento"
    consulta = consulta+" FROM tesoreria.solicitud_avance sol"
    consulta = consulta+" INNER JOIN tesoreria.beneficiario bene ON bene.id_beneficiario=sol.id_beneficiario"
    consulta = consulta+" INNER JOIN tesoreria.estado_avance est_av ON est_av.id_solicitud=sol.id_solicitud"
    consulta = consulta+" INNER JOIN tesoreria.estados est ON est.id_estado=est_av.id_estado AND fecha_registro=(SELECT MAX(fecha_registro) FROM tesoreria.estado_avance WHERE id_solicitud=est_av.id_solicitud)"
    consulta = consulta+" WHERE "
    consulta = consulta+" sol.vigencia=$1 "
    consulta = consulta+" ORDER BY sol.vigencia DESC, sol.consecutivo DESC"
    //fmt.Println("dat :",consulta,vigencia)
	_, err := connectionDB.Select(&solicitudesavance, consulta,vigencia)

	msg := utilidades.CheckErr(err, "Error consultando la solicitud de avance")
	return solicitudesavance, msg

}


func FindOneBeneficiario(beneficiario solicitudavance.Beneficiario) (solicitudavance.Beneficiario, model.MessageReturn) {	
	
	idBene:=beneficiario.IdBeneficiario
	doc:=beneficiario.Documento
	
	var beneficiarioavance solicitudavance.Beneficiario
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

func FindOneSolicitudSec(solicitud solicitudavance.Solicitudavance) (solicitudavance.Solicitudavance, model.MessageReturn) {	

    var solicitudavance solicitudavance.Solicitudavance
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

func FindOneSolicitudTipoAvance(tipoavance solicitudavance.Solicitudtipoavance) (solicitudavance.Solicitudtipoavance, model.MessageReturn) {	

	var solicitudtipoavance solicitudavance.Solicitudtipoavance
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


func FindOneEstadoAvance(estadosolicitud solicitudavance.Estadosolicitudavance) (solicitudavance.Estadosolicitudavance, model.MessageReturn) {	

	var estadosolicitudes solicitudavance.Estadosolicitudavance
	var consulta string
  	consulta = " SELECT est_av.id_estado, est_av.id_solicitud, est_av.fecha_registro, est_av.observaciones, est_av.usuario, est_av.estado"
    consulta = consulta+" FROM tesoreria.estado_avance est_av"
	consulta = consulta+" WHERE est_av.id_solicitud=$1 "
    consulta = consulta+" AND est_av.id_estado=$2 "
    consulta = consulta+" AND est_av.estado='A' "
    consulta = consulta+" ORDER BY est_av.id_solicitud "    
    //fmt.Println("datSQL :",consulta, estadosolicitud.IdSolicitud, estadosolicitud.IdEstado)
	err := connectionDB.SelectOne(&estadosolicitudes, consulta, estadosolicitud.IdSolicitud, estadosolicitud.IdEstado)
	msg := utilidades.CheckErr(err, "No existe el Estado de la la solicitud")
	return estadosolicitudes, msg
}

func FindOneEstado(estado solicitudavance.Estados) (solicitudavance.Estados, model.MessageReturn) {	

	var estados solicitudavance.Estados
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