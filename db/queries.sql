-- ALIMENTACION --

-- name: ListarAlimentaciones :many
SELECT id, nombre_local, descripcion,ubicacion, direccion, horario_apertura, horario_cierre, telefono, tipo_comida, precio_promedio, provider_id FROM alimentacion;

-- name: BuscarAlimentacionPorID :one
SELECT id, nombre_local, descripcion, ubicacion, direccion, horario_apertura, horario_cierre, telefono, tipo_comida, precio_promedio, provider_id FROM alimentacion
WHERE id = ?;

-- name: CrearAlimentacion :one
INSERT INTO alimentacion (nombre_local, descripcion, ubicacion, direccion, horario_apertura, horario_cierre, telefono, tipo_comida, precio_promedio, provider_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, nombre_local, descripcion, ubicacion, direccion, horario_apertura, horario_cierre, telefono, tipo_comida, precio_promedio, provider_id;

-- name: ActualizarAlimentacion :one
UPDATE alimentacion
SET nombre_local = ?, descripcion = ?, ubicacion = ?, direccion = ?,
    horario_apertura = ?, horario_cierre = ?, telefono = ?,
    tipo_comida = ?, precio_promedio = ?, provider_id = ?
WHERE id = ?
RETURNING id, nombre_local, descripcion, ubicacion, direccion, horario_apertura, horario_cierre, telefono, tipo_comida, precio_promedio, provider_id;

-- name: BorrarAlimentacion :execrows
DELETE FROM alimentacion WHERE id = ?;


-- MENU DIARIO --

-- name: ListarMenuDiarios :many
SELECT id, fecha, alimentacion_id FROM menu_diario;

-- name: BuscarMenuDiarioPorID :one
SELECT id, fecha, alimentacion_id FROM menu_diario
WHERE id = ?;

-- name: CrearMenuDiario :one
INSERT INTO menu_diario (fecha, alimentacion_id)
VALUES (?, ?)
RETURNING id, fecha, alimentacion_id;

-- name: ActualizarMenuDiario :one
UPDATE menu_diario
SET fecha = ?, alimentacion_id = ?
WHERE id = ?
RETURNING id, fecha, alimentacion_id;

-- name: BorrarMenuDiario :execrows
DELETE FROM menu_diario WHERE id = ?;


-- PLATOS --

-- name: ListarPlatos :many
SELECT id, nombre_plato, descripcion, categoria, precio, menu_diario_id FROM plato;

-- name: BuscarPlatoPorID :one
SELECT id, nombre_plato, descripcion, categoria, precio, menu_diario_id FROM plato
WHERE id = ?;

-- name: CrearPlato :one
INSERT INTO plato (nombre_plato, descripcion, categoria, precio, menu_diario_id)
VALUES (?, ?, ?, ?, ?)
RETURNING id, nombre_plato, descripcion, categoria, precio, menu_diario_id;

-- name: ActualizarPlato :one
UPDATE plato
SET nombre_plato = ?, descripcion = ?, categoria = ?,
    precio = ?, menu_diario_id = ?
WHERE id = ?
RETURNING id, nombre_plato, descripcion, categoria, precio, menu_diario_id;

-- name: BorrarPlato :execrows
DELETE FROM plato WHERE id = ?;


-- RESENAS --

-- name: ListarResenas :many
SELECT id, comentario, calificacion, alimentacion_id FROM resena;

-- name: BuscarResenaPorID :one
SELECT id, comentario, calificacion, alimentacion_id FROM resena
WHERE id = ?;

-- name: CrearResena :one
INSERT INTO resena (comentario, calificacion, alimentacion_id)
VALUES (?, ?, ?)
RETURNING id, comentario, calificacion, alimentacion_id;

-- name: ActualizarResena :one
UPDATE resena
SET comentario = ?, calificacion = ?, alimentacion_id = ?
WHERE id = ?
RETURNING id, comentario, calificacion, alimentacion_id;

-- name: BorrarResena :execrows
DELETE FROM resena WHERE id = ?;


-- COOPERATIVAS --

-- name: ListarCooperativas :many
SELECT id, nombre, telefono, descripcion FROM cooperativa;

-- name: BuscarCooperativaPorID :one
SELECT id, nombre, telefono, descripcion FROM cooperativa
WHERE id = ?;

-- name: CrearCooperativa :one
INSERT INTO cooperativa (nombre, telefono, descripcion)
VALUES (?, ?, ?)
RETURNING id, nombre, telefono, descripcion;

-- name: ActualizarCooperativa :one
UPDATE cooperativa
SET nombre = ?, telefono = ?, descripcion = ?
WHERE id = ?
RETURNING id, nombre, telefono, descripcion;

-- name: BorrarCooperativa :execrows
DELETE FROM cooperativa WHERE id = ?;


-- RUTAS DE TRANSPORTE --

-- name: ListarRutas :many
SELECT id, nombre_linea, frecuencia_aprox, precio, descripcion_ruta, cooperativa_id, sector_origen_id, sector_destino_id, parada_bus_id FROM ruta_transporte;

-- name: BuscarRutaPorID :one
SELECT id, nombre_linea, frecuencia_aprox, precio, descripcion_ruta, cooperativa_id, sector_origen_id, sector_destino_id, parada_bus_id FROM ruta_transporte
WHERE id = ?;

-- name: CrearRuta :one
INSERT INTO ruta_transporte (nombre_linea, frecuencia_aprox, precio, descripcion_ruta, cooperativa_id, sector_origen_id, sector_destino_id, parada_bus_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, nombre_linea, frecuencia_aprox, precio, descripcion_ruta, cooperativa_id, sector_origen_id, sector_destino_id, parada_bus_id;

-- name: ActualizarRuta :one
UPDATE ruta_transporte
SET nombre_linea = ?, frecuencia_aprox = ?, precio = ?,
    descripcion_ruta = ?, cooperativa_id = ?,
    sector_origen_id = ?, sector_destino_id = ?, parada_bus_id = ?  
WHERE id = ?
RETURNING id, nombre_linea, frecuencia_aprox, precio, descripcion_ruta, cooperativa_id, sector_origen_id, sector_destino_id, parada_bus_id;

-- name: BorrarRuta :execrows
DELETE FROM ruta_transporte WHERE id = ?;

--PARADAS BUS--

-- name: ListarParadas :many
SELECT id, nombre_parada, direccion, descripcion FROM paradas_bus;

-- name: BuscarParadaPorID :one
SELECT id, nombre_parada, direccion, descripcion FROM paradas_bus
WHERE id = ?;

-- name: CrearParada :one
INSERT INTO paradas_bus (nombre_parada, direccion, descripcion)
VALUES (?, ?, ?)
RETURNING id, nombre_parada, direccion, descripcion;

-- name: ActualizarParada :one
UPDATE paradas_bus
SET nombre_parada = ?, direccion = ?, descripcion = ?
WHERE id = ?
RETURNING id, nombre_parada, direccion, descripcion;

-- name: BorrarParada :execrows
DELETE FROM paradas_bus WHERE id = ?;

-- VIVIENDAS --

-- name: ListarViviendas :many
SELECT vivienda_id, titulo, tipo_vivienda, precio, garantia, precio_garantia, direccion, luz, agua, amueblado, internet, bano_privado, numero_habitaciones, mascotas, genero_preferido, reglas_convivencia, telefono, email, estado, comentario, sector_id, proveedor_id FROM vivienda;

-- name: BuscarViviendaPorID :one
SELECT vivienda_id, titulo, tipo_vivienda, precio, garantia, precio_garantia, direccion, luz, agua, amueblado, internet, bano_privado, numero_habitaciones, mascotas, genero_preferido, reglas_convivencia, telefono, email, estado, comentario, sector_id, proveedor_id FROM vivienda
WHERE vivienda_id = ?;

-- name: CrearVivienda :one
INSERT INTO vivienda (titulo, tipo_vivienda, precio, garantia, precio_garantia, direccion, luz, agua, amueblado, internet, bano_privado, numero_habitaciones, mascotas, genero_preferido, reglas_convivencia, telefono, email, estado, comentario, sector_id, proveedor_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING vivienda_id, titulo, tipo_vivienda, precio, garantia, precio_garantia, direccion, luz, agua, amueblado, internet, bano_privado, numero_habitaciones, mascotas, genero_preferido, reglas_convivencia, telefono, email, estado, comentario, sector_id, proveedor_id;

-- name: ActualizarVivienda :one
UPDATE vivienda
SET titulo = ?, tipo_vivienda = ?, precio = ?, garantia = ?,
    precio_garantia = ?, direccion = ?, luz = ?, agua = ?,
    amueblado = ?, internet = ?, bano_privado = ?,
    numero_habitaciones = ?, mascotas = ?, genero_preferido = ?,
    reglas_convivencia = ?, telefono = ?, email = ?,
    estado = ?, comentario = ?, sector_id = ?, proveedor_id = ?
WHERE vivienda_id = ?
RETURNING vivienda_id, titulo, tipo_vivienda, precio, garantia, precio_garantia, direccion, luz, agua, amueblado, internet, bano_privado, numero_habitaciones, mascotas, genero_preferido, reglas_convivencia, telefono, email, estado, comentario, sector_id, proveedor_id;

-- name: EliminarVivienda :execrows
DELETE FROM vivienda WHERE vivienda_id = ?;


-- FOTOS --

-- name: ListarFotos :many
SELECT foto_id, url, vivienda_id FROM foto;

-- name: BuscarFotoPorID :one
SELECT foto_id, url, vivienda_id FROM foto
WHERE foto_id = ?;

-- name: CrearFoto :one
INSERT INTO foto (url, vivienda_id)
VALUES (?, ?)
RETURNING foto_id, url, vivienda_id;

-- name: ActualizarFoto :one
UPDATE foto
SET url = ?, vivienda_id = ?
WHERE foto_id = ?
RETURNING foto_id, url, vivienda_id;

-- name: BorrarFoto :execrows
DELETE FROM foto WHERE foto_id = ?;


-- APLICAR VIVIENDA --

-- name: ListarAplicarViviendas :many
SELECT aplicar_vivienda_id, estudiante_id, vivienda_id, estado FROM aplicar_vivienda;

-- name: BuscarAplicarViviendaPorID :one
SELECT aplicar_vivienda_id, estudiante_id, vivienda_id, estado FROM aplicar_vivienda
WHERE aplicar_vivienda_id = ?;

-- name: CrearAplicarVivienda :one
INSERT INTO aplicar_vivienda (estudiante_id, vivienda_id, estado)
VALUES (?, ?, ?)
RETURNING aplicar_vivienda_id, estudiante_id, vivienda_id, estado;

-- name: ActualizarAplicarVivienda :one
UPDATE aplicar_vivienda
SET estudiante_id = ?, vivienda_id = ?, estado = ?
WHERE aplicar_vivienda_id = ?
RETURNING aplicar_vivienda_id, estudiante_id, vivienda_id, estado;

-- name: BorrarAplicarVivienda :execrows
DELETE FROM aplicar_vivienda WHERE aplicar_vivienda_id = ?;

-- SECTORES --

-- name: ListarSectores :many
SELECT sector_id, nombre FROM sector;

-- name: BuscarSectorPorID :one
SELECT sector_id, nombre FROM sector
WHERE sector_id = ?;

-- name: CrearSector :one
INSERT INTO sector (nombre)
VALUES (?)
RETURNING sector_id, nombre;

-- name: ActualizarSector :one
UPDATE sector
SET nombre = ?
WHERE sector_id = ?
RETURNING sector_id, nombre;

-- name: BorrarSector :execrows
DELETE FROM sector WHERE sector_id = ?;

