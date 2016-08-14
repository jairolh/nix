package tipoavance

type Tipoavance struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdTipo         int64   `db:"id" json:"IdTipo" `
    Referencia      string  `db:"referencia" json:"Referencia"`
    Nombre          string  `db:"nombre" json:"Nombre"`
    Descripcion     string  `db:"descripcion" json:"Descripcion"`
    Estado          string  `db:"estado" json:"Estado"`
    FechaRegistro  string  `db:"fecha_registro" json:"FechaRegistro"`

}
/*
func GetNewTipoavance (Id_tipoAux int64, ReferenciaAux string, NombreAux string,DescripcionAux string, EstadoAux string, Fecha_registroAux string) Tipoavance{
	return Tipoavance{Id_tipo:Id_tipoAux,  Referencia:ReferenciaAux , Nombre:NombreAux ,Descripcion:DescripcionAux , Estado:EstadoAux string, Fecha_registro:Fecha_registroAux}
}*/
