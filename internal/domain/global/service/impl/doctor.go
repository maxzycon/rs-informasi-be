package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/authutil"
	"github.com/maxzycon/rs-informasi-be/pkg/model"
	"github.com/maxzycon/rs-informasi-be/pkg/util/pagination"
)

func (s *GlobalService) GetDoctorPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	resp, err = s.globalRepository.FindDoctorPaginated(ctx, payload)
	if err != nil {
		s.log.Error("err get Doctor paginated")
		return
	}

	respToDto := make([]*dto.DoctorRow, 0)
	list, ok := resp.Items.([]*model.Doctor)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.DoctorRow{
				ID:                 v.ID,
				Name:               v.Name,
				SpecializationName: v.Specialization.Name,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *GlobalService) GetDoctorPluck(ctx context.Context) (resp []*dto.DefaultPluck, err error) {
	rows, err := s.globalRepository.FindAllDoctor(ctx)
	if err != nil {
		s.log.Error("err get Doctor paginated")
		return
	}
	resp = make([]*dto.DefaultPluck, 0)
	for _, row := range rows {
		resp = append(resp, &dto.DefaultPluck{
			ID:   row.ID,
			Name: row.Name,
		})
	}
	return
}

func (s *GlobalService) GetDoctorById(ctx context.Context, id int) (resp *dto.DoctorByIdRow, err error) {
	row, err := s.globalRepository.FindDoctorById(ctx, id)
	if err != nil {
		s.log.Error("err get Doctor paginated")
		return
	}
	resp = &dto.DoctorByIdRow{
		ID:               row.ID,
		Name:             row.Name,
		SpecializationID: row.SpecializationID,
		Photo:            row.Photo,
		Skills:           make([]*dto.DoctorSkillsRow, 0),
		Educations:       make([]*dto.DoctorEducationRow, 0),
		Slots:            make([]*dto.DoctorSlotRow, 0),
	}

	if row.Photo != nil {
		temp := s.conf.AWS_S3_URL + "/" + *row.Photo
		resp.Photo = &temp
	}

	for _, v := range row.Education {
		resp.Educations = append(resp.Educations, &dto.DoctorEducationRow{
			ID:    v.ID,
			Grade: v.Grade,
			Major: v.Major,
			Name:  v.Name,
		})
	}

	for _, v := range row.Skill {
		resp.Skills = append(resp.Skills, &dto.DoctorSkillsRow{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	for _, v := range row.Slot {
		resp.Slots = append(resp.Slots, &dto.DoctorSlotRow{
			ID:        v.ID,
			Day:       v.Day,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	return
}

func (s *GlobalService) CreateDoctor(ctx context.Context, payload *dto.PayloadDoctor) (resp *int64, err error) {
	user, _ := authutil.GetCredential(ctx)
	entity := &model.Doctor{
		Name:             payload.Name,
		SpecializationID: payload.SpecializationID,
		MerchantID:       *user.MerchantID,
		Photo:            payload.Photo,
	}

	for _, v := range payload.Educations {
		entity.Education = append(entity.Education, model.DoctorEducation{
			Grade: v.Grade,
			Major: v.Major,
			Name:  v.Name,
		})
	}

	for _, v := range payload.Skills {
		entity.Skill = append(entity.Skill, model.DoctorSkill{
			Name: v.Name,
		})
	}

	for _, v := range payload.Slots {
		entity.Slot = append(entity.Slot, model.DoctorSlot{
			Day:       v.Day,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	resp, err = s.globalRepository.CreateDoctor(ctx, entity)
	if err != nil {
		s.log.Error("err Doctor status")
		return
	}
	return
}

func (s *GlobalService) UpdateDoctorById(ctx context.Context, id int, payload *dto.PayloadDoctor) (resp *int64, err error) {
	row, err := s.GetDoctorById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	entity := &model.Doctor{
		Name:             payload.Name,
		SpecializationID: payload.SpecializationID,
	}

	if payload.Photo != nil && *row.Photo != *payload.Photo {
		entity.Photo = payload.Photo
	}

	for _, v := range payload.Educations {
		entity.Education = append(entity.Education, model.DoctorEducation{
			Grade: v.Grade,
			Major: v.Major,
			Name:  v.Name,
		})
	}

	for _, v := range payload.Skills {
		entity.Skill = append(entity.Skill, model.DoctorSkill{
			Name: v.Name,
		})
	}

	for _, v := range payload.Slots {
		entity.Slot = append(entity.Slot, model.DoctorSlot{
			Day:       v.Day,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	err = s.db.Model(&model.DoctorSkill{}).Unscoped().Where("doctor_id = ?", id).Delete(&model.DoctorSkill{}).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = s.db.Model(&model.DoctorEducation{}).Unscoped().Where("doctor_id = ?", id).Delete(&model.DoctorEducation{}).Error
	if err != nil {
		return nil, err
	}

	err = s.db.Model(&model.DoctorSlot{}).Unscoped().Where("doctor_id = ?", id).Delete(&model.DoctorSlot{}).Error
	if err != nil {
		return nil, err
	}

	resp, err = s.globalRepository.UpdateDoctorById(ctx, id, entity)
	if err != nil {
		s.log.Errorf("err update Doctor %d", id)
		return
	}

	return
}

func (s *GlobalService) DeleteDoctorById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeleteDoctorById(ctx, id)
	if err != nil {
		s.log.Errorf("err delete Doctor %d", id)
		return
	}
	return
}
