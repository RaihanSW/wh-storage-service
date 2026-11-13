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

type storageRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewStorageRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) StorageRepositoryInterface {
	return &storageRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *storageRepository) InsertStorage(ctx context.Context, data entity.Storage) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var uid string
	err := r.db.QueryRowContext(ctx, query.InsertStorage,
		data.Name,
		data.GoodsUID,
		data.WarehouseID,
		data.BlockID,
		data.Status,
		data.XRow,
		data.YRow,
		r.serviceName,
	).Scan(&uid)
	if err != nil {
		return "", fmt.Errorf("repo.InsertStorage (Scan): %w", err)
	}

	return uid, nil
}

func (r *storageRepository) GetAllStorage(ctx context.Context) ([]entity.Storage, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllStorage)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllStorage (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.Storage
	for rows.Next() {
		var data entity.Storage
		if err := scanStorage(rows, &data); err != nil {
			return nil, fmt.Errorf("repo.GetAllStorage (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllStorage (Rows): %w", err)
	}

	return list, nil
}

func (r *storageRepository) GetStorageByUID(ctx context.Context, uid string) (entity.Storage, error) {
	var data entity.Storage

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetStorageByUID, uid).Scan(
		&data.UID,
		&data.Name,
		&data.GoodsUID,
		&data.WarehouseID,
		&data.BlockID,
		&data.Status,
		&data.XRow,
		&data.YRow,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.Storage{}, fmt.Errorf("repo.GetStorageByUID (Scan): %w", err)
	}

	return data, nil
}

func (r *storageRepository) UpdateStorageByUID(ctx context.Context, data entity.Storage) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateStorageByUID,
		data.Name,
		data.GoodsUID,
		data.WarehouseID,
		data.BlockID,
		data.Status,
		data.XRow,
		data.YRow,
		r.serviceName,
		data.UID,
	)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateStorageByUID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateStorageByUID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *storageRepository) DeleteStorageByUID(ctx context.Context, uid string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteStorageByUID, uid)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteStorageByUID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteStorageByUID (RowsAffected): %w", err)
	}

	return ra, nil
}

func scanStorage(rows *sql.Rows, data *entity.Storage) error {
	return rows.Scan(
		&data.UID,
		&data.Name,
		&data.GoodsUID,
		&data.WarehouseID,
		&data.BlockID,
		&data.Status,
		&data.XRow,
		&data.YRow,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
}
