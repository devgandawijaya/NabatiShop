CREATE TABLE shops (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE warehouses (
    id SERIAL PRIMARY KEY,
    shop_id INTEGER REFERENCES shops(id),
    name VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE warehouse_stocks (
    id SERIAL PRIMARY KEY,
    warehouse_id INTEGER REFERENCES warehouses(id),
    product_id INTEGER DEFAULT 0,
    available_qty INTEGER DEFAULT 0,
    reserved_qty INTEGER DEFAULT 0,
    UNIQUE (warehouse_id)
);


CREATE TABLE warehouse_transfers (
    id SERIAL PRIMARY KEY,
    product_id  INTEGER DEFAULT 0,
    from_warehouse_id INTEGER REFERENCES warehouses(id),
    to_warehouse_id INTEGER REFERENCES warehouses(id),
    quantity INTEGER NOT NULL,
    transferred_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO shops (id, name) VALUES
(1, 'Toko Elektronik Jaya'),
(2, 'Toko Fashion Trendy');


INSERT INTO warehouses (id, shop_id, name, is_active) VALUES
(1, 1, 'Gudang Utama Jakarta', TRUE),
(2, 1, 'Gudang Cadangan Bandung', TRUE),
(3, 2, 'Gudang Fashion Jakarta', TRUE),
(4, 2, 'Gudang Fashion Surabaya', FALSE);


INSERT INTO warehouse_stocks (warehouse_id, product_id, available_qty, reserved_qty) VALUES
(1, 1, 60, 0), 
(2, 1, 40, 0),
(3, 2, 30, 0), 
(4, 2, 20, 0); 


