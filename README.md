# 🏪 Go Shop Warehouse API

**Go Shop Warehouse** adalah Microservice RESTful API berbasis Golang yang menangani manajemen stok gudang, produk, toko, dan proses transfer stok antar gudang. Proyek ini cocok untuk sistem inventaris toko yang membutuhkan pengelolaan stok lintas gudang secara efisien dan otomatis.

---

## 🚀 Fitur Utama

- ✅ CRUD Toko dan Gudang
- 📦 Manajemen stok produk per gudang
- 🔄 Transfer stok antar gudang
- ❗ Validasi stok saat transfer
- 📅 Riwayat transfer lengkap

---

## 🐳 Jalankan dengan Docker

### Build dan Jalankan Manual
```bash
# Build Docker image
docker build -t go-shop-warehouse .

# Jalankan container
docker run -d -p 8080:8080 --name go-shop-warehouse go-shop-warehouse
