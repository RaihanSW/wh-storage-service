package service

// =============================================================================
// CONTOH UNIT TEST — StorageStatusService (lookup id+name).
//
// Jalankan:  go test ./app/service/ -run TestStorageStatus -v
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

type stubStorageStatusRepo struct {
	repository.StorageStatusRepositoryInterface

	insertID   int
	insertData entity.StorageStatus
	insertErr  error
	getAllData []entity.StorageStatus
	getAllErr  error
	getData    entity.StorageStatus
	getErr     error
	updateRows int64
	updateErr  error
	deleteID   int
	deleteRows int64
	deleteErr  error
}

func (s *stubStorageStatusRepo) InsertStorageStatus(ctx context.Context, data entity.StorageStatus) (int, error) {
	s.insertData = data
	return s.insertID, s.insertErr
}

func (s *stubStorageStatusRepo) GetAllStorageStatus(ctx context.Context) ([]entity.StorageStatus, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubStorageStatusRepo) GetStorageStatusByID(ctx context.Context, id int) (entity.StorageStatus, error) {
	return s.getData, s.getErr
}

func (s *stubStorageStatusRepo) UpdateStorageStatusByID(ctx context.Context, data entity.StorageStatus) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubStorageStatusRepo) DeleteStorageStatusByID(ctx context.Context, id int) (int64, error) {
	s.deleteID = id
	return s.deleteRows, s.deleteErr
}

func newStorageStatusService(repo *stubStorageStatusRepo) StorageStatusServiceInterface {
	return NewStorageStatusServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestStorageStatusCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, id dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubStorageStatusRepo{insertID: 7}
		svc := newStorageStatusService(repo)

		data, status, err := svc.StorageStatusCreate(ctx, request.StorageStatusRequest{Name: "Occupied"})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.ID != 7 || data.Name != "Occupied" {
			t.Fatalf("unexpected data: %+v", data)
		}
		if repo.insertData.Name != "Occupied" {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubStorageStatusRepo{insertErr: errors.New("db down")}
		svc := newStorageStatusService(repo)

		_, status, err := svc.StorageStatusCreate(ctx, request.StorageStatusRequest{Name: "X"})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestStorageStatusGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubStorageStatusRepo{getAllData: []entity.StorageStatus{{ID: 1}, {ID: 2}}}
		svc := newStorageStatusService(repo)

		data, status, err := svc.StorageStatusGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})
}

func TestStorageStatusGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubStorageStatusRepo{getData: entity.StorageStatus{ID: 1, Name: "Empty"}}
		svc := newStorageStatusService(repo)

		data, status, err := svc.StorageStatusGetByID(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.ID != 1 || data.Name != "Empty" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubStorageStatusRepo{getErr: sql.ErrNoRows}
		svc := newStorageStatusService(repo)

		_, status, err := svc.StorageStatusGetByID(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestStorageStatusUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubStorageStatusRepo{updateRows: 1}
		svc := newStorageStatusService(repo)

		status, err := svc.StorageStatusUpdate(ctx, 1, request.StorageStatusRequest{Name: "Renamed"})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubStorageStatusRepo{updateRows: 0}
		svc := newStorageStatusService(repo)

		status, err := svc.StorageStatusUpdate(ctx, 99, request.StorageStatusRequest{Name: "X"})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestStorageStatusDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, id diteruskan ke repo", func(t *testing.T) {
		repo := &stubStorageStatusRepo{deleteRows: 1}
		svc := newStorageStatusService(repo)

		status, err := svc.StorageStatusDelete(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteID != 1 {
			t.Fatalf("expect deleted id 1, got %d", repo.deleteID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubStorageStatusRepo{deleteRows: 0}
		svc := newStorageStatusService(repo)

		status, err := svc.StorageStatusDelete(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
