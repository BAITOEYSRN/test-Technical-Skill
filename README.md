# User Management API

API สำหรับจัดการข้อมูลผู้ใช้ (User Management) ที่พัฒนาด้วย Go และ Gin Framework

## 📋 สิ่งที่ต้องมีก่อนเริ่มต้น

- **Go** เวอร์ชัน 1.25.0 ขึ้นไป ([ดาวน์โหลดที่นี่](https://go.dev/dl/))
- **PostgreSQL** เวอร์ชัน 12 ขึ้นไป ([ดาวน์โหลดที่นี่](https://www.postgresql.org/download/))
- **Git** สำหรับ clone repository

## 🚀 วิธีติดตั้งและใช้งาน

### 1. Clone Repository

```bash
git clone <repository-url>
cd test-Technical-Skill
```

### 2. ติดตั้ง Dependencies

```bash
go mod download
```

### 3. สร้างไฟล์ `.env`

สร้างไฟล์ `.env` ที่ root ของโปรเจกต์ และใส่ค่าตามนี้:

```env
# Server Configuration
PORT=4000

# Database Configuration
DB_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_SSLMODE=disable
DB_SCHEMA_NAME=develop

# Migration Configuration
PATH_MIGRATIONS=internal/infrastructure/migrations
MIGRATION_DB_AUTO=true
```

**หมายเหตุ:** แก้ไขค่าต่างๆ ให้ตรงกับ PostgreSQL ของคุณ

### 4. สร้าง Database และ Schema

เข้าไปใน PostgreSQL และสร้าง database กับ schema:

```sql
-- สร้าง database
CREATE DATABASE your_database_name;

-- สร้าง schema
CREATE SCHEMA develop;
```

### 5. รัน Migration (ถ้า MIGRATION_DB_AUTO=true จะรันอัตโนมัติ)

ถ้าตั้งค่า `MIGRATION_DB_AUTO=true` ในไฟล์ `.env` ระบบจะรัน migration อัตโนมัติเมื่อเริ่มแอปพลิเคชัน

### 6. รันแอปพลิเคชัน

```bash
go run cmd/main.go
```

หรือถ้าต้องการ build ก่อน:

```bash
go build -o app cmd/main.go
./app
```

แอปพลิเคชันจะรันที่ `http://localhost:4000` (หรือ port ที่กำหนดใน `.env`)

## 📡 API Endpoints

### Health Check
```
GET /health
```
ตรวจสอบสถานะของ API

### User Endpoints

#### 1. สร้างผู้ใช้ใหม่
```
POST /api/v1/users
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "date_of_birth": "1990-01-15",
  "age": 34,
  "address": "123 Main Street, Bangkok"
}
```

#### 2. ดึงรายชื่อผู้ใช้ทั้งหมด
```
GET /api/v1/users/list
```

#### 3. ดึงข้อมูลผู้ใช้ตาม ID
```
GET /api/v1/users/:id
```


