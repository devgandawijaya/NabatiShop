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

## 🐳 Jalankan dengan Docker


### Build Docker image

```bash
  docker build -t go-shop-warehouse .
```

### Contoh DockerFile
```bash
  # Gunakan image Go resmi sebagai base
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy semua file ke dalam image
COPY . .

# Download dependency
RUN go mod download

# Build binary
RUN go build -o main .

# Expose port (sesuaikan dengan port yang digunakan)
EXPOSE 8080

# Command default saat container dijalankan
CMD ["./main"]

```

### Contoh Docker-Compose.yml
```bash
version: "3.8"

services:
  warehouse-api:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - db
    networks:
      - go-network

  db:
    image: postgres:14
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: pass
      POSTGRES_DB: go_warehouse
    ports:
      - "5432:5432"
    networks:
      - go-network

networks:
  go-network:

```



## API Reference

#### API Ambil Stock dan product berdasarkan warehouse

```http
  curl --location --request GET 'http://localhost:8080/warehouses/2/stocks'
```

```bash
Response :
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
```

####  Update Stock berdasarkan ID warehouse, ID Stock dan ID Product

```http
  curl --location --request PUT 'http://localhost:8080/warehouses/2/stocks/2/product/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "available_qty": 100
  }'
```

```bash
  Response :
  {
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T06:27:15Z",
    "data": {
        "id": 2,
        "warehouse_id": 2,
        "product_id": 1,
        "available_qty": 100,
        "reserved_qty": 0
    },
    "message": "Stock updated successfully"
}
```



####  Stock baru berdasarkan product baru

```http
  curl --location --request POST 'http://localhost:8080/warehouses/2/stocks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product_id": 2,
    "available_qty": 60,
    "reserved_qty": 0
  }'
```

```bash
  Response :
  {
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T06:52:27Z",
    "data": {
        "id": 11,
        "warehouse_id": 2,
        "product_id": 2,
        "available_qty": 60,
        "reserved_qty": 0
    },
    "message": "Stock created successfully"
}
```



####  Transfer Stock dari dan menuju gudang

```http
  curl --location --request POST 'http://localhost:8080/transfers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product_id": 1,
    "from_warehouse_id": 2,
    "to_warehouse_id": 3,
    "quantity": 1000
}'
```

```bash
  Response :
  {
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T08:38:14Z",
    "data": null,
    "message": "stok gudang asal tidak mencukupi"
}
```


####   Ambil seluruh stock di semua gudang

```http
  curl --location --request GET 'http://localhost:8080/transfers'

```

```bash
  Response :
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T08:37:57Z",
    "data": [
        {
            "id": 6,
            "product_id": 1,
            "from_warehouse_id": 2,
            "to_warehouse_id": 3,
            "quantity": 30,
            "transferred_at": "2025-05-22T08:15:14.601997Z"
        },
        {
            "id": 5,
            "product_id": 1,
            "from_warehouse_id": 2,
            "to_warehouse_id": 3,
            "quantity": 30,
            "transferred_at": "2025-05-22T08:11:36.134325Z"
        }
    ],
    "message": "Success"
}
```




####  ambil  stock berdasarkan warehouse dan product

```http
  curl --location --request GET 'http://localhost:8080/warehouses/1/stocks/2'

```

```bash
  Response :
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T09:07:06Z",
    "data": [
        {
            "id": 7,
            "warehouse_id": 1,
            "product_id": 2,
            "available_qty": 60,
            "reserved_qty": 0
        }
    ],
    "message": "Success"
}
```



####   ambil semua list shop

```http
  curl --location --request GET 'http://localhost:8080/shops'

```

```bash
 Response :
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T09:25:57Z",
    "data": [
        {
            "id": 2,
            "name": "Toko Fashion Trendy",
            "created_at": "2025-05-21T08:09:53.50108Z",
            "updated_at": "2025-05-21T08:09:53.50108Z",
            "warehouses": [
                {
                    "id": 3,
                    "shop_id": 2,
                    "name": "Gudang Fashion Jakarta",
                    "is_active": true,
                    "created_at": "2025-05-21T08:09:53.501436",
                    "updated_at": "2025-05-21T08:09:53.501436"
                },
                {
                    "id": 4,
                    "shop_id": 2,
                    "name": "Gudang Fashion Surabaya",
                    "is_active": false,
                    "created_at": "2025-05-21T08:09:53.501436",
                    "updated_at": "2025-05-21T08:09:53.501436"
                }
            ]
        },
        {
            "id": 7,
            "name": "Toko Berkah 5",
            "created_at": "2025-05-21T14:24:52.538991Z",
            "updated_at": "2025-05-21T14:25:22.777909Z",
            "warehouses": []
        },
        {
            "id": 1,
            "name": "Toko Elektronik Jaya",
            "created_at": "2025-05-21T08:09:53.50108Z",
            "updated_at": "2025-05-21T08:09:53.50108Z",
            "warehouses": [
                {
                    "id": 1,
                    "shop_id": 1,
                    "name": "Gudang Utama Jakarta",
                    "is_active": true,
                    "created_at": "2025-05-21T08:09:53.501436",
                    "updated_at": "2025-05-21T08:09:53.501436"
                },
                {
                    "id": 2,
                    "shop_id": 1,
                    "name": "Gudang Cadangan Bandung",
                    "is_active": true,
                    "created_at": "2025-05-21T08:09:53.501436",
                    "updated_at": "2025-05-21T08:09:53.501436"
                },
                {
                    "id": 5,
                    "shop_id": 1,
                    "name": "uptdae berhasil",
                    "is_active": true,
                    "created_at": "2025-05-22T02:34:01.283911",
                    "updated_at": "2025-05-22T02:39:55.822482"
                }
            ]
        },
        {
            "id": 3,
            "name": "test manual",
            "created_at": "2025-05-21T10:06:43.249718Z",
            "updated_at": "2025-05-21T10:06:43.249718Z",
            "warehouses": []
        }
    ],
    "message": "Success"
}

```



####  Tambah Shop Baru 

```http
 curl --location --request POST 'http://localhost:8080/shops' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"toko berkah jaya selalu"
}'
```

```bash

Response :
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T09:27:47Z",
    "data": {
        "id": 8,
        "name": "toko berkah jaya selalu",
        "created_at": "2025-05-22T09:27:47.890586Z",
        "updated_at": "2025-05-22T09:27:47.890586Z"
    },
    "message": "Shop created successfully"
}

```



####  Update Shop 

```http
 curl --location --request PUT 'http://localhost:8080/shops/8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"toko berkah jaya selalu 1"
}'
```

```bash

Response :
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T09:28:58Z",
    "data": {
        "id": 8,
        "name": "toko berkah jaya selalu 1",
        "created_at": "2025-05-22T09:27:47.890586Z",
        "updated_at": "2025-05-22T09:28:58.302015Z"
    },
    "message": "Shop updated successfully"
}

```



####  Delete Shops

```http
 curl --location --request DELETE 'http://localhost:8080/shops/8'

```

```bash

Response :
{
    "app": "go-shop-warehouse",
    "version": "1.0.0",
    "date": "2025-05-22T09:29:46Z",
    "data": null,
    "message": "Shop deleted successfully"
}

```
