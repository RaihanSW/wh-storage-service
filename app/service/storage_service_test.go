package service

// =============================================================================
// CONTOH UNIT TEST — StorageService (CRUD di "Storage".storage).
//
// Jalankan:  go test ./app/service/ -run TestStorage -v
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

type stubStorageRepo struct {
	repository.StorageRepositoryInterface

	insertUID  string
	insertData entity.Storage
	insertErr  error
	getAllData []entity.Storage
	getAllErr  error
	getData    entity.Storage
	getErr     error
	updateRows int64
	updateErr  error
	deleteUID  string
	deleteRows int64
	deleteErr  error
}

func (s *stubStorageRepo) InsertStorage(ctx context.Context, data entity.Storage) (string, error) {
	s.insertData = data
	return s.insertUID, s.insertErr
}

func (s *stubStorageRepo) GetAllStorage(ctx context.Context) ([]entity.Storage, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubStorageRepo) GetStorageByUID(ctx context.Context, uid string) (entity.Storage, error) {
	return s.getData, s.getErr
}

func (s *stubStorageRepo) UpdateStorageByUID(ctx context.Context, data entity.Storage) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubStorageRepo) DeleteStorageByUID(ctx context.Context, uid string) (int64, error) {
	s.deleteUID = uid
	return s.deleteRows, s.deleteErr
}

func newStorageService(repo *stubStorageRepo) StorageServiceInterface {
	return NewStorageServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestStorageCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, uid dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubStorageRepo{insertUID: "FEDCBA0987654321"}
		svc := newStorageService(repo)

		data, status, err := svc.StorageCreate(ctx, request.StorageRequest{
			Name:        "A-01-01",
			WarehouseID: 1,
			BlockID:     2,
			Status:      1,
		})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.UID != "FEDCBA0987654321" {
			t.Fatalf("expect uid from repo, got %s", data.UID)
		}
		if repo.insertData.Name != "A-01-01" || repo.insertData.WarehouseID != 1 || repo.insertData.BlockID != 2 {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubStorageRepo{insertErr: errors.New("db down")}
		svc := newStorageService(repo)

		_, status, err := svc.StorageCreate(ctx, request.StorageRequest{Name: "S1", WarehouseID: 1, BlockID: 1, Status: 1})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestStorageGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubStorageRepo{getAllData: []entity.Storage{{UID: "S1"}, {UID: "S2"}}}
		svc := newStorageService(repo)

		data, status, err := svc.StorageGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubStorageRepo{getAllErr: errors.New("conn lost")}
		svc := newStorageService(repo)

		_, status, err := svc.StorageGetAll(ctx)
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestStorageGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubStorageRepo{getData: entity.Storage{UID: "S1", Name: "A-01-01"}}
		svc := newStorageService(repo)

		data, status, err := svc.StorageGetByID(ctx, "S1")
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.UID != "S1" || data.Name != "A-01-01" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubStorageRepo{getErr: sql.ErrNoRows}
		svc := newStorageService(repo)

		_, status, err := svc.StorageGetByID(ctx, "NOPE")
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestStorageUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubStorageRepo{updateRows: 1}
		svc := newStorageService(repo)

		status, err := svc.StorageUpdate(ctx, "S1", request.StorageRequest{Name: "A-01-02", WarehouseID: 1, BlockID: 2, Status: 2})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubStorageRepo{updateRows: 0}
		svc := newStorageService(repo)

		status, err := svc.StorageUpdate(ctx, "NOPE", request.StorageRequest{})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubStorageRepo{updateErr: errors.New("db down")}
		svc := newStorageService(repo)

		status, err := svc.StorageUpdate(ctx, "S1", request.StorageRequest{})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestStorageDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, uid diteruskan ke repo", func(t *testing.T) {
		repo := &stubStorageRepo{deleteRows: 1}
		svc := newStorageService(repo)

		status, err := svc.StorageDelete(ctx, "S1")
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteUID != "S1" {
			t.Fatalf("expect deleted uid S1, got %s", repo.deleteUID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubStorageRepo{deleteRows: 0}
		svc := newStorageService(repo)

		status, err := svc.StorageDelete(ctx, "NOPE")
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
