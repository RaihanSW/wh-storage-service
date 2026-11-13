package query

const InsertStorageStatus = `
INSERT INTO "Storage".storage_status (name, created_by, updated_by)
VALUES ($1, $2, $2)
RETURNING id;
`

const GetAllStorageStatus = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Storage".storage_status
ORDER BY id;
`

const GetStorageStatusByID = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Storage".storage_status
WHERE id = $1;
`

const UpdateStorageStatusByID = `
UPDATE "Storage".storage_status
SET name = $1,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $2
WHERE id = $3;
`

const DeleteStorageStatusByID = `
DELETE FROM "Storage".storage_status
WHERE id = $1;
`
