package tipoavancerepository

import(

//"fmt"
"time"

"gopkg.in/gorp.v1"
"nix/repository"
"nix/model"
"nix/model/tesoreriaModel/avances/tipoAvance"
"nix/utilidades"
)

var connectionDB *gorp.DbMap

func Init() {
	connectionDB = repository.GetConnectionDB()
}

func Create(tipoavanceIns tipoavance.Tipoavance) model.MessageReturn{

	tipoavanceIns.FechaRegistro= time.Now().Format("2006-01-02 15:04:05")
	//err := connectionDB.Insert(&tipoavance)
	var consultaIns string 
	consultaIns = "INSERT INTO tesoreria.tipo_avance(id,referencia, nombre, descripcion, estado, fecha_registro) VALUES ( DEFAULT,$1, $2,$3,'A',$4)  RETURNING id"
	//fmt.Printf("con :",consulta)
	//resultado, err := connectionDB.Exec(consultaIns, tipoavanceIns.Referencia, tipoavanceIns.Nombre, tipoavanceIns.Descripcion,tipoavanceIns.FechaRegistro)
	_, err := connectionDB.Exec(consultaIns, tipoavanceIns.Referencia, tipoavanceIns.Nombre, tipoavanceIns.Descripcion,tipoavanceIns.FechaRegistro)
	//fmt.Println("res :%s ",&resultado)
	//IdTipo,_ := resultado.LastInsertId()
	//fmt.Println("ID :",IdTipo)
		
	msg := utilidades.CheckErr(err, "Error Insertando el Tipo de avance")

	if msg.Code == 0{
		return utilidades.CheckInfo(" Se creo el tipo de avance "+tipoavanceIns.Referencia+" "+tipoavanceIns.Nombre+" exitosamente.")
	 }
	
	return msg
}


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

func FindAll() ([]tipoavance.Tipoavance, model.MessageReturn) {
	var tiposavance []tipoavance.Tipoavance
	var consulta string
	//consulta = "SELECT tav.id_tipo, tav.referencia, tav.nombre, tav.descripcion, tav.estado, tav.fecha_registro FROM tesoreria.tipo_avance tav"
	consulta = "SELECT * FROM tesoreria.tipo_avance tav order by tav.referencia"
    //fmt.Println("SELECT :",consulta)
	_, err := connectionDB.Select(&tiposavance, consulta)

	msg := utilidades.CheckErr(err, "Error consultando la tabla tipos de avance")
	return tiposavance, msg

}
