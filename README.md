# API_para_estudiantes_for-neos-
API REST desarrollada en Go para ayudar a estudiantes foráneos en Manta a encontrar servicios esenciales al llegar a la ciudad:
Se busca centralizar información sobre:
    -Viviendas disponibles
    -Lugares de Alimentacion
    -Rutas de Transporte
Actualmente muchos estudiantes buscan esta informacion mediantes grupos de WhatsApp, publicaciones en FaceBook y recomendaciones informales.

Nuestra API busca organizar esta información mediante recomendaciones basadas en:
    -Presupuesto
    -Ubicación
    -Horarios
    -Disponibilidad

Problema
Los estudiantes foraneos que llegan a Manta para estudiar un la ULEAM tienen difucultades para encontrar:
    -habitaciones accesibles
    -comida economica
    -transporte adecuado

Usuarios
-Estudiantes
Busca opciones de vivienda, comida y transporte
-Proveedor
Publica información sobre:
    -habitaciones
    -menús
    -rutas o servicios disponibles

Módulos
    -Modulo de Vivienda
    Gestión y busqueda de viviendas
    -Modulo de Transporte
    Gestion y busqueda de transporte.
    -Modulo de Alimentación
    Gestión y busqueda de opciones de alimentación.
    -Modulo de Usuario
    Registro e inicio de sesión para estudiantes y proveedores

Tech Stack
    -Go

Project Status
Project currently in Hito 1 (Discovery Phase)

Current progress:
    - Problem validation
    - Users interviews
    - Initial domain desing
    - API planning
Development of endpoints and business logic will begin in the next Hito.

Team Members
- Yuleisi
- Joyce
- Pablo

# Tecnologías utilizadas

- Go 1.26
- Chi Router
- GORM
- SQLite
- PostgreSQL
- JWT
- Docker
- Docker Compose
- GitHub Actions

---

# Arquitectura

La aplicación sigue una arquitectura por capas.

```
cmd/
internal/
    alimentacion/
        handlers/
        models/
        service/

    transporte/
        handlers/
        models/
        service/

    vivienda/
        handlers/
        models/
        service/

    usuario/
        handlers/
        models/
        service/

    config/
    httpserver/
    middleware/
    storage/
```

Cada módulo se encuentra dividido en:

- Models
- Service
- Handlers

Además existen componentes compartidos para:

- Configuración
- Middleware
- Storage
- HTTP Server

---

# Funcionalidades

## Autenticación

- Registro de usuarios
- Login
- JWT
- Middleware de autenticación

---

## Viviendas

- CRUD Viviendas
- CRUD Fotos
- CRUD Sectores
- CRUD Aplicar Vivienda

---

## Alimentación

- CRUD Alimentación
- CRUD Menú Diario
- CRUD Platos
- CRUD Reseñas

---

## Transporte

- CRUD Rutas
- CRUD Cooperativas
- CRUD Paradas

---

# Patrones implementados

## Factory

Centraliza la inicialización del almacenamiento.

Archivo:

```
internal/storage/factory.go
```

Permite cambiar entre SQLite y PostgreSQL sin modificar la lógica de negocio.

---

## Options Pattern

Utilizado para configurar el servicio de autenticación.

Permite inyectar:

- secreto JWT
- duración del token

sin utilizar variables globales.

---

## Parameter Object

Los handlers


