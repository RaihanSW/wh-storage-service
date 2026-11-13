package service

// =============================================================================
// CONTOH UNIT TEST — WarehouseService (CRUD di "Storage".warehouse).
//
// Jalankan:  go test ./app/service/ -run TestWarehouse -v
// =============================================================================

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"
	"wh-storage-service/app/model/entity"
	"wh-storage-service/app/model/request"
	"wh-storage-service/app/repository"
	"wh-storage-service/config"
)

type stubWarehouseRepo struct {
	repository.WarehouseRepositoryInterface

	insertID   int
	insertData entity.Warehouse
	insertErr  error
	getAllData []entity.Warehouse
	getAllErr  error
	getData    entity.Warehouse
	getErr     error
	updateRows int64
	updateErr  error
	deleteID   int
	deleteRows int64
	deleteErr  error
}

func (s *stubWarehouseRepo) InsertWarehouse(ctx context.Context, data entity.Warehouse) (int, error) {
	s.insertData = data
	return s.insertID, s.insertErr
}

func (s *stubWarehouseRepo) GetAllWarehouse(ctx context.Context) ([]entity.Warehouse, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubWarehouseRepo) GetWarehouseByID(ctx context.Context, id int) (entity.Warehouse, error) {
	return s.getData, s.getErr
}

func (s *stubWarehouseRepo) UpdateWarehouseByID(ctx context.Context, data entity.Warehouse) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubWarehouseRepo) DeleteWarehouseByID(ctx context.Context, id int) (int64, error) {
	s.deleteID = id
	return s.deleteRows, s.deleteErr
}

func newWarehouseService(repo *stubWarehouseRepo) WarehouseServiceInterface {
	return NewWarehouseServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestWarehouseCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, id dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubWarehouseRepo{insertID: 7}
		svc := newWarehouseService(repo)

		data, status, err := svc.WarehouseCreate(ctx, request.WarehouseRequest{Name: "WH Jakarta", Code: "WHJKT"})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.ID != 7 || data.Name != "WH Jakarta" || data.Code != "WHJKT" {
			t.Fatalf("unexpected data: %+v", data)
		}
		if repo.insertData.Name != "WH Jakarta" || repo.insertData.Code != "WHJKT" {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubWarehouseRepo{insertErr: errors.New("db down")}
		svc := newWarehouseService(repo)

		_, status, err := svc.WarehouseCreate(ctx, request.WarehouseRequest{Name: "X", Code: "X"})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestWarehouseGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubWarehouseRepo{getAllData: []entity.Warehouse{{ID: 1}, {ID: 2}}}
		svc := newWarehouseService(repo)

		data, status, err := svc.WarehouseGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})
}

func TestWarehouseGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubWarehouseRepo{getData: entity.Warehouse{ID: 1, Name: "WH Jakarta", Code: "WHJKT"}}
		svc := newWarehouseService(repo)

		data, status, err := svc.WarehouseGetByID(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.ID != 1 || data.Name != "WH Jakarta" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubWarehouseRepo{getErr: sql.ErrNoRows}
		svc := newWarehouseService(repo)

		_, status, err := svc.WarehouseGetByID(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestWarehouseUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubWarehouseRepo{updateRows: 1}
		svc := newWarehouseService(repo)

		status, err := svc.WarehouseUpdate(ctx, 1, request.WarehouseRequest{Name: "WH Depok", Code: "WHDPK"})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubWarehouseRepo{updateRows: 0}
		svc := newWarehouseService(repo)

		status, err := svc.WarehouseUpdate(ctx, 99, request.WarehouseRequest{Name: "X", Code: "X"})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestWarehouseDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, id diteruskan ke repo", func(t *testing.T) {
		repo := &stubWarehouseRepo{deleteRows: 1}
		svc := newWarehouseService(repo)

		status, err := svc.WarehouseDelete(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteID != 1 {
			t.Fatalf("expect deleted id 1, got %d", repo.deleteID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubWarehouseRepo{deleteRows: 0}
		svc := newWarehouseService(repo)

		status, err := svc.WarehouseDelete(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
