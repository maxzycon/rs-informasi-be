package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (r *GlobalRepository) FindDoctorById(ctx context.Context, id int) (resp *model.Doctor, err error) {
	tx := r.db.WithContext(ctx).Preload("Skill").Preload("Education").Preload("Slot").First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllDoctor(ctx context.Context) (resp []*model.Doctor, err error) {
	resp = make([]*model.Doctor, 0)
	tx := r.db.WithContext(ctx).Model(&model.Doctor{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindDoctorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Doctors []*model.Doctor = make([]*model.Doctor, 0)
	sql := r.db.Debug().WithContext(ctx).Preload("Specialization")
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Doctors, &resp.Paginator, sql)).Find(&Doctors)
	resp.Items = Doctors
	return
}

func (r *GlobalRepository) CreateDoctor(ctx context.Context, entity *model.Doctor) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Doctor{}).Create(&entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) UpdateDoctorById(ctx context.Context, id int, entity *model.Doctor) (resp *int64, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return &tx.RowsAffected, tx.Error
}

func (r *GlobalRepository) DeleteDoctorById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Doctor{}, id)
	return &tx.RowsAffected, tx.Error
}