package query

const InsertBlockLayout = `
INSERT INTO "Storage".block_layout (warehouse_id, block_name, floor, x_sum, y_sum, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $6)
RETURNING id;
`

const GetAllBlockLayout = `
SELECT id, warehouse_id, block_name, floor, x_sum, y_sum,
	created_at, created_by, updated_at, updated_by
FROM "Storage".block_layout
ORDER BY id;
`

const GetBlockLayoutByID = `
SELECT id, warehouse_id, block_name, floor, x_sum, y_sum,
	created_at, created_by, updated_at, updated_by
FROM "Storage".block_layout
WHERE id = $1;
`

const UpdateBlockLayoutByID = `
UPDATE "Storage".block_layout
SET warehouse_id = $1,
	block_name = $2,
	floor = $3,
	x_sum = $4,
	y_sum = $5,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $6
WHERE id = $7;
`

const DeleteBlockLayoutByID = `
DELETE FROM "Storage".block_layout
WHERE id = $1;
`
