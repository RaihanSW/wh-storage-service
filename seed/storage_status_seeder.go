package seed

import (
	"context"
	"database/sql"
)

func init() {
	Register(Seeder{
		Name:  "storage_status",
		DBKey: "warehouse_db",
		Run: func(ctx context.Context, db *sql.DB) error {
			// TODO: sesuaikan nilai lookup
			statuses := []struct {
				id   int
				name string
			}{
				{1, "Empty"},
				{2, "Occupied"},
				{3, "Maintenance"},
			}

			for _, s := range statuses {
				_, err := db.ExecContext(ctx, `
					INSERT INTO "Storage".storage_status (id, name, created_by, updated_by)
					VALUES ($1, $2, $3, $3)
					ON CONFLICT (id) DO NOTHING
				`, s.id, s.name, "seeder")
				if err != nil {
					return err
				}
			}

			// id eksplisit di atas tidak menggerakkan identity sequence —
			// selaraskan agar insert berikutnya tidak bentrok PK
			_, err := db.ExecContext(ctx, `
				SELECT setval(
					pg_get_serial_sequence('"Storage".storage_status', 'id'),
					(SELECT MAX(id) FROM "Storage".storage_status)
				)
			`)
			return err
		},
	})
}
