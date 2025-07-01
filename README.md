# JalanKerja Backend

## Project Structures

```bash
.
└── src/
    ├── api/                                 # Layer presentasi (HTTP/API layer)
    │   ├── controllers/
    │   │   └── v1/                          # Versi API v1
    │   │       ├── dto/                     # DTO (Data Transfer Object)
    │   │       │   └── squad_dto.go         # Struct DTO untuk operasi Squad
    │   │       ├── handler/                 # Handler yang menangani logic
    │   │       │   └── squad_handler.go     # Handler fungsi CRUD untuk Squad
    │   │       ├── mapper                   # (Opsional) Mapping antara model, DTO
    │   │       └── squad_controller.go      # Menghubungkan route ke handler Squad
    │   ├── presenters/                      # Output formatter
    │   ├── middleware/                      # Middleware seperti Auth, Logging, dll.
    │   └── routes/                          # Routing layer
    │       └── v1/                          # Versi API v1
    │          ├── routes.go                 # Entry point untuk semua route
    │          └── squad_routes.go           # Daftar route terkait Squad
    ├── cmd/                                 # Entry point aplikasi
    │   └── main.go                          # Fungsi utama untuk menjalankan server
    ├── internal/
    │   └── database/
    │       ├── config/                      # Konfigurasi
    │       │   └── postgre.go               # Setup PostgreSQL (DSN, koneksi, migrasi)
    │       ├── models/                      # Model ORM (representasi tabel database)
    │       │   └── squads.go                # Struct model Squad
    │       └── seeder/                      # Seeder data awal ke database
    └── pkg/
        ├── repository/                      # Layer repository (akses data ke DB)
        │   ├── adapter/
        │   │   └── sql/                     # Implementasi repository berbasis SQL
        │   │       └── squad.go             # Repository interface Squad
        │   └── query/
        │       └── sql/
        │           └── squad.go             # Implement repo Squad
        └── services/                        # Logika bisnis
            └── v1/                          # Versi API v1
                └── squad_service.go         # Business logic squad
```

# Flow

![Flow Diagram](https://jam.dev/cdn-cgi/image/width=1000,quality=100,dpr=1/https://cdn-jam-screenshots.jam.dev/4bee731580457b0e55664da35511cc00/screenshot/da869424-f1e0-4435-a921-c790171b9c9d.png)

# Installations

1. Clone Repo

   ```bash
   git clone <repo>
   cd <repo>
   ```

2. Copy environtment

   ```bash
   cp src/.env.example src/.env
   ```

3. Run App

   - Development

   ```javascript
   cd src
   go mod tidy
   go run cmd/main.go
   ```

   - Production

   ```bash
   make build
   make run

   # for log
   make logs
   ```

App run on <http://localhost:5000>

## Using Generator

1. Masuk `src` directory
2. Run command

   ```javascript
   // Untuk Generate API Layer Files
   ./generator api:generate <entities>

   // Untuk Generate PKG Layer Files
   ./generator pkg:generate <entities>
   ```
