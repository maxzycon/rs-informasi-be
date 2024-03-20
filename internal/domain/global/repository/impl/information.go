package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindInformationById(ctx context.Context, id int) (resp *model.Information, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllInformation(ctx context.Context) (resp []*model.Information, err error) {
	resp = make([]*model.Information, 0)
	tx := r.db.WithContext(ctx).Model(&model.Information{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindInformationPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Informations []*model.Information = make([]*model.Information, 0)
	sql := r.db.Debug().WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Informations, &resp.Paginator, sql)).Find(&Informations)
	resp.Items = Informations
	return
}

func (r *GlobalRepository) CreateInformation(ctx context.Context, entity *model.Information) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Information{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateInformationById(ctx context.Context, id int, entity *model.Information) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteInformationById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Information{}, id)
	return &tx.RowsAffected, tx.Error
}
