package impl

import (
	"context"

	"github.com/maxzycon/rs-farmasi-be/pkg/model"
)

func (r *GlobalRepository) FindQueueById(ctx context.Context, id int) (resp *model.Queue, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllQueue(ctx context.Context) (resp []*model.Queue, err error) {
	resp = make([]*model.Queue, 0)
	tx := r.db.WithContext(ctx).Model(&model.Queue{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) CreateQueue(ctx context.Context, entity *model.Queue) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Queue{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateQueueById(ctx context.Context, id int, entity *model.Queue) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteQueueById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Queue{}, id)
	return &tx.RowsAffected, tx.Error
}
