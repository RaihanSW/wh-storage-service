package service

// =============================================================================
// CONTOH UNIT TEST — BlockLayoutService (CRUD di "Storage".block_layout).
//
// Jalankan:  go test ./app/service/ -run TestBlockLayout -v
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

type stubBlockLayoutRepo struct {
	repository.BlockLayoutRepositoryInterface

	insertID   int
	insertData entity.BlockLayout
	insertErr  error
	getAllData []entity.BlockLayout
	getAllErr  error
	getData    entity.BlockLayout
	getErr     error
	updateRows int64
	updateErr  error
	deleteID   int
	deleteRows int64
	deleteErr  error
}

func (s *stubBlockLayoutRepo) InsertBlockLayout(ctx context.Context, data entity.BlockLayout) (int, error) {
	s.insertData = data
	return s.insertID, s.insertErr
}

func (s *stubBlockLayoutRepo) GetAllBlockLayout(ctx context.Context) ([]entity.BlockLayout, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubBlockLayoutRepo) GetBlockLayoutByID(ctx context.Context, id int) (entity.BlockLayout, error) {
	return s.getData, s.getErr
}

func (s *stubBlockLayoutRepo) UpdateBlockLayoutByID(ctx context.Context, data entity.BlockLayout) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubBlockLayoutRepo) DeleteBlockLayoutByID(ctx context.Context, id int) (int64, error) {
	s.deleteID = id
	return s.deleteRows, s.deleteErr
}

func newBlockLayoutService(repo *stubBlockLayoutRepo) BlockLayoutServiceInterface {
	return NewBlockLayoutServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestBlockLayoutCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, id dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{insertID: 7}
		svc := newBlockLayoutService(repo)

		data, status, err := svc.BlockLayoutCreate(ctx, request.BlockLayoutRequest{WarehouseID: 1, BlockName: "A"})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.ID != 7 || data.BlockName != "A" || data.WarehouseID != 1 {
			t.Fatalf("unexpected data: %+v", data)
		}
		if repo.insertData.BlockName != "A" || repo.insertData.WarehouseID != 1 {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{insertErr: errors.New("db down")}
		svc := newBlockLayoutService(repo)

		_, status, err := svc.BlockLayoutCreate(ctx, request.BlockLayoutRequest{WarehouseID: 1, BlockName: "X"})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestBlockLayoutGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{getAllData: []entity.BlockLayout{{ID: 1}, {ID: 2}}}
		svc := newBlockLayoutService(repo)

		data, status, err := svc.BlockLayoutGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})
}

func TestBlockLayoutGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{getData: entity.BlockLayout{ID: 1, WarehouseID: 1, BlockName: "A"}}
		svc := newBlockLayoutService(repo)

		data, status, err := svc.BlockLayoutGetByID(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.ID != 1 || data.BlockName != "A" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{getErr: sql.ErrNoRows}
		svc := newBlockLayoutService(repo)

		_, status, err := svc.BlockLayoutGetByID(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestBlockLayoutUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{updateRows: 1}
		svc := newBlockLayoutService(repo)

		status, err := svc.BlockLayoutUpdate(ctx, 1, request.BlockLayoutRequest{WarehouseID: 2, BlockName: "B"})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{updateRows: 0}
		svc := newBlockLayoutService(repo)

		status, err := svc.BlockLayoutUpdate(ctx, 99, request.BlockLayoutRequest{WarehouseID: 1, BlockName: "X"})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestBlockLayoutDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, id diteruskan ke repo", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{deleteRows: 1}
		svc := newBlockLayoutService(repo)

		status, err := svc.BlockLayoutDelete(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteID != 1 {
			t.Fatalf("expect deleted id 1, got %d", repo.deleteID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubBlockLayoutRepo{deleteRows: 0}
		svc := newBlockLayoutService(repo)

		status, err := svc.BlockLayoutDelete(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
