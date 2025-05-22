# 🏪 Go Shop Warehouse API

**Go Shop Warehouse** adalah RESTful API berbasis Golang yang dirancang untuk mengelola **stok gudang**, **produk**, **toko**, dan **transfer stok antar gudang**. Cocok digunakan untuk sistem inventaris multigudang dan skala toko retail.

---

## 🚀 Fitur Utama

- ✅ CRUD Toko dan Gudang
- 📦 Manajemen stok produk per gudang
- 🔄 Transfer stok antar gudang
- ❗ Validasi stok sebelum transfer
- 📅 Riwayat transfer stok antar gudang

---

## 📦 Endpoint API & Contoh `curl`

### 1. 🔍 Ambil Stok Berdasarkan Gudang
```bash
curl --location --request GET 'http://localhost:8080/warehouses/2/stocks'
