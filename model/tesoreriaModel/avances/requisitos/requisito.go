package requisito

type Requisito struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdReq          int64   `db:"id" json:"IdReq" `
    Referencia     string  `db:"referencia" json:"Referencia"`
    Nombre         string  `db:"nombre" json:"Nombre"`
    Descripcion    string  `db:"descripcion" json:"Descripcion"`
    Estado         string  `db:"estado" json:"Estado"`
    Etapa          string  `db:"etapa" json:"Etapa"`
    FechaRegistro  string  `db:"fecha_registro" json:"FechaRegistro"`

}
