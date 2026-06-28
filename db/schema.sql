-- alimentacion --

CREATE TABLE alimentacion (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre_local TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    ubicacion TEXT NOT NULL,
    direccion TEXT NOT NULL,
    horario_apertura TEXT NOT NULL,
    horario_cierre TEXT NOT NULL,
    telefono TEXT NOT NULL,
    tipo_comida TEXT NOT NULL,
    precio_promedio REAL NOT NULL,
    provider_id INTEGER NOT NULL
);

CREATE TABLE menu_diario (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fecha TEXT NOT NULL,
    alimentacion_id INTEGER NOT NULL,
    FOREIGN KEY (alimentacion_id) REFERENCES alimentacion(id)
);

CREATE TABLE plato (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre_plato TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    categoria TEXT NOT NULL,
    precio REAL NOT NULL,
    menu_diario_id INTEGER NOT NULL,
    FOREIGN KEY (menu_diario_id) REFERENCES menu_diario(id)
);

CREATE TABLE resena (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    comentario TEXT NOT NULL,
    calificacion INTEGER NOT NULL,
    alimentacion_id INTEGER NOT NULL,
    FOREIGN KEY (alimentacion_id) REFERENCES alimentacion(id)
);

-- transporte --

CREATE TABLE cooperativa (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    telefono TEXT NOT NULL,
    descripcion TEXT NOT NULL
);

CREATE TABLE ruta_transporte (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre_linea TEXT NOT NULL,
    frecuencia_aprox TEXT NOT NULL,
    precio REAL NOT NULL,
    descripcion_ruta TEXT NOT NULL,
    cooperativa_id INTEGER NOT NULL,
    sector_origen_id INTEGER NOT NULL,
    sector_destino_id INTEGER NOT NULL,
    parada_bus_id INTEGER NOT NULL,
    FOREIGN KEY (cooperativa_id) REFERENCES cooperativa(id),
    FOREIGN KEY (sector_origen_id) REFERENCES sector(id),
    FOREIGN KEY (sector_destino_id) REFERENCES sector(id)
);

CREATE TABLE paradas_bus (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre_parada TEXT NOT NULL,
    direccion TEXT NOT NULL,
    descripcion TEXT NOT NULL
);

-- vivienda --

CREATE TABLE vivienda (
    vivienda_id INTEGER PRIMARY KEY AUTOINCREMENT,
    titulo TEXT NOT NULL,
    tipo_vivienda TEXT NOT NULL,
    precio REAL NOT NULL,
    garantia BOOLEAN NOT NULL,
    precio_garantia REAL NOT NULL,
    direccion TEXT NOT NULL,
    luz BOOLEAN NOT NULL,
    agua BOOLEAN NOT NULL,
    amueblado BOOLEAN NOT NULL,
    internet BOOLEAN NOT NULL,
    bano_privado BOOLEAN NOT NULL,
    numero_habitaciones INTEGER NOT NULL,
    mascotas BOOLEAN NOT NULL,
    genero_preferido TEXT NOT NULL,
    reglas_convivencia TEXT NOT NULL,
    telefono TEXT NOT NULL,
    email TEXT NOT NULL,
    estado TEXT NOT NULL,
    comentario TEXT NOT NULL,
    sector_id INTEGER NOT NULL,
    proveedor_id INTEGER NOT NULL,
    FOREIGN KEY (sector_id) REFERENCES sector(id)
);

CREATE TABLE foto (
    foto_id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    vivienda_id INTEGER NOT NULL,
    FOREIGN KEY (vivienda_id) REFERENCES vivienda(vivienda_id)
);

CREATE TABLE aplicar_vivienda (
    aplicar_vivienda_id INTEGER PRIMARY KEY AUTOINCREMENT,
    estudiante_id INTEGER NOT NULL,
    vivienda_id INTEGER NOT NULL,
    estado TEXT NOT NULL,
    FOREIGN KEY (vivienda_id) REFERENCES vivienda(vivienda_id)
);

CREATE TABLE sector (
    sector_id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL
);
