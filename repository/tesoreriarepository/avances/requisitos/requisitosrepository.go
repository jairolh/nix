package requisitorepository

import(

//"fmt"
"time"

"gopkg.in/gorp.v1"
"nix/repository"
"nix/model"
"nix/model/tesoreriaModel/avances/requisitos"
"nix/utilidades"
)

var connectionDB *gorp.DbMap

func Init() {
	connectionDB = repository.GetConnectionDB()
}

func Create(requisitoIns requisito.Requisito) model.MessageReturn{

	requisitoIns.FechaRegistro= time.Now().Format("2006-01-02 15:04:05")
	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.requisito_avance(id,referencia, nombre, descripcion, estado,etapa, fecha_registro) VALUES ( DEFAULT,$1, $2,$3,'A',$4,$5)  RETURNING id"
	_, err := connectionDB.Exec(consultaIns, requisitoIns.Referencia, requisitoIns.Nombre, requisitoIns.Descripcion,requisitoIns.Etapa,requisitoIns.FechaRegistro)
	//fmt.Println("res :%s ",&resultado)
	//IdTipo,_ := resultado.LastInsertId()
	//fmt.Println("ID :",IdTipo)
		
	msg := utilidades.CheckErr(err, "Error Insertando el Requisito")

	if msg.Code == 0{
		return utilidades.CheckInfo(" Se creo el Requisito de avance "+requisitoIns.Referencia+" "+requisitoIns.Nombre+" exitosamente.")
	 }
	
	return msg
}


func Update(requisitoUpd requisito.Requisito) model.MessageReturn {
	var consultaUpd string 
	consultaUpd = "UPDATE tesoreria.requisito_avance "
	consultaUpd = consultaUpd+"SET referencia=$1, nombre=$2, descripcion=$3, estado=$4, etapa=$5 "
	consultaUpd = consultaUpd+"WHERE  id=$6"
	//fmt.Println("con  :",consultaUpd)
	//fmt.Println("dat :",consultaUpd, requisitoUpd.Referencia, requisitoUpd.Nombre, requisitoUpd.Descripcion,requisitoUpd.Estado,requisitoUpd.Etapa,requisitoUpd.IdReq)
	_, err := connectionDB.Exec(consultaUpd, requisitoUpd.Referencia, requisitoUpd.Nombre, requisitoUpd.Descripcion,requisitoUpd.Estado,requisitoUpd.Etapa,requisitoUpd.IdReq)
	msg := utilidades.CheckErr(err, "Error Actualizando el requisito")

	if msg.Code == 0{
		return utilidades.CheckInfo(" Se Actualizo el requisito exitosamente.")
	}

	return msg
}




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

func FindOne( id int64) (requisito.Requisito, model.MessageReturn) {
	var requisito requisito.Requisito
    var consulta string
	consulta = "SELECT req.id, req.referencia, req.nombre, req.descripcion, req.estado, req.etapa, req.fecha_registro"
    consulta = consulta+" FROM tesoreria.requisito_avance req"
    consulta = consulta+" WHERE req.id=$1"
  	consulta = consulta+" ORDER BY req.referencia"

	err := connectionDB.SelectOne(&requisito, consulta, id)
	msg := utilidades.CheckErr(err, "Error consultando el tipo de avance por ID")

	return requisito, msg
}

func FindAll() ([]requisito.Requisito, model.MessageReturn) {
	var requisitos []requisito.Requisito
	var consulta string
	consulta = "SELECT req.id, req.referencia, req.nombre, req.descripcion, req.estado, req.etapa, req.fecha_registro"
    consulta = consulta+" FROM tesoreria.requisito_avance req"
  	consulta = consulta+" ORDER BY req.referencia"
  	_, err := connectionDB.Select(&requisitos, consulta)
	msg := utilidades.CheckErr(err, "Error consultando la tabla requisitos")
	return requisitos, msg

}
