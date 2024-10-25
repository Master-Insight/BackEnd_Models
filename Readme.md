# Backend Insight

## Estructura
El backend esta basado en Modelo, Vista, Controlador (MVC) en formato Modular segun su funcionalidad.

```
/app
│
├── /cmd                # Puntos de entrada (diferentes servicios o microservicios)
│   └── server.go
│
├── /modules            # Lógica interna de la aplicación (no exportada)
│   ├── /auth           # Módulo de autenticación
│   │   ├── api         # Capa Vista por módulo
│   │   │   ├── routes.go
│   │   │   ├── validators.go
│   │   │   ├── controller.go
│   │   │   └── dto.go
│   │   ├── data        # Capa Modelo por módulo
│   │   │   ├── dao.go
│   │   │   └── model.go
│   │   ├── logic       # Capa Controllador por módulo
│   │   │   ├── service.go
│   │   │   └── xxxxxxx.go
│   │   └── i18n        # Internacionalización por módulo
│   │       ├── en.json
│   │       └── es.json
│   │
│   ├── /user           # Módulo de usuarios
│   ├── /security       # Módulo de seguridad (autorización, roles)
│   ├── /product        # Módulo de productos
│   └── /order          # Otros módulos
│
├── /pkg                # Código compartido (reutilizable por otros proyectos)
│   ├── /middleware     # Middlewares compartidos
│   │   ├── auth_middleware.go
│   │   ├── cors_middleware.go
│   │   ├── logger_middleware.go
│   │   └── rate_limiter.go
│   ├── /validators     # Validadores generales
│   ├── /services       # Servicios compartidos (antes "libraries" o "utils")
│   │   ├── jwt         # Servicio de autenticación con JWT
│   │   ├── db          # Conexiones a bases de datos (MongoDB, MySQL, etc.)
│   │   ├── env         # Manejo de variables de entorno
│   │   ├── logs        # Servicio de logging
│   │   ├── email       # Servicio de envío de correos electrónicos
│   │   └── cache       # Servicio de cache (Redis, Memcached, etc.)
│   └── /custom         # Estructuras reutilizables
│       ├── helper_class.go
│       ├── factory_class.go
│       └── utilities.go
│
├── /config             # Configuración global de la aplicación
│   └── config.go
│
├── /migrations         # Migraciones de base de datos
│
├── /translations       # Sistema de internacionalización general (i18n)
│   ├── /en             # Archivos de traducción en inglés
│   └── /es             # Archivos de traducción en español
│
└── /scripts            # Scripts útiles (despliegue, testing, etc.)
```

## Explicación de cada parte:

1. **`/cmd`**:
   - Contiene los puntos de entrada de la aplicación, normalmente archivos como `main.go` o `server.go` que inician los microservicios o servicios principales.
   
2. **`/app`**:
   - Lógica interna de la aplicación dividida por módulos, lo cual permite un fácil escalamiento y separación de responsabilidades.
   - Cada módulo (por ejemplo, `auth`, `user`, `security`) tiene sus propias subcapas:
     - **`api`**: Capa de vista, incluye rutas, validadores, controladores, y objetos DTO (Data Transfer Objects).
     - **`data`**: Capa de modelo, responsable del acceso a la base de datos (DAO) y definición de modelos de datos.
     - **`logic`**: Capa de controlador, contiene la lógica de negocio principal, como los servicios.
     - **`i18n`**: Internacionalización por módulo, con archivos JSON para diferentes idiomas (en este caso, inglés y español).

3. **`/pkg`**:
   - Código compartido y reutilizable por otros módulos o proyectos.
   - **`middleware`**: Middlewares comunes y compartidos como autenticación, CORS, logging, y limitación de tasas de solicitudes.
   - **`validators`**: Validadores compartidos para diversas verificaciones de datos.
   - **`services`**: Servicios compartidos como autenticación JWT, manejo de bases de datos, variables de entorno, logging, envío de correos electrónicos y caching.
   - **`custom`**: Clases y estructuras reutilizables para tareas específicas, como helpers y factories.

4. **`/config`**:
   - Almacena la configuración global de la aplicación, como conexiones a base de datos, servicios externos, entre otros.

5. **`/migrations`**:
   - Almacena los scripts de migración de bases de datos, para mantener versiones y actualizaciones de la estructura de la base de datos.

6. **`/translations`**:
   - Sistema de internacionalización (i18n) global, con archivos de traducción que permiten que la aplicación soporte múltiples idiomas. En este caso, incluye inglés (`/en`) y español (`/es`).

7. **`/scripts`**:
   - Scripts útiles para despliegue, testing, migraciones, u otros procesos que faciliten el mantenimiento o la automatización de tareas dentro del entorno de desarrollo.

## Componentes

* Modulos:
  * auth: manejo de usuario, referente a sesiones y contraseñas
