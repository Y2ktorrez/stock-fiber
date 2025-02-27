# 🚀 Stock Management

![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15%2B-brightgreen)
![Docker](https://img.shields.io/badge/Docker-24.0%2B-blueviolet)

Servicio para gestión de inventario desarrollado en **Go**, con persistencia en **PostgreSQL** y containerizado con **Docker**. Ideal para sistemas de **e-commerce** o **gestión logística**.

---

## 📦 Características Principales
✅ API RESTful con endpoints CRUD para gestión de stock
✅ Arquitectura containerizada con Docker Compose
✅ Base de datos PostgreSQL con volumen persistente
✅ Configuración mediante variables de entorno
✅ Health checks integrados para la base de datos
✅ Sistema escalable y listo para producción

---

## 🚀 Empezando Rápido

### 🔧 Prerrequisitos
- **Docker** 24.0+
- **Docker Compose** 2.20+

### ▶️ Ejecutar el Proyecto

1️⃣ Clonar el repositorio:
```bash
git clone https://github.com/tu-usuario/stock-service.git
cd stock-service
```

2️⃣ Construir y levantar los contenedores:
```bash
docker-compose up --build
```

3️⃣ Los servicios estarán disponibles en:
- **API:** [http://localhost:8000](http://localhost:8000)
- **PostgreSQL:** `postgres://y2k:root@localhost:5432/stock`

---

## 🛠 Comandos Esenciales de Docker

| Comando | Descripción |
|---------|------------|
| `docker-compose up -d` | Iniciar servicios en segundo plano |
| `docker-compose down` | Detener y eliminar contenedores |
| `docker-compose logs -f stock_go_app` | Ver logs de la aplicación |
| `docker-compose restart stock_go_app` | Reiniciar solo la aplicación |
| `docker volume inspect stock-service_pgdata` | Inspeccionar volumen de PostgreSQL |

---

## 🔧 Configuración del Entorno
El servicio se configura mediante variables de entorno (modificables en `docker-compose.yml`):

```env
PORT=8000
DB_HOST=stock_go_db
DB_USER=y2k
DB_PASS=root
DB_NAME=stock
DB_PORT=5432
```

---

## 📡 Endpoints de la API

### 🔍 Obtener todo el stock
```http
GET /api/v1/items
```

### ➕ Crear nuevo ítem
```http
POST /api/v1/items
Content-Type: application/json
```
```json
{
  "product_id": "PROD-001",
  "quantity": 50,
  "location": "warehouse-a"
}
```

### 🔄 Actualizar stock
```http
PUT /api/v1/items/:id
Content-Type: application/json
```
```json
{
  "quantity": 25
}
```

---

## 🛠 Desarrollo Local

### 📌 Iniciar solo la base de datos:
```bash
docker-compose up -d stock_go_db
```

### 🔄 Ejecutar la aplicación con live-reload (requiere Air):
```bash
air -c .air.toml
```

### ✅ Ejecutar tests:
```bash
go test -v ./...
```

---

## 📊 Health Check del Sistema

Verificar estado de los servicios:
```bash
curl http://localhost:8000/health
```

Respuesta esperada:
```json
{
  "status": "OK",
  "db_status": "Connected",
  "timestamp": "2024-02-20T12:34:56Z"
}
```
