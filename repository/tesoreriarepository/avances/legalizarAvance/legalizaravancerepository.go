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
