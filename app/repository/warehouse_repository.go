package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/repository/query"
	"wh-storage-service/config"
)

type warehouseRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewWarehouseRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) WarehouseRepositoryInterface {
	return &warehouseRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *warehouseRepository) InsertWarehouse(ctx context.Context, data entity.Warehouse) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var id int
	err := r.db.QueryRowContext(ctx, query.InsertWarehouse,
		data.Name,
		data.Code,
		data.Address,
		data.TtlArea,
		r.serviceName,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.InsertWarehouse (Scan): %w", err)
	}

	return id, nil
}

func (r *warehouseRepository) GetAllWarehouse(ctx context.Context) ([]entity.Warehouse, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllWarehouse)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllWarehouse (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.Warehouse
	for rows.Next() {
		var data entity.Warehouse
		if err := scanWarehouse(rows, &data); err != nil {
			return nil, fmt.Errorf("repo.GetAllWarehouse (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllWarehouse (Rows): %w", err)
	}

	return list, nil
}

func (r *warehouseRepository) GetWarehouseByID(ctx context.Context, id int) (entity.Warehouse, error) {
	var data entity.Warehouse

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetWarehouseByID, id).Scan(
		&data.ID,
		&data.Name,
		&data.Code,
		&data.Address,
		&data.TtlArea,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.Warehouse{}, fmt.Errorf("repo.GetWarehouseByID (Scan): %w", err)
	}

	return data, nil
}

func (r *warehouseRepository) UpdateWarehouseByID(ctx context.Context, data entity.Warehouse) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateWarehouseByID,
		data.Name,
		data.Code,
		data.Address,
		data.TtlArea,
		r.serviceName,
		data.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateWarehouseByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateWarehouseByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *warehouseRepository) DeleteWarehouseByID(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteWarehouseByID, id)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteWarehouseByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteWarehouseByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func scanWarehouse(rows *sql.Rows, data *entity.Warehouse) error {
	return rows.Scan(
		&data.ID,
		&data.Name,
		&data.Code,
		&data.Address,
		&data.TtlArea,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
}
