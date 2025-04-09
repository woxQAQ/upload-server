package stores

import (
	"context"

	"github.com/woxQAQ/upload-server/internal/models"
	"gorm.io/gorm"
)

type ApproveOption struct {
	Id       uint64
	Approver string
	Opinion  *string
}

type GetOption struct {
	Id       string
	TaskId   string
	OrgName  string
	UserId   string
	Status   string
	Approver string
}

type UpdateOption struct {
	Status        string
	ApproveOption string
}

type ProgressStore interface {
	ListProgressByUser(ctx context.Context, userId string) ([]models.Progress, error)
	GetProgress(ctx context.Context, progressId uint64) (models.Progress, error)
	CreateProgress(ctx context.Context, progress models.Progress) error
	ApproveProgress(ctx context.Context, opt ApproveOption) error
	UpdateProgress(ctx context.Context, id *string, opt UpdateOption) error
}

type store struct {
	db *gorm.DB
}

// GetProgress implements ProgressStore.
func (s store) GetProgress(ctx context.Context, progressId uint64) (models.Progress, error) {
	panic("unimplemented")
}

// ApproveProgress implements ProgressStore.
func (s store) ApproveProgress(ctx context.Context, opt ApproveOption) error {
	panic("unimplemented")
}

// CreateProgress implements ProgressStore.
func (s store) CreateProgress(ctx context.Context, progress models.Progress) error {
	panic("unimplemented")
}

// ListProgress implements ProgressStore.
func (s store) ListProgressByUser(ctx context.Context, userId string) ([]models.Progress, error) {
	panic("unimplemented")
}

// UpdateProgress implements ProgressStore.
func (s store) UpdateProgress(ctx context.Context, id *string, opt UpdateOption) error {
	panic("unimplemented")
}

var _ ProgressStore = store{}

func NewProgressStore(db *gorm.DB) ProgressStore {
	return store{
		db,
	}
}
