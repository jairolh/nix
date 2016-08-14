package requisitotipoavance
type RequisitoTipoavance struct {
    // db tag lets you specify the column name if it differs from the struct field
    IdTipo         int64   `db:"id_tipo" json:"IdTipo" `
    IdReq          int64   `db:"id_req"  json:"IdReq" `
    Estado         string  `db:"estado"  json:"Estado"`
    FechaRegistro  string  `db:"fecha_registro" json:"FechaRegistro"`
    ReferenciaAvn  string  `db:"referenciaavn" json:"ReferenciaAvn"`
    NombreAvn      string  `db:"nombreavn" json:"NombreAvn"`
    ReferenciaReq  string  `db:"referenciareq" json:"ReferenciaReq"`
    NombreReq      string  `db:"nombrereq" json:"NombreReq"`
    DescripcionReq string  `db:"descripcionreq" json:"DescripcionReq"`
    EtapaReq       string  `db:"etapareq" json:"EtapaReq"`
}

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

