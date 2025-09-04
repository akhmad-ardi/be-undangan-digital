# 🎉 Platform Undangan Digital Berbasis Web - Backend (Go)

Backend platform **undangan digital berbasis web** menggunakan **Golang**.  
Menyediakan API untuk membuat, mengelola, dan membagikan undangan digital ke berbagai platform (misalnya WhatsApp).

---

## 🚀 Fitur Utama
- **Manajemen Undangan**
  - CRUD undangan digital
- **Share Undangan**
  - Generate link undangan
  - Bagikan ke WhatsApp

---

## 🛠️ Teknologi
- **Bahasa**: Go (Golang)
- **Framework**: Gofiber
- **Database**: supabase(PostgreSQL)
- **Konfigurasi**: `.env` dengan [godotenv](https://github.com/joho/godotenv)

---

## 📂 Struktur Proyek

BE-UNDANGAN-DIGITAL/ <br>
├── cmd/ # Entry point (CLI/server command) <br>
├── config/ # Konfigurasi (env, DB, dll.) <br>
├── controllers/ # Handler request/response <br>
├── docs/ # Dokumentasi API (Swagger/OpenAPI) <br>
├── lib/ # Library/helper umum <br>
├── middleware/ # Middleware (auth, logging, dll.) <br>
├── models/ # Definisi model/data <br>
├── public/ # File publik (jika ada) <br>
├── requests/ # Validasi request DTO <br>
├── rest_client/ # Client untuk API eksternal (mis. WhatsApp API) <br>
├── routes/ # Routing endpoint <br>
├── services/ # Logika bisnis utama <br>
├── tests/ # Unit & integration tests <br>
├── tmp/ # File sementara<br>
├── validations/ # Validator custom <br>
├── .air.toml # Konfigurasi live-reload (air) <br>
├── .env # Environment variabel (jangan commit) <br>
├── .env.example # Contoh environment <br>
├── go.mod # Module Go <br>
├── go.sum # Dependency lock <br>
├── main.go # Entry point utama aplikasi <br>
└── README.md # Dokumentasi proyek <br>

---

## ⚙️ Instalasi & Menjalankan
1. **Clone repositori**
   ```bash
   git clone https://github.com/akhmad-ardi/BE-UNDANGAN-DIGITAL.git
   cd BE-UNDANGAN-DIGITAL

2. Install dependencies
  ```bash
  go mod tidy
  ```

3. Set environment
  - Copy file ```.env.example``` ke ```.env```
  - Atur variabel sesuai kebutuhan:
  ``` .env
  DB_HOST=
  DB_USER=
  DB_PASSWORD=
  DB_NAME=
  DB_PORT=
  DB_POOL_MODE=

  JWT_SECRET=

  FRONT_END=
  ```

4. Jalankan aplikasi
  ``` bash
  go run main.go
  ```
  atau jika menggunakan air(hot reload):
  ``` bash
  air
  ```