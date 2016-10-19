package solicitudavance


type SolicitudGeneral struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdSolicitud     	int64   `db:"id_solicitud" json:"IdSolicitud" `
    IdBeneficiario  	int64   `db:"id_beneficiario" json:"IdBeneficiario" `
    Vigencia        	string  `db:"vigencia" json:"Vigencia"`
    Consecutivo        	string  `db:"consecutivo" json:"Consecutivo"`
    Objetivo        	string  `db:"objetivo" json:"Objetivo"`
    Justificacion   	string  `db:"justificacion" json:"Justificacion"`
    ValorTotal      	float64 `db:"valor_total" json:"ValorTotal" `
    CodigoDependencia 	string  `db:"codigo_dependencia" json:"CodigoDependencia"`
    Dependencia 		string  `db:"dependencia" json:"Dependencia"`
    CodigoFacultad  	string  `db:"codigo_facultad" json:"CodigoFacultad"`
    Facultad 			string  `db:"facultad" json:"Facultad"`
    CodigoProyectoCur 	string  `db:"codigo_proyecto_curricular" json:"CodigoProyectoCur"`
    ProyectoCurricular	string  `db:"proyecto_curricular" json:"ProyectoCurricular"`
    CodigoConvenio  	string  `db:"codigo_convenio" json:"CodigoConvenio"`
    Convenio			string  `db:"convenio" json:"Convenio"`
    CodigoProyectoInv  	string  `db:"codigo_proyecto_inv" json:"CodigoProyectoInv"`
    ProyectoInv			string  `db:"proyecto_inv" json:"ProyectoInv"`
    EstadoActual 	    string  `db:"estado_actual" json:"EstadoActual"`
    IdEstado            int64   `db:"id_estado" json:"IdEstado" `
    FechaRegistro       string  `db:"fecha_registro" json:"FechaRegistro"`
    Observaciones       string  `db:"observaciones" json:"Observaciones"`
    Usuario             string  `db:"usuario" json:"Usuario"`
    IdTipo              int64   `db:"id" json:"IdTipo" `
    Descripcion         string  `db:"descripcion" json:"Descripcion"`
    Valor       	    float64 `db:"valor" json:"Valor" `
    Nombre              string  `db:"nombres" json:"Nombre"`
    Apellido            string  `db:"apellidos" json:"Apellido"`
    TipoDocumento       string  `db:"tipo_documento" json:"TipoDocumento"`
    Documento           string  `db:"documento" json:"Documento"`
    LugarDocumento      string  `db:"lugar_documento" json:"LugarDocumento"`
    Direccion           string  `db:"direccion" json:"Direccion"`
    Correo              string  `db:"correo" json:"Correo"`
    Telefono            string  `db:"telefono" json:"Telefono"`
    Celular             string  `db:"celular" json:"Celular"`
}


type Solicitud struct {
        Solicitud Solicitudavance `json:"Solicitud"`
        Beneficiario Beneficiario `json:"Beneficiario"`
        Tipoavance Solicitudtipoavance `json:"Tipoavance"`
        Estadosolicitud Estadosolicitudavance `json:"Estadosolicitud"`
    }

type Consecutivoavance struct {
    // db tag lets you specify the column name if it differs from the struct field
    Vigencia            string  `db:"vigencia" json:"Vigencia"`
    Consecutivo         string  `db:"consecutivo" json:"Consecutivo"`
}    

type Solicitudavance struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdSolicitud     	int64   `db:"id_solicitud" json:"IdSolicitud" `
    IdBeneficiario  	int64   `db:"id_beneficiario" json:"IdBeneficiario" `
    Vigencia        	string  `db:"vigencia" json:"Vigencia"`
    Consecutivo        	string  `db:"consecutivo" json:"Consecutivo"`
    Objetivo        	string  `db:"objetivo" json:"Objetivo"`
    Justificacion   	string  `db:"justificacion" json:"Justificacion"`
    ValorTotal      	float64 `db:"valor_total" json:"ValorTotal" `
    CodigoDependencia 	string  `db:"codigo_dependencia" json:"CodigoDependencia"`
    Dependencia 		string  `db:"dependencia" json:"Dependencia"`
    CodigoFacultad  	string  `db:"codigo_facultad" json:"CodigoFacultad"`
    Facultad 			string  `db:"facultad" json:"Facultad"`
    CodigoProyectoCur 	string  `db:"codigo_proyecto_curricular" json:"CodigoProyectoCur"`
    ProyectoCurricular	string  `db:"proyecto_curricular" json:"ProyectoCurricular"`
    CodigoConvenio  	string  `db:"codigo_convenio" json:"CodigoConvenio"`
    Convenio			string  `db:"convenio" json:"Convenio"`
    CodigoProyectoInv  	string  `db:"codigo_proyecto_inv" json:"CodigoProyectoInv"`
    ProyectoInv			string  `db:"proyecto_inv" json:"ProyectoInv"`
    Estado 		        string  `db:"estado" json:"Estado"`

}

type Estadosolicitudavance struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdEstado        int64   `db:"id_estado" json:"IdEstado" `
    IdSolicitud     int64   `db:"id_solicitud" json:"IdSolicitud" `
    FechaRegistro   string  `db:"fecha_registro" json:"FechaRegistro"`
    Observaciones   string  `db:"observaciones" json:"Observaciones"`
    Usuario         string  `db:"usuario" json:"Usuario"`
    Estado          string  `db:"estado" json:"Estado"`
}

type Solicitudtipoavance struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdSolicitud     int64   `db:"id_solicitud" json:"IdSolicitud" `
    IdTipo          int64   `db:"id_tipo" json:"IdTipo" `
    Descripcion     string  `db:"descripcion" json:"Descripcion"`
    Valor       	float64 `db:"valor" json:"Valor" `
    Estado          string  `db:"estado" json:"Estado"`
    Referencia      string  `db:"referencia" json:"Referencia"`
    Nombre          string  `db:"nombre" json:"Nombre"`
 }


type Beneficiario struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdBeneficiario  int64   `db:"id_beneficiario" json:"IdBeneficiario" `
    Nombre          string  `db:"nombres" json:"Nombre"`
    Apellido        string  `db:"apellidos" json:"Apellido"`
    TipoDocumento   string  `db:"tipo_documento" json:"TipoDocumento"`
    Documento       string  `db:"documento" json:"Documento"`
    LugarDocumento  string  `db:"lugar_documento" json:"LugarDocumento"`
    Direccion       string  `db:"direccion" json:"Direccion"`
    Correo          string  `db:"correo" json:"Correo"`
    Telefono        string  `db:"telefono" json:"Telefono"`
    Celular          string  `db:"celular" json:"Celular"`
	Estado          string  `db:"estado" json:"Estado"`
}

type Estados struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdEstado        int64   `db:"id_estado" json:"IdEstado" `
    Nombre          string  `db:"nombre" json:"Nombre"`
    Descripcion     string  `db:"descripcion" json:"Descripcion"`
    Estado          string  `db:"estado" json:"Estado"`
    Proceso         string  `db:"proceso" json:"Proceso"`
}

/*
func GetNewTipoavance (Id_tipoAux int64, ReferenciaAux string, NombreAux string,DescripcionAux string, EstadoAux string, Fecha_registroAux string) Tipoavance{
	return Tipoavance{Id_tipo:Id_tipoAux,  Referencia:ReferenciaAux , Nombre:NombreAux ,Descripcion:DescripcionAux , Estado:EstadoAux string, Fecha_registro:Fecha_registroAux}
}*/
