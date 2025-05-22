# 🏪 Go Shop Warehouse API

**Go Shop Warehouse** adalah RESTful API berbasis Golang yang menangani manajemen stok gudang, produk, toko, dan proses transfer stok antar gudang. Cocok untuk sistem inventaris retail dengan banyak cabang dan gudang, mendukung pengelolaan stok yang efisien dan transparan.

---

## 🚀 Fitur Utama

- ✅ CRUD Toko (Shop)
- 🏬 Manajemen Gudang per Toko
- 📦 Manajemen Stok Produk di Gudang
- 🔄 Transfer Stok Antar Gudang
- 🔍 Pencarian stok berdasarkan gudang dan produk
- 🐳 Dukungan penggunaan dengan Docker

---

## 📦 API Endpoints

### 1. 📥 Ambil Stok Berdasarkan Warehouse
**GET** `/warehouses/{warehouse_id}/stocks`
```bash
curl --location --request GET 'http://localhost:8080/warehouses/2/stocks'
