# Go PostgreSQL User Management App

Ushbu loyiha Go tilida yozilgan bo'lib, foydalanuvchilarni ro'yxatga olish va ularning ro'yxatini ko'rish imkonini beruvchi oddiy RESTful API hisoblanadi. Ma'lumotlar bazasi sifatida PostgreSQL ishlatilgan.

## 🛠 Texnologiyalar
- **Go** (Golang)
- **PostgreSQL** (Ma'lumotlar bazasi)
- **Docker & Docker Compose** (Konteynerizatsiya)
- **Gorilla Mux** (Router)
- **Bcrypt** (Parollarni xavfsiz saqlash uchun)

### 3. Ma'lumotlar bazasini ishga tushirish (Docker orqali)
PostgreSQL bazasini Docker-da ishga tushirish uchun:

Bu buyruq bazani ishga tushiradi va `docker/postgres/init.sql` ichidagi jadvalni avtomatik yaratadi.

### 4. Go serverni ishga tushirish

Server ishga tushgach, `http://localhost:8080` manzilida xizmat ko'rsatadi.

---

## 📡 API Murojaatlari (Endpoints)

### 1. Server holatini tekshirish
*   **URL:** `/health`
*   **Metod:** `GET`
*   **Natija:** Server ishlayotgan bo'lsa `OK Health ishlati...` matni qaytadi.

### 2. Yangi foydalanuvchi qo'shish
*   **URL:** `/api/users`
*   **Metod:** `POST`
*   **Jo'natiladigan ma'lumot (JSON):**
    ```json
    {
        "username": "ali_valiyev",
        "email": "ali@example.com",
        "password": "maxfiy_parol"
    }
    ```
*   **Muvaffaqiyatli holat (201 Created):** Foydalanuvchi ma'lumotlari qaytadi (parol qaytarilmaydi).
*   **Xatolik holati (400 Bad Request):** Ma'lumotlar to'liq bo'lmasa yoki email band bo'lsa xato qaytadi.

### 3. Foydalanuvchilar ro'yxatini olish
*   **URL:** `/api/users`
*   **Metod:** `GET`
*   **Natija:** Bazadagi barcha foydalanuvchilarning ro'yxati (oxirgi qo'shilganlar birinchi bo'lib chiqadi).

---

## 📂 Loyiha tuzilishi (Folder Structure)

- `cmd/app/main.go` — Dasturga kirish nuqtasi va marshrutlar (routes).
- `internal/user/` — Foydalanuvchilar bilan bog'liq barcha biznes mantiq (Handler, Service, Repository).
- `internal/db/` — Ma'lumotlar bazasiga ulanish sozlamalari.
- `docker/` — PostgreSQL uchun boshlang'ich SQL skriptlari.
- `.env` — Maxfiy sozlamalar va ulanish ma'lumotlari.

## 🔒 Xavfsizlik
- Foydalanuvchi parollari bazada ochiq holatda saqlanmaydi.
- `bcrypt` kutubxonasi yordamida parollar xeshlangan (hash) holda saqlanadi.
- API-dan foydalanuvchi ma'lumotlarini qaytarganda parol xeshi ham JSON-dan olib tashlangan (`json:"-"`).
