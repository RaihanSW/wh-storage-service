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

type blockLayoutRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewBlockLayoutRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) BlockLayoutRepositoryInterface {
	return &blockLayoutRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *blockLayoutRepository) InsertBlockLayout(ctx context.Context, data entity.BlockLayout) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var id int
	err := r.db.QueryRowContext(ctx, query.InsertBlockLayout,
		data.WarehouseID,
		data.BlockName,
		data.Floor,
		data.XSum,
		data.YSum,
		r.serviceName,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.InsertBlockLayout (Scan): %w", err)
	}

	return id, nil
}

func (r *blockLayoutRepository) GetAllBlockLayout(ctx context.Context) ([]entity.BlockLayout, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllBlockLayout)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllBlockLayout (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.BlockLayout
	for rows.Next() {
		var data entity.BlockLayout
		if err := scanBlockLayout(rows, &data); err != nil {
			return nil, fmt.Errorf("repo.GetAllBlockLayout (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllBlockLayout (Rows): %w", err)
	}

	return list, nil
}

func (r *blockLayoutRepository) GetBlockLayoutByID(ctx context.Context, id int) (entity.BlockLayout, error) {
	var data entity.BlockLayout

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetBlockLayoutByID, id).Scan(
		&data.ID,
		&data.WarehouseID,
		&data.BlockName,
		&data.Floor,
		&data.XSum,
		&data.YSum,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.BlockLayout{}, fmt.Errorf("repo.GetBlockLayoutByID (Scan): %w", err)
	}

	return data, nil
}

func (r *blockLayoutRepository) UpdateBlockLayoutByID(ctx context.Context, data entity.BlockLayout) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateBlockLayoutByID,
		data.WarehouseID,
		data.BlockName,
		data.Floor,
		data.XSum,
		data.YSum,
		r.serviceName,
		data.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateBlockLayoutByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateBlockLayoutByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *blockLayoutRepository) DeleteBlockLayoutByID(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteBlockLayoutByID, id)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteBlockLayoutByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteBlockLayoutByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func scanBlockLayout(rows *sql.Rows, data *entity.BlockLayout) error {
	return rows.Scan(
		&data.ID,
		&data.WarehouseID,
		&data.BlockName,
		&data.Floor,
		&data.XSum,
		&data.YSum,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
}
