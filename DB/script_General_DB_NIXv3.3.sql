-- Crear Usuario

CREATE ROLE nixsgf LOGIN password '4dm1n=NIX2016'
  NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
COMMENT ON ROLE adminnix IS 'Usuario NIX, sistema Gestión Financiero.';

-- Crear Base de Datos


CREATE DATABASE "nix"
  WITH OWNER = nixsgf
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'es_CO.UTF-8'
       LC_CTYPE = 'es_CO.UTF-8'
       CONNECTION LIMIT = -1;

COMMENT ON DATABASE "nix"
  IS 'Base de Datos para sistema de Gestión Financiero';


/****************************************************************/

-- crear usuario esquema avances 
/*
CREATE ROLE avancesnix LOGIN password '4v4nc3s=NIX2016'
  NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
COMMENT ON ROLE adminnix IS 'Usuario para modulo de avances en sistema Gestión Financiero.';
*/
-- Crear esquema

CREATE SCHEMA tesoreria
  AUTHORIZATION nixsgf;

COMMENT ON SCHEMA tesoreria
  IS 'Esquema para el almacenamiento de información de los procesos de avances.';

-- crear tablas tipos de avances

CREATE TABLE tesoreria.requisito_avance (
                id serial NOT NULL,
                referencia VARCHAR(50) DEFAULT 0 NOT NULL,
                nombre VARCHAR(100) NOT NULL,
                descripcion text NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                etapa VARCHAR(16) NOT NULL,
                fecha_registro DATE NOT NULL,
                CONSTRAINT requisito_avance_pk PRIMARY KEY (id)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.requisito_avance
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.requisito_avance IS 'Almacena los requisitos para los avances.';
COMMENT ON COLUMN tesoreria.requisito_avance.id IS 'identificador unico del requisito para avance.';
COMMENT ON COLUMN tesoreria.requisito_avance.referencia IS 'Indica codigo asignado al requisito para el avance., por defecto es 0';
COMMENT ON COLUMN tesoreria.requisito_avance.nombre IS 'Indica el nombre del requisito para el avance.';
COMMENT ON COLUMN tesoreria.requisito_avance.descripcion IS 'detalla la descripcion del requisito para el avance.';
COMMENT ON COLUMN tesoreria.requisito_avance.estado IS 'Indica el estado del registro A-activo, I-inactivo';
COMMENT ON COLUMN tesoreria.requisito_avance.etapa IS 'Indica en que etapa del avance aplica el requisito, solicitud o legalizacion';
COMMENT ON COLUMN tesoreria.requisito_avance.fecha_registro IS 'fecha en que se registra el requisito';


CREATE TABLE tesoreria.tipo_avance (
                id serial NOT NULL,
                referencia VARCHAR(50) DEFAULT 0 NOT NULL,
                nombre VARCHAR(100) NOT NULL,
                descripcion text NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                fecha_registro DATE NOT NULL,
                CONSTRAINT tipo_avance_pk PRIMARY KEY (id)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.tipo_avance
  OWNER TO nixsgf ;


COMMENT ON COLUMN tesoreria.tipo_avance.id IS 'Identificador unico del tipo de avance';
COMMENT ON COLUMN tesoreria.tipo_avance.referencia IS 'Indica codigo asignado al tipo de avance, por defecto es 0';
COMMENT ON COLUMN tesoreria.tipo_avance.nombre IS 'Indica el nombre del tipo de avance.';
COMMENT ON COLUMN tesoreria.tipo_avance.descripcion IS 'detalla la descripcion del tipo de avance.';
COMMENT ON COLUMN tesoreria.tipo_avance.estado IS 'Indica el estado del registro A-activo, I-inactivo';
COMMENT ON COLUMN tesoreria.tipo_avance.fecha_registro IS 'Fecha en que se registra el tipo de avance';


CREATE TABLE tesoreria.requisito_tipo_avance (
                id_tipo INTEGER NOT NULL,
                id_req INTEGER NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
		fecha_registro VARCHAR NOT NULL,
                CONSTRAINT requisito_tipo_avance_pk PRIMARY KEY (id_tipo, id_req)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.requisito_tipo_avance
  OWNER TO nixsgf ;

COMMENT ON COLUMN tesoreria.requisito_tipo_avance.id_tipo IS 'Identificador unico del tipo de avance';
COMMENT ON COLUMN tesoreria.requisito_tipo_avance.id_req IS 'identificador unico del requisito para avance.';
COMMENT ON COLUMN tesoreria.requisito_tipo_avance.estado IS 'Indica el estado del registro A-activo, I-inactivo';
COMMENT ON COLUMN tesoreria.requisito_tipo_avance.fecha_registro IS 'Fecha en que se relaciona el requisito al tipo.';


ALTER TABLE tesoreria.requisito_tipo_avance ADD CONSTRAINT requisito_avance_requisito_tipo_avance_fk
FOREIGN KEY (id_req)
REFERENCES tesoreria.requisito_avance (id)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.requisito_tipo_avance ADD CONSTRAINT tipo_avance_requisito_tipo_avance_fk
FOREIGN KEY (id_tipo)
REFERENCES tesoreria.tipo_avance (id)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

/***************************** Tablas para solicitud de avances **************/


CREATE TABLE tesoreria.beneficiario (
                id_beneficiario INTEGER NOT NULL,
                nombres VARCHAR(100) NOT NULL,
                apellidos VARCHAR(100) NOT NULL,
                tipo_documento VARCHAR(8) NOT NULL,
                documento VARCHAR(16) NOT NULL,
                lugar_documento VARCHAR(100) NOT NULL,
                direccion VARCHAR(100) NOT NULL,
                correo VARCHAR(100) NOT NULL,
                telefono VARCHAR(32) NOT NULL,
                celular VARCHAR(24) NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                CONSTRAINT beneficiario_pk PRIMARY KEY (id_beneficiario)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.beneficiario
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.beneficiario IS 'Datos basicos de la persona que solicita el avance';
COMMENT ON COLUMN tesoreria.beneficiario.id_beneficiario IS 'Identificador unico del beneficiarios para los avances';
COMMENT ON COLUMN tesoreria.beneficiario.nombres IS 'Indica los nombres del beneficiario  del avance.';
COMMENT ON COLUMN tesoreria.beneficiario.apellidos IS 'Indica los apellidos del beneficiario  del avance.';
COMMENT ON COLUMN tesoreria.beneficiario.tipo_documento IS 'Indica el tipo de documento del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.documento IS 'Indica el numero de documento del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.lugar_documento IS 'Indica el lugar de expedición del documento del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.direccion IS 'Indica la direccion de residencia del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.correo IS 'Indica el correo electronico del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.telefono IS 'Indica el numero telefónico del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.celular IS 'Indica el numero de celular del beneficiario';
COMMENT ON COLUMN tesoreria.beneficiario.estado IS 'Indica el estado del registro A-activo, I-inactivo';


CREATE TABLE tesoreria.solicitud_avance (
                id_solicitud serial NOT NULL,
                id_beneficiario INTEGER NOT NULL,
                vigencia VARCHAR(16) NOT NULL,
                consecutivo VARCHAR(16) NOT NULL,
                objetivo text NOT NULL,
                justificacion text NOT NULL,
                valor_total DOUBLE PRECISION NOT NULL,
                codigo_dependencia VARCHAR(8) NOT NULL,
                dependencia VARCHAR(250) NOT NULL,
                codigo_facultad VARCHAR(8),
                facultad VARCHAR(250),
                codigo_proyecto_curricular VARCHAR(8) NOT NULL,
                proyecto_curricular VARCHAR(250) NOT NULL,
                codigo_convenio VARCHAR(50),
                convenio VARCHAR(250),
                codigo_convenio_inv VARCHAR(50),
                proyecto_inv VARCHAR(250),
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                CONSTRAINT solicitud_avance_pk PRIMARY KEY (id_solicitud)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.solicitud_avance
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.solicitud_avance IS 'Tabla en que se registran las solicitudes de avances.';
COMMENT ON COLUMN tesoreria.solicitud_avance.id_solicitud IS 'Identificador unico de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.id_beneficiario IS 'Identificador unico del beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.vigencia IS 'Indica la vigencia del consecutivo de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.consecutivo IS 'Indica el consecutivo de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.objetivo IS 'Indica los objetivos a alcanzar con el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.justificacion IS 'Indica la justificacion para la solicitud del avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.valor_total IS 'Indica el valor total de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.codigo_dependencia IS 'Indica el codigo de la dependencia asociada al beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.dependencia IS 'Indica el nombre de  la dependencia asociada al beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.codigo_facultad IS 'Indica el codigo de la facultad asociada al beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.facultad IS 'Indica el nombre de la facultad asociada al beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.codigo_proyecto_curricular IS 'Indica el codigo del proyecto curricular o la dependencia asociada al beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.proyecto_curricular IS 'Indica el nombre del proyecto curricular o la dependencia asociada al beneficiario que solicita el avance';
COMMENT ON COLUMN tesoreria.solicitud_avance.codigo_convenio IS 'Indica el codigo del convenio asociado al avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.convenio IS 'Indica los nombres de los convenios que se realizan para el avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.codigo_convenio_inv IS 'Indica el codigo del proyecto de investigación asociado al avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.proyecto_inv IS 'Indica los nombres del proyecto de investigación asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.solicitud_avance.estado IS 'Indica el estado del registro A-activo, I-inactivo';


CREATE TABLE tesoreria.financiacion_avance (
                id_solicitud INTEGER NOT NULL,
                interno_rubro VARCHAR(8) NOT NULL,
		nombre_rubro text NOT NULL,
                vigencia VARCHAR(8) NOT NULL,
		unidad_ejecutora VARCHAR(4) NOT NULL,
		necesidad VARCHAR(16) NOT NULL,
                fecha_necesidad VARCHAR(32) NULL,
                valor_necesidad DOUBLE PRECISION NOT NULL,
                objeto text NOT NULL,
                disponibilidad VARCHAR(16) NULL,
                fecha_disp VARCHAR(32) NULL,
                valor_disp DOUBLE PRECISION NULL,
                registro VARCHAR(16) NULL,
                fecha_registro VARCHAR(32) NULL,
                valor_registro DOUBLE PRECISION NULL,
                compromiso VARCHAR(16) NULL,
                orden_pago VARCHAR(16) NULL,
                fecha_orden VARCHAR(32) NULL,
                valor_orden DOUBLE PRECISION NULL,
                fecha_certifica_giro VARCHAR(32) NULL,
                CONSTRAINT financiacion_avance_pk PRIMARY KEY (id_solicitud, interno_rubro)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.financiacion_avance
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.financiacion_avance IS 'Indica los datos presupuestales al avance aprobado';
COMMENT ON COLUMN tesoreria.financiacion_avance.id_solicitud IS 'Identificador unico de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.interno_rubro IS 'Indica el código del rubro presupuestal asociado al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.nombre_rubro IS 'Indica el Nombre del rubro presupuestal asociado al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.vigencia IS 'Indica la vigencia presupuestal de los documentos financieros asociados al avance';
COMMENT ON COLUMN tesoreria.financiacion_avance.unidad_ejecutora IS 'Indica el la unidad ejecutora del rubro presupuestal asociado al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.necesidad IS 'Indica el consecutivo de la necesidad presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.fecha_necesidad IS 'Fecha en que se registra la necesidad presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.objeto IS 'Descripcion de la necesidad presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.valor_necesidad IS 'Indica el valor de la necesidad presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.disponibilidad IS 'Indica el consecutivo del certificado de disponibilidad presupuestal asociado al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.fecha_disp IS 'Fecha en que se elabora la disponibilidad presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.valor_disp IS 'Indica el valor de la disponibilidad presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.registro IS 'Indica el consecutivo del certificado de registro presupuestal asociado al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.fecha_registro IS 'Fecha en que se elabora el registro presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.valor_registro IS 'Indica el valor del registro presupuestal asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.compromiso IS 'Indica el consecutivo del compromiso presupuestal asociado al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.orden_pago IS 'Indica el consecutivo de la orden_pago asociada al avance aprobado.';
COMMENT ON COLUMN tesoreria.financiacion_avance.fecha_orden IS 'Fecha en que se elabora la orden_pago asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.valor_orden IS 'Indica el valor de la orden_pago asociado a la solicitud de avance.';
COMMENT ON COLUMN tesoreria.financiacion_avance.fecha_certifica_giro IS 'Fecha en que se certifica el giro e inicia el tiempo para legalizar el avance.';





CREATE TABLE tesoreria.estados (
                id_estado INTEGER NOT NULL,
                nombre VARCHAR(100) NOT NULL,
                descripcion text NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                CONSTRAINT estados_pk PRIMARY KEY (id_estado)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.estados
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.estados IS 'Se indican los estados de los avances';
COMMENT ON COLUMN tesoreria.estados.id_estado IS 'Identificador unico del estado del avance';
COMMENT ON COLUMN tesoreria.estados.nombre IS 'Indica el nombre del tipo de avance.';
COMMENT ON COLUMN tesoreria.estados.descripcion IS 'detalla la descripcion del estado del avance.';
COMMENT ON COLUMN tesoreria.estados.estado IS 'Indica el estado del registro A-activo, I-inactivo';


CREATE TABLE tesoreria.estado_avance (
                id_estado INTEGER NOT NULL,
                id_solicitud INTEGER NOT NULL,
                fecha_registro DATE NOT NULL,
                observaciones text NOT NULL,
                usuario VARCHAR(100) NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                CONSTRAINT estado_avance_pk PRIMARY KEY (id_estado, id_solicitud)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.estado_avance
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.estado_avance IS 'Se indica los estados de una solicitud de avance, en relacion al tiempo.';
COMMENT ON COLUMN tesoreria.estado_avance.id_estado IS 'Identificador unico del estado del avance';
COMMENT ON COLUMN tesoreria.estado_avance.id_solicitud IS 'Identificador unico de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.estado_avance.fecha_registro IS 'Fecha en que se registra el estado del avance';
COMMENT ON COLUMN tesoreria.estado_avance.observaciones IS 'detalla las observaciones de los estados del avance.';
COMMENT ON COLUMN tesoreria.estado_avance.usuario IS 'Indica el usuario que realiza el registro del estado del avance.';
COMMENT ON COLUMN tesoreria.estado_avance.estado IS 'Indica el estado del registro A-activo, I-inactivo';



CREATE TABLE tesoreria.solicitud_tipo_avance (
                id_solicitud INTEGER NOT NULL,
                id_tipo INTEGER NOT NULL,
                descripcion text NOT NULL,
                valor DOUBLE PRECISION NOT NULL,
                estado VARCHAR(1) DEFAULT 'A' NOT NULL,
                CONSTRAINT solicitud_tipo_avance_pk PRIMARY KEY (id_solicitud, id_tipo)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.solicitud_tipo_avance
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.solicitud_tipo_avance IS 'Se indica los tipos de avance asociados a una solicitud de avance.';
COMMENT ON COLUMN tesoreria.solicitud_tipo_avance.id_solicitud IS 'Identificador unico de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_tipo_avance.id_tipo IS 'Identificador unico del tipo de avance';
COMMENT ON COLUMN tesoreria.solicitud_tipo_avance.descripcion IS 'detalla la descripcion de la asignacion para el tipo de avance en la solicitud.';
COMMENT ON COLUMN tesoreria.solicitud_tipo_avance.valor IS 'Indica el valor asignado al tipo de avance, en la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_tipo_avance.estado IS 'Indica el estado del registro A-activo, I-inactivo';


CREATE TABLE tesoreria.solicitud_requisito_tipo_avance (
                id_tipo INTEGER NOT NULL,
                id_req INTEGER NOT NULL,
                id_solicitud INTEGER NOT NULL,
                valido VARCHAR(3),
                observaciones text,
                fecha_registro DATE NOT NULL,
                documento VARCHAR(250),
                estado VARCHAR(1) NOT NULL,
                ubicacion_doc VARCHAR(250),
                CONSTRAINT solicitud_requisito_tipo_avance_pk PRIMARY KEY (id_tipo, id_req, id_solicitud)
)WITH (OIDS=FALSE);

ALTER TABLE tesoreria.solicitud_requisito_tipo_avance
  OWNER TO nixsgf ;

COMMENT ON TABLE tesoreria.solicitud_requisito_tipo_avance IS 'Indica los requisitos del avance segun los tipos de avance y se evalua si el documento es valido para el requisito.';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.id_tipo IS 'Identificador unico del tipo de avance';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.id_req IS 'identificador unico del requisito para avance.';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.id_solicitud IS 'Identificador unico de la solicitud del avance.';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.valido IS 'Indica si el documento adjunto es valido para el requisito definido para el avance.  Valor SI,NO,N/A';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.observaciones IS 'detalla las observaciones de los estados del avance.';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.fecha_registro IS 'fecha en que se registra el requisito';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.documento IS 'Indica el nombre del documento que se adjunta para cumplir con el requisito definido para el avance.';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.estado IS 'Estado del registro';
COMMENT ON COLUMN tesoreria.solicitud_requisito_tipo_avance.ubicacion_doc IS 'En lace que Indica la direccion en que se cargo el documento que se adjunta para cumplir con el requisito definido para el avance.';


ALTER TABLE tesoreria.solicitud_avance ADD CONSTRAINT beneficiario_solicitud_avance_fk
FOREIGN KEY (id_beneficiario)
REFERENCES tesoreria.beneficiario (id_beneficiario)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.estado_avance ADD CONSTRAINT solicitud_avance_estado_avance_fk
FOREIGN KEY (id_solicitud)
REFERENCES tesoreria.solicitud_avance (id_solicitud)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.solicitud_tipo_avance ADD CONSTRAINT solicitud_avance_solicitud_tipo_avance_fk
FOREIGN KEY (id_solicitud)
REFERENCES tesoreria.solicitud_avance (id_solicitud)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.financiacion_avance ADD CONSTRAINT solicitud_avance_financiacion_avance_fk
FOREIGN KEY (id_solicitud)
REFERENCES tesoreria.solicitud_avance (id_solicitud)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.estado_avance ADD CONSTRAINT estados_estado_avance_fk
FOREIGN KEY (id_estado)
REFERENCES tesoreria.estados (id_estado)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;


ALTER TABLE tesoreria.solicitud_tipo_avance ADD CONSTRAINT tipo_avance_solicitud_tipo_avance_fk
FOREIGN KEY (id_tipo)
REFERENCES tesoreria.tipo_avance (id)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.solicitud_requisito_tipo_avance ADD CONSTRAINT solicitud_tipo_avance_solicitud_requisito_tipo_avance_fk
FOREIGN KEY (id_solicitud, id_tipo)
REFERENCES tesoreria.solicitud_tipo_avance (id_solicitud, id_tipo)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE tesoreria.solicitud_requisito_tipo_avance ADD CONSTRAINT requisito_tipo_avance_valida_requisito_avance_fk
FOREIGN KEY (id_tipo, id_req)
REFERENCES tesoreria.requisito_tipo_avance (id_tipo, id_req)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

/***Ajustes nuevos a solicitud****/

ALTER TABLE tesoreria.estados ADD COLUMN proceso character varying(100);
COMMENT ON COLUMN tesoreria.estados.proceso IS 'Indica el nombre del proceso a que corresponde el estado.';
ALTER TABLE tesoreria.estado_avance ADD COLUMN fecha_registro timestamp without time zone;

ALTER TABLE tesoreria.solicitud_avance ADD COLUMN facultad character varying(100);

