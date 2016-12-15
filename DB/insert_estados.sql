INSERT INTO tesoreria.estados(
id_estado, nombre, descripcion, estado, proceso)
VALUES (0,'Sin Estado','No aplica estado para el proceso','A',''),
(1,'Inactivo','registro inactivo','A',''),
(2,'Activo','Registro Activo','A',''),
(3,'Registrado','Estado que aplica cuando el dato ingresa a la base de datos.','A','avances'),
(4,'Cancelado','Proceso cancelado','A','avances'),
(5,'Verificado','Estado que aplica cuando se verifican los requisitos de la solicitud de avance','A','avances'),
(6,'Aprobado','Estado que aplica cuando se aprueba la solicitud de avance, para continuar el proceso para giro.','A','avances'),
(7,'Girado','Estado que aplica cuando se hace efectivo el Giro al beneficiario del avance.','A','avances'),
(8,'Legalizado','Estado que aplica cuando se hace efectivo el proceso de legalización del avance por parte del beneficiario del avance.','A','avances');
