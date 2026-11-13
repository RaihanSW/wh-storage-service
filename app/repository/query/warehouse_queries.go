package query

const InsertWarehouse = `
INSERT INTO "Storage".warehouse (name, code, address, ttl_area, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $5)
RETURNING id;
`

const GetAllWarehouse = `
SELECT id, name, code, address, ttl_area,
	created_at, created_by, updated_at, updated_by
FROM "Storage".warehouse
ORDER BY id;
`

const GetWarehouseByID = `
SELECT id, name, code, address, ttl_area,
	created_at, created_by, updated_at, updated_by
FROM "Storage".warehouse
WHERE id = $1;
`

const UpdateWarehouseByID = `
UPDATE "Storage".warehouse
SET name = $1,
	code = $2,
	address = $3,
	ttl_area = $4,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $5
WHERE id = $6;
`

const DeleteWarehouseByID = `
DELETE FROM "Storage".warehouse
WHERE id = $1;
`
