package requisitotipoavancerepository

import(

//"fmt"
"time"

"gopkg.in/gorp.v1"
"nix/repository"
"nix/model"
"nix/model/tesoreriaModel/avances/requisitoTipoAvance"
"nix/model/tesoreriaModel/avances/requisitos"
"nix/utilidades"
)

var connectionDB *gorp.DbMap

func Init() {
	connectionDB = repository.GetConnectionDB()
}

func Create(registroIns requisitotipoavance.RequisitoTipoavance) model.MessageReturn{

	registroIns.FechaRegistro= time.Now().Format("2006-01-02 15:04:05")
	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.requisito_tipo_avance "
    consultaIns = consultaIns+" (id_tipo, id_req, estado, fecha_registro) "
    consultaIns = consultaIns+"  VALUES ($1, $2,'A',$3) "
	_, err := connectionDB.Exec(consultaIns, registroIns.IdTipo, registroIns.IdReq, registroIns.FechaRegistro)
	//fmt.Println("res :%s ",&resultado)
	//IdTipo,_ := resultado.LastInsertId()
	//fmt.Println("ID :",IdTipo)
	msg := utilidades.CheckErr(err, "Error Insertando el Registro")
	if msg.Code == 0{
		return utilidades.CheckInfo(" Se almaceno el Registro exitosamente.")
	 }
	return msg
}


func Update(registroUpd requisitotipoavance.RequisitoTipoavance) model.MessageReturn {
	var consultaUpd string 
	consultaUpd = "UPDATE tesoreria.requisito_tipo_avance "
	consultaUpd = consultaUpd+" SET estado=$3 "
    consultaUpd = consultaUpd+" WHERE id_tipo=$1 "
    consultaUpd = consultaUpd+" AND id_req=$2"
	//fmt.Println("con  :",consultaUpd)
	//fmt.Println("dat :",consultaUpd, requisitoUpd.Referencia, requisitoUpd.Nombre, requisitoUpd.Descripcion,requisitoUpd.Estado,requisitoUpd.Etapa,requisitoUpd.IdReq)
	_, err := connectionDB.Exec(consultaUpd, registroUpd.IdTipo, registroUpd.IdReq, registroUpd.Estado)
	msg := utilidades.CheckErr(err, "Error Actualizando el registro")

	if msg.Code == 0{
		return utilidades.CheckInfo(" Se Actualizo el registro exitosamente.")
	}

	return msg
}


/*

func Delete(id int64) model.MessageReturn{

	//consulta primero para ver si el registro existe
	var requisitoDel requisito.Requisito
	var consulta,consultaDel string
	consulta = "SELECT * FROM tesoreria.requisito_avance req where req.id=$1"
	err := connectionDB.SelectOne(&requisitoDel, consulta, id)
    msg := utilidades.CheckErr(err, "No se encontro registro Pre eliminacion")

    if msg.Code != 0{
		return msg
	}
    //se elimina registro
    consultaDel = "DELETE FROM tesoreria.requisito_avance req where req.id=$1"
    //fmt.Println("DEL :",consultaDel)
	_, err = connectionDB.Exec(consultaDel, id)
	//_, err = connectionDB.Delete(&requisitoDel)
    msg = utilidades.CheckErr(err, "Error Eliminando registro")
    if msg.Code != 0{
		return msg
	}else{
		return utilidades.CheckInfo(" Se Elimino el requisito exitosamente.")
	}

}
*/
func FindOne(idtipo int64, idreq int64) (requisitotipoavance.RequisitoTipoavance, model.MessageReturn) {
	var requisitoAvance requisitotipoavance.RequisitoTipoavance
    var consulta string
	consulta = "SELECT DISTINCT reqav.id_tipo, reqav.id_req, reqav.estado, reqav.fecha_registro,"
    consulta = consulta+" tav.referencia referenciaavn, tav.nombre nombreavn,"
    consulta = consulta+" req.referencia referenciareq, req.nombre nombrereq, req.descripcion descripcionreq, req.etapa etapareq"
    consulta = consulta+" FROM tesoreria.requisito_tipo_avance reqav"
    consulta = consulta+" INNER JOIN tesoreria.tipo_avance tav ON reqav.id_tipo=tav.id"
    consulta = consulta+" INNER JOIN tesoreria.requisito_avance req ON reqav.id_tipo=req.id"
    consulta = consulta+" WHERE reqav.id_tipo=$1"
    consulta = consulta+" AND reqav.id_req=$2"
    consulta = consulta+" ORDER BY tav.referencia,req.referencia"

   //fmt.Println("dat :",consulta,idtipo,idreq)
	err := connectionDB.SelectOne(&requisitoAvance, consulta, idtipo,idreq)
	msg := utilidades.CheckErr(err, "Error consultando el requisito por tipo de avance")

	return requisitoAvance, msg
}

func FindAll(idtipo int64) ([]requisitotipoavance.RequisitoTipoavance,model.MessageReturn) {
	
    var requisitosAvance []requisitotipoavance.RequisitoTipoavance
	var consulta string
    //fmt.Println("fnd :",idtipo)
	consulta = "SELECT DISTINCT reqav.id_tipo, reqav.id_req, reqav.estado, reqav.fecha_registro,"
    consulta = consulta+" tav.referencia referenciaavn, tav.nombre nombreavn,"
    consulta = consulta+" req.referencia referenciareq, req.nombre nombrereq, req.descripcion descripcionreq, req.etapa etapareq"
    consulta = consulta+" FROM tesoreria.requisito_tipo_avance reqav"
    consulta = consulta+" INNER JOIN tesoreria.tipo_avance tav ON reqav.id_tipo=tav.id"
    consulta = consulta+" INNER JOIN tesoreria.requisito_avance req ON reqav.id_req=req.id"
    consulta = consulta+" WHERE reqav.id_tipo=$1"
    consulta = consulta+" ORDER BY tav.referencia,req.referencia"
	//fmt.Println("dat :",consulta,idtipo)
  	_, err := connectionDB.Select(&requisitosAvance, consulta,idtipo)
	msg := utilidades.CheckErr(err, "Error consultando la tabla requisitos de avance")
	return requisitosAvance, msg

}


func FindSelected(idtipo int64) ([]requisito.Requisito,model.MessageReturn) {
	
    var requisitos []requisito.Requisito
	var consulta string
    
	consulta = "SELECT req.id, req.referencia, req.nombre, req.descripcion, req.estado, req.etapa, req.fecha_registro"
    consulta = consulta+" FROM tesoreria.requisito_avance req"
    consulta = consulta+" WHERE req.estado = 'A' "
    consulta = consulta+" AND req.id NOT IN "
    consulta = consulta+" (SELECT DISTINCT ra.id_req FROM tesoreria.requisito_tipo_avance ra WHERE ra.id_tipo=$1) "
  	consulta = consulta+" ORDER BY req.referencia"
  	_, err := connectionDB.Select(&requisitos, consulta,idtipo)
	msg := utilidades.CheckErr(err, "Error consultando la tabla requisitos")
	return requisitos, msg

}

