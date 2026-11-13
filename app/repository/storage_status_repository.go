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

type storageStatusRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewStorageStatusRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) StorageStatusRepositoryInterface {
	return &storageStatusRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *storageStatusRepository) InsertStorageStatus(ctx context.Context, data entity.StorageStatus) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var id int
	err := r.db.QueryRowContext(ctx, query.InsertStorageStatus, data.Name, r.serviceName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.InsertStorageStatus (Scan): %w", err)
	}

	return id, nil
}

func (r *storageStatusRepository) GetAllStorageStatus(ctx context.Context) ([]entity.StorageStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllStorageStatus)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllStorageStatus (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.StorageStatus
	for rows.Next() {
		var data entity.StorageStatus
		if err := rows.Scan(&data.ID, &data.Name, &data.CreatedAt, &data.CreatedBy, &data.UpdatedAt, &data.UpdatedBy); err != nil {
			return nil, fmt.Errorf("repo.GetAllStorageStatus (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllStorageStatus (Rows): %w", err)
	}

	return list, nil
}

func (r *storageStatusRepository) GetStorageStatusByID(ctx context.Context, id int) (entity.StorageStatus, error) {
	var data entity.StorageStatus

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetStorageStatusByID, id).Scan(
		&data.ID,
		&data.Name,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.StorageStatus{}, fmt.Errorf("repo.GetStorageStatusByID (Scan): %w", err)
	}

	return data, nil
}

func (r *storageStatusRepository) UpdateStorageStatusByID(ctx context.Context, data entity.StorageStatus) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateStorageStatusByID, data.Name, r.serviceName, data.ID)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateStorageStatusByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateStorageStatusByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *storageStatusRepository) DeleteStorageStatusByID(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteStorageStatusByID, id)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteStorageStatusByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteStorageStatusByID (RowsAffected): %w", err)
	}

	return ra, nil
}
