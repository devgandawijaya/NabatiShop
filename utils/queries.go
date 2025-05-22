package utils

const (
	GetAllShopActive = `
		SELECT 
				s.id,
				s.name,
				s.created_at,
				s.updated_at,
				COALESCE(
					JSON_AGG(
					json_build_object(
						'id', w.id,
						'shop_id', w.shop_id,
						'name', w.name,
						'is_active', w.is_active,
						'created_at', w.created_at,
						'updated_at', w.updated_at
					)
					) FILTER (WHERE w.id IS NOT NULL), '[]'
				) AS warehouses
				FROM shops s
				LEFT JOIN warehouses w ON w.shop_id = s.id
				GROUP BY s.id
	`

	InsertShop = `
		INSERT INTO shops (name, created_at, updated_at)
		VALUES ($1, NOW(), NOW())
		RETURNING id, name, created_at, updated_at
	`

	GetAllWarehouseWithStock = `
	SELECT 
		w.id,
		w.shop_id,
		w.name,
		w.is_active,
		w.created_at,
		w.updated_at,
		COALESCE(
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', ws.id,
					'warehouse_id', ws.warehouse_id,
					'product_id', ws.product_id,
					'available_qty', ws.available_qty,
					'reserved_qty', ws.reserved_qty
				)
			) FILTER (WHERE ws.id IS NOT NULL), '[]'
		) AS stocks
	FROM warehouses w
	LEFT JOIN warehouse_stocks ws ON ws.warehouse_id = w.id
	GROUP BY w.id
`

	InsertWarehouse = `
		INSERT INTO warehouses (shop_id, name, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, shop_id, name, is_active, created_at, updated_at
	`

	UpdateWarehouse = `
		UPDATE warehouses
		SET name = $1, is_active = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, shop_id, name, is_active, created_at, updated_at
	`

	DeleteWarehouse = `DELETE FROM warehouses WHERE id = $1`

	GetStockByWarehouseID = `SELECT id, warehouse_id, product_id, available_qty, reserved_qty FROM warehouse_stocks WHERE warehouse_id = $1;
	`

	GetStockByWarehouseProductID = `SELECT id, warehouse_id, product_id, available_qty, reserved_qty FROM warehouse_stocks WHERE warehouse_id = $1 AND product_id = $2;
	`

	InsertTransfer = `
		INSERT INTO warehouse_transfers (product_id, from_warehouse_id, to_warehouse_id, quantity)
		VALUES ($1, $2, $3, $4)
		RETURNING id, product_id, from_warehouse_id, to_warehouse_id, quantity, transferred_at`

	GetAllTransfers = `
		SELECT id, product_id, from_warehouse_id, to_warehouse_id, quantity, transferred_at
		FROM warehouse_transfers
		ORDER BY transferred_at DESC;`

	GetStock = `
				SELECT * FROM warehouse_stocks
				WHERE warehouse_id = $1 AND product_id = $2`

	UpdateStock = `
				UPDATE warehouse_stocks
				SET available_qty = $1
				WHERE id = $2`

	CheckWarehouseActive = `SELECT is_active FROM warehouses WHERE id = $1
			`
)
