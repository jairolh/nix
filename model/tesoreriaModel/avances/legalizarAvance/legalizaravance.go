package legalizaravance

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
    FechaEstado         string  `db:"fecha_estado" json:"FechaEstado"`
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
    InternoRubro        int64   `db:"interno_rubro" json:"InternoRubro" `
    NombreRubro         string  `db:"nombre_rubro" json:"NombreRubro" `
    UnidadEjecutora     string  `db:"unidad_ejecutora" json:"UnidadEjecutora"`
    NumeroNecesidad     int64   `db:"necesidad" json:"NumeroNecesidad"`
    FechaNecesidad      string  `db:"fecha_necesidad" json:"FechaNecesidad"`
    ValorNecesidad      float64 `db:"valor_necesidad" json:"ValorNecesidad" `
    Objeto              string  `db:"objeto" json:"Objeto"`
    Disponibilidad      int64   `db:"disponibilidad" json:"Disponibilidad"`
    FechaDisp           string  `db:"fecha_disp" json:"FechaDisp"`
    ValorDisp           float64 `db:"valor_disp" json:"ValorDisp" `
    Registro            int64   `db:"registro" json:"Registro"`
    FechaRegistro       string  `db:"fecha_registro" json:"FechaRegistro"`
    ValorRegistro       float64 `db:"valor_registro" json:"ValorRegistro" `
    Compromiso          int64   `db:"compromiso" json:"Compromiso"`
    OrdenPago           int64   `db:"orden_pago" json:"OrdenPago"`
    FechaOrden          string  `db:"fecha_orden" json:"FechaOrden"`
    ValorOrden          float64 `db:"valor_orden" json:"ValorOrden" `
    FechaCertificacion  string  `db:"fecha_certifica_giro" json:"FechaCertificacion"`
}

/*
func GetNewTipoavance (Id_tipoAux int64, ReferenciaAux string, NombreAux string,DescripcionAux string, EstadoAux string, Fecha_registroAux string) Tipoavance{
	return Tipoavance{Id_tipo:Id_tipoAux,  Referencia:ReferenciaAux , Nombre:NombreAux ,Descripcion:DescripcionAux , Estado:EstadoAux string, Fecha_registro:Fecha_registroAux}
}*/
