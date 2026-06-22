-- ALIMENTACION --

-- name: ListarAlimentaciones :many
SELECT * FROM alimentacion;

-- name: ObtenerAlimentacion :one
SELECT * FROM alimentacion
WHERE id = ?;

-- name: CrearAlimentacion :exec
INSERT INTO alimentacion (
    nombre_local, descripcion, ubicacion, direccion,
    horario_apertura, horario_cierre, telefono,
    tipo_comida, precio_promedio, provider_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ActualizarAlimentacion :exec
UPDATE alimentacion SET
    nombre_local = ?, descripcion = ?, ubicacion = ?, direccion = ?,
    horario_apertura = ?, horario_cierre = ?, telefono = ?,
    tipo_comida = ?, precio_promedio = ?, provider_id = ?
WHERE id = ?;

-- name: EliminarAlimentacion :exec
DELETE FROM alimentacion
WHERE id = ?;


-- MENU DIARIO --

-- name: ListarMenuDiarios :many
SELECT * FROM menu_diario;

-- name: ObtenerMenuDiario :one
SELECT * FROM menu_diario
WHERE id = ?;

-- name: ListarMenuDiariosPorAlimentacion :many
SELECT * FROM menu_diario
WHERE alimentacion_id = ?;

-- name: CrearMenuDiario :exec
INSERT INTO menu_diario (fecha, alimentacion_id)
VALUES (?, ?);

-- name: ActualizarMenuDiario :exec
UPDATE menu_diario SET
    fecha = ?, alimentacion_id = ?
WHERE id = ?;

-- name: EliminarMenuDiario :exec
DELETE FROM menu_diario
WHERE id = ?;


-- PLATOS --

-- name: ListarPlatos :many
SELECT * FROM plato;

-- name: ObtenerPlato :one
SELECT * FROM plato
WHERE id = ?;

-- name: ListarPlatosPorMenuDiario :many
SELECT * FROM plato
WHERE menu_diario_id = ?;

-- name: CrearPlato :exec
INSERT INTO plato (
    nombre_plato, descripcion, categoria, precio, menu_diario_id
) VALUES (?, ?, ?, ?, ?);

-- name: ActualizarPlato :exec
UPDATE plato SET
    nombre_plato = ?, descripcion = ?, categoria = ?,
    precio = ?, menu_diario_id = ?
WHERE id = ?;

-- name: EliminarPlato :exec
DELETE FROM plato
WHERE id = ?;


-- RESENAS --

-- name: ListarResenas :many
SELECT * FROM resena;

-- name: ObtenerResena :one
SELECT * FROM resena
WHERE id = ?;

-- name: ListarResenasPorAlimentacion :many
SELECT * FROM resena
WHERE alimentacion_id = ?;

-- name: CrearResena :exec
INSERT INTO resena (comentario, calificacion, alimentacion_id)
VALUES (?, ?, ?);

-- name: ActualizarResena :exec
UPDATE resena SET
    comentario = ?, calificacion = ?, alimentacion_id = ?
WHERE id = ?;

-- name: EliminarResena :exec
DELETE FROM resena
WHERE id = ?;


-- COOPERATIVAS --

-- name: ListarCooperativas :many
SELECT * FROM cooperativa;

-- name: ObtenerCooperativa :one
SELECT * FROM cooperativa
WHERE id = ?;

-- name: CrearCooperativa :exec
INSERT INTO cooperativa (nombre, telefono, descripcion)
VALUES (?, ?, ?);

-- name: ActualizarCooperativa :exec
UPDATE cooperativa SET
    nombre = ?, telefono = ?, descripcion = ?
WHERE id = ?;

-- name: EliminarCooperativa :exec
DELETE FROM cooperativa
WHERE id = ?;


-- RUTAS DE TRANSPORTE --

-- name: ListarRutas :many
SELECT * FROM ruta_transporte;

-- name: ObtenerRuta :one
SELECT * FROM ruta_transporte
WHERE id = ?;

-- name: CrearRuta :exec
INSERT INTO ruta_transporte (
    nombre_linea, frecuencia_aprox, precio, descripcion_ruta,
    cooperativa_id, sector_origen_id, sector_destino_id, provider_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: ActualizarRuta :exec
UPDATE ruta_transporte SET
    nombre_linea = ?, frecuencia_aprox = ?, precio = ?,
    descripcion_ruta = ?, cooperativa_id = ?,
    sector_origen_id = ?, sector_destino_id = ?, provider_id = ?
WHERE id = ?;

-- name: EliminarRuta :exec
DELETE FROM ruta_transporte
WHERE id = ?;

-- VIVIENDAS --

-- name: ListarViviendas :many
SELECT * FROM vivienda;

-- name: ObtenerVivienda :one
SELECT * FROM vivienda
WHERE vivienda_id = ?;

-- name: CrearVivienda :exec
INSERT INTO vivienda (
    titulo, tipo_vivienda, precio, garantia, precio_garantia,
    direccion, luz, agua, amueblado, internet, bano_privado,
    numero_habitaciones, mascotas, genero_preferido, reglas_convivencia,
    telefono, email, estado, comentario, sector_id, proveedor_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ActualizarVivienda :exec
UPDATE vivienda SET
    titulo = ?, tipo_vivienda = ?, precio = ?, garantia = ?,
    precio_garantia = ?, direccion = ?, luz = ?, agua = ?,
    amueblado = ?, internet = ?, bano_privado = ?,
    numero_habitaciones = ?, mascotas = ?, genero_preferido = ?,
    reglas_convivencia = ?, telefono = ?, email = ?,
    estado = ?, comentario = ?, sector_id = ?, proveedor_id = ?
WHERE vivienda_id = ?;

-- name: EliminarVivienda :exec
DELETE FROM vivienda
WHERE vivienda_id = ?;


-- FOTOS --

-- name: ListarFotos :many
SELECT * FROM foto;

-- name: ObtenerFoto :one
SELECT * FROM foto
WHERE foto_id = ?;

-- name: ListarFotosPorVivienda :many
SELECT * FROM foto
WHERE vivienda_id = ?;

-- name: CrearFoto :exec
INSERT INTO foto (url, vivienda_id)
VALUES (?, ?);

-- name: ActualizarFoto :exec
UPDATE foto SET
    url = ?, vivienda_id = ?
WHERE foto_id = ?;

-- name: EliminarFoto :exec
DELETE FROM foto
WHERE foto_id = ?;


-- APLICAR VIVIENDA --

-- name: ListarAplicarViviendas :many
SELECT * FROM aplicar_vivienda;

-- name: ObtenerAplicarVivienda :one
SELECT * FROM aplicar_vivienda
WHERE aplicar_vivienda_id = ?;

-- name: CrearAplicarVivienda :exec
INSERT INTO aplicar_vivienda (estudiante_id, vivienda_id, estado)
VALUES (?, ?, ?);

-- name: ActualizarAplicarVivienda :exec
UPDATE aplicar_vivienda SET
    estudiante_id = ?, vivienda_id = ?, estado = ?
WHERE aplicar_vivienda_id = ?;

-- name: EliminarAplicarVivienda :exec
DELETE FROM aplicar_vivienda
WHERE aplicar_vivienda_id = ?;