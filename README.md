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

** Response :
```bash
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T05:51:13Z",
    "data": [
        {
            "id": 2,
            "warehouse_id": 2,
            "product_id": 1,
            "available_qty": 100,
            "reserved_qty": 0
        }
    ],
    "message": "Success"
}
