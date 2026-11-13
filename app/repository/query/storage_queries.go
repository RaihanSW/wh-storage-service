package query

const InsertStorage = `
INSERT INTO "Storage".storage
	(name, goods_uid, warehouse_id, block_id, status, x_row, y_row, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $8)
RETURNING uid;
`

const GetAllStorage = `
SELECT uid, name, goods_uid, warehouse_id, block_id, status, x_row, y_row,
	created_at, created_by, updated_at, updated_by
FROM "Storage".storage
ORDER BY created_at;
`

const GetStorageByUID = `
SELECT uid, name, goods_uid, warehouse_id, block_id, status, x_row, y_row,
	created_at, created_by, updated_at, updated_by
FROM "Storage".storage
WHERE uid = $1;
`

const UpdateStorageByUID = `
UPDATE "Storage".storage
SET name = $1,
	goods_uid = $2,
	warehouse_id = $3,
	block_id = $4,
	status = $5,
	x_row = $6,
	y_row = $7,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $8
WHERE uid = $9;
`

const DeleteStorageByUID = `
DELETE FROM "Storage".storage
WHERE uid = $1;
`
