package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/maxzycon/rs-informasi-be/internal/domain/global/dto"
	"github.com/maxzycon/rs-informasi-be/pkg/util/timeutil"
)

func (s *GlobalService) GetInformationListKiosk(ctx context.Context, categoryId int, merchantUUID string) (resp []*dto.InformationKioskList, err error) {
	cond := squirrel.And{squirrel.Eq{
		"m.id_str":     merchantUUID,
		"i.deleted_at": nil,
	}}

	if categoryId > 0 {
		cond = append(cond, squirrel.Eq{
			"i.information_category_id": categoryId,
		})
	}

	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo, ic.name as category_information, i.created_at").
		From("information as i").
		LeftJoin("information_categories as ic ON ic.id = i.information_category_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(cond).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}
	fmt.Println(merchantUUID)
	fmt.Println(sqlStr)
	rows, err := s.db.Raw(sqlStr, args...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	resp = make([]*dto.InformationKioskList, 0)

	for rows.Next() {
		temp := dto.InformationKioskList{}
		var tempTime time.Time
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.CategoryName, &tempTime)
		if err != nil {
			s.log.Error("err scan information")
			return
		}

		temp.CreatedAt = timeutil.ToString(tempTime)

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		resp = append(resp, &temp)
	}
	return
}

func (s *GlobalService) GetInformationKiosk(ctx context.Context, id string, merchantUUID string) (resp *dto.InformationKiosk, err error) {
	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo, ic.name as category_information, i.created_at, i.desc").
		From("information as i").
		LeftJoin("information_categories as ic ON ic.id = i.information_category_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.id_str":     id,
			"i.deleted_at": nil,
		}).
		ToSql()

	fmt.Println(sqlStr)

	if err != nil {
		s.log.Error("err generate sql")
		return
	}

	row := s.db.Raw(sqlStr, args...).Row()
	if row.Err() != nil {
		s.log.Error("err raw rows")
		return
	}

	temp := dto.InformationKiosk{}
	var tempTime time.Time
	err = row.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.CategoryName, &tempTime, &temp.Description)
	if err != nil {
		s.log.Error("err scan information")
		return
	}

	temp.CreatedAt = timeutil.ToString(tempTime)

	if temp.Photo != nil {
		s := s.conf.AWS_S3_URL + "/" + *temp.Photo
		temp.Photo = &s
	}

	resp = &temp
	return
}

func (s *GlobalService) GetFacilitiesListKiosk(ctx context.Context, merchantUUID string) (resp []*dto.FacilitiesListKiosk, err error) {

	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo").
		From("facilities as i").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.deleted_at": nil,
		}).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}
	fmt.Println(merchantUUID)
	fmt.Println(sqlStr)
	rows, err := s.db.Raw(sqlStr, args...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	resp = make([]*dto.FacilitiesListKiosk, 0)

	for rows.Next() {
		temp := dto.FacilitiesListKiosk{}
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Photo)
		if err != nil {
			s.log.Error("err scan information")
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		resp = append(resp, &temp)
	}
	return
}

func (s *GlobalService) GetFacilitieskioskById(ctx context.Context, id string, merchantUUID string) (resp *dto.FacilitiesKiosk, err error) {

	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo, i.desc").
		From("facilities as i").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.id_str":     id,
			"i.deleted_at": nil,
		}).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}
	row := s.db.Raw(sqlStr, args...).Row()
	if row.Err() != nil {
		s.log.Error("err raw rows")
		return
	}

	temp := dto.FacilitiesKiosk{}
	err = row.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.Description)
	if err != nil {
		s.log.Error("err scan information")
		return
	}

	if temp.Photo != nil {
		s := s.conf.AWS_S3_URL + "/" + *temp.Photo
		temp.Photo = &s
	}

	resp = &temp

	return
}

func (s *GlobalService) GetRoomsListKiosk(ctx context.Context, floorId int, merchantUUID string) (resp []*dto.RoomsListKiosk, err error) {

	cond := squirrel.And{
		squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.deleted_at": nil,
		},
	}

	if floorId > 0 {
		cond = append(cond, squirrel.Eq{
			"i.floor_id": floorId,
		})
	}

	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo").
		From("rooms as i").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(cond).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}
	fmt.Println(merchantUUID)
	fmt.Println(sqlStr)
	rows, err := s.db.Raw(sqlStr, args...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	resp = make([]*dto.RoomsListKiosk, 0)

	for rows.Next() {
		temp := dto.RoomsListKiosk{}
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Photo)
		if err != nil {
			s.log.Error("err scan information")
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		resp = append(resp, &temp)
	}
	return
}

func (s *GlobalService) GetRoomskioskById(ctx context.Context, id string, merchantUUID string) (resp *dto.RoomsKiosk, err error) {
	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo, i.description").
		From("rooms as i").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.id_str":     id,
			"i.deleted_at": nil,
		}).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}
	row := s.db.Raw(sqlStr, args...).Row()
	if row.Err() != nil {
		s.log.Errorf("err raw rows %v", row.Err())
		return
	}

	temp := dto.RoomsKiosk{}
	err = row.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.Description)
	if err != nil {
		s.log.Error("err scan information")
		return
	}

	if temp.Photo != nil {
		s := s.conf.AWS_S3_URL + "/" + *temp.Photo
		temp.Photo = &s
	}

	resp = &temp

	return
}

func (s *GlobalService) GetServicesListKiosk(ctx context.Context, merchantUUID string) (resp []*dto.ServicesListKiosk, err error) {
	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo").
		From("services as i").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.deleted_at": nil,
		}).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}

	rows, err := s.db.Raw(sqlStr, args...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	resp = make([]*dto.ServicesListKiosk, 0)

	for rows.Next() {
		temp := dto.ServicesListKiosk{}
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Photo)
		if err != nil {
			s.log.Error("err scan service")
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		resp = append(resp, &temp)
	}
	return
}

func (s *GlobalService) GetServicekioskById(ctx context.Context, id string, merchantUUID string) (resp *dto.ServiceKiosk, err error) {
	sqlStr, args, err := squirrel.
		Select("i.id_str, i.name, i.photo, i.desc").
		From("services as i").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(squirrel.Eq{
			"m.id_str":     merchantUUID,
			"i.id_str":     id,
			"i.deleted_at": nil,
		}).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql")
		return
	}
	row := s.db.Raw(sqlStr, args...).Row()
	if row.Err() != nil {
		s.log.Errorf("err raw rows %v", row.Err())
		return
	}

	temp := dto.ServiceKiosk{}
	err = row.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.Description)
	if err != nil {
		s.log.Error("err scan service")
		return
	}

	if temp.Photo != nil {
		s := s.conf.AWS_S3_URL + "/" + *temp.Photo
		temp.Photo = &s
	}

	resp = &temp

	return
}

func (s *GlobalService) GetProductsListKiosk(ctx context.Context, categoryId int, merchantUUID string) (resp []*dto.ProductListKiosk, err error) {
	cond := squirrel.And{
		squirrel.Eq{
			"i.deleted_at": nil,
			"m.id_str":     merchantUUID,
		},
	}

	if categoryId > 0 {
		cond = append(cond, squirrel.Eq{
			"i.product_category_id": categoryId,
		})
	}

	productId := []int{}

	rootSql, argsRoot, err := squirrel.
		Select("i.id_str, i.id, i.name, i.price, i.is_discount, i.amount_discount, i.discount_start_date, i.discount_end_date, i.photo, pc.name").
		From("products as i").
		LeftJoin("product_categories as pc ON pc.id = i.product_category_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(cond).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql parent")
		return
	}

	rows, err := s.db.Raw(rootSql, argsRoot...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	resp = make([]*dto.ProductListKiosk, 0)

	for rows.Next() {
		temp := dto.ProductListKiosk{}
		var tempId int64
		var tempStartDate, tempEndDate *time.Time
		err = rows.Scan(&temp.ID, &tempId, &temp.Name, &temp.Price, &temp.IsDiscount, &temp.AmountDiscount, &tempStartDate, &tempEndDate, &temp.Photo, &temp.CategoryName)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		if temp.IsDiscount {
			s := timeutil.ToString(*tempStartDate)
			e := timeutil.ToString(*tempEndDate)
			temp.DiscountStartDate = &s
			temp.DiscountEndDate = &e
		}

		productId = append(productId, int(tempId))

		temp.Detail = make([]*dto.ProductListDetailKiosk, 0)
		resp = append(resp, &temp)
	}

	childSql, argsChild, err := squirrel.
		Select("i.id_str, p.id_str, i.name").
		From("detail_products as i").
		LeftJoin("products as p ON p.id = i.product_id").
		Where(squirrel.Eq{
			"i.product_id": productId,
		}).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql child")
		return
	}

	rows, err = s.db.Raw(childSql, argsChild...).Rows()
	if err != nil {
		s.log.Error("err raw rows child")
		return
	}

	tempDetail := make([]*dto.TempProductListDetailKiosk, 0)
	for rows.Next() {
		tempD := dto.TempProductListDetailKiosk{}

		err = rows.Scan(&tempD.ID, &tempD.ProductID, &tempD.Name)
		if err != nil {
			s.log.Errorf("err generate sql child %v", err)
			return
		}
		tempDetail = append(tempDetail, &tempD)
	}

	// ------ join to response
	for _, v := range resp {
		for _, c := range tempDetail {
			if v.ID == c.ProductID {
				v.Detail = append(v.Detail, &dto.ProductListDetailKiosk{
					ID:   c.ID,
					Name: c.Name,
				})
			}
		}
	}

	return
}

func (s *GlobalService) GetDoctorsListKiosk(ctx context.Context, specializationId int, day int, merchantUUID string) (resp []*dto.DoctorListKiosk, err error) {
	cond := squirrel.And{
		squirrel.Eq{
			"i.deleted_at": nil,
			"m.id_str":     merchantUUID,
		},
	}

	if specializationId > 0 {
		cond = append(cond, squirrel.Eq{
			"i.specialization_id": specializationId,
		})
	}

	doctorIds := []string{}

	p := squirrel.
		Select("i.id_str, i.name, i.photo, pc.name, og.name").
		From("doctors as i").
		LeftJoin("specializations as pc ON pc.id = i.specialization_id").
		// LeftJoin("organs as og ON og.id = pc.organ_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id")

	if day > 0 {
		p = p.Join("doctor_slots as ds ON ds.doctor_id = i.id AND ds.day = ?", day)
	}

	rootSql, argsRoot, err := p.Where(cond).
		GroupBy("i.id").
		ToSql()

	if err != nil {
		s.log.Error("err generate sql parent")
		return
	}

	rows, err := s.db.Raw(rootSql, argsRoot...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	resp = make([]*dto.DoctorListKiosk, 0)

	for rows.Next() {
		temp := dto.DoctorListKiosk{}
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.SpecializationName)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		doctorIds = append(doctorIds, temp.ID)

		temp.DoctorScheduleKiosk = make([]*dto.DoctorScheduleKiosk, 0)
		resp = append(resp, &temp)
	}

	slotSql, slotArgs, err := squirrel.
		Select("i.id_str, i.day, i.start_time, i.end_time, d.id_str").
		From("doctor_slots as i").
		LeftJoin("doctors as d ON d.id = i.doctor_id").
		Where(squirrel.Eq{"d.id_str": doctorIds}).
		OrderBy("i.day ASC, i.start_time ASC, i.end_time ASC").
		ToSql()

	if err != nil {
		s.log.Error("err generate sql slot str")
		return
	}

	rows, err = s.db.Raw(slotSql, slotArgs...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	tmpSlot := make([]*dto.TempDoctorScheduleKiosk, 0)

	for rows.Next() {
		temp := dto.TempDoctorScheduleKiosk{}
		err = rows.Scan(&temp.ID, &temp.Day, &temp.Start, &temp.End, &temp.DoctorID)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}
		tmpSlot = append(tmpSlot, &temp)
	}

	for _, v := range resp {
		for _, c := range tmpSlot {
			if v.ID == c.DoctorID {
				v.DoctorScheduleKiosk = append(v.DoctorScheduleKiosk, &dto.DoctorScheduleKiosk{
					ID:    c.ID,
					Day:   c.Day,
					Start: c.Start,
					End:   c.End,
				})
			}
		}
	}

	return
}

func (s *GlobalService) GetDoctorsKiosk(ctx context.Context, id string, merchantUUID string) (resp *dto.DoctorKiosk, err error) {
	cond := squirrel.And{
		squirrel.Eq{
			"i.deleted_at": nil,
			"m.id_str":     merchantUUID,
			"i.id_str":     id,
		},
	}

	rootSql, argsRoot, err := squirrel.
		Select("i.id_str, i.name, i.photo, pc.name, og.name").
		From("doctors as i").
		LeftJoin("specializations as pc ON pc.id = i.specialization_id").
		// LeftJoin("organs as og ON og.id = pc.organ_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(cond).ToSql()

	if err != nil {
		s.log.Error("err generate sql parent")
		return
	}

	row := s.db.Raw(rootSql, argsRoot...).Row()
	if row.Err() != nil {
		s.log.Errorf("err raw rows %v", row.Err())
		return
	}

	temp := dto.DoctorKiosk{}
	err = row.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.SpecializationName)
	if err != nil {
		s.log.Errorf("err scan service : %v", err)
		return
	}

	if temp.Photo != nil {
		s := s.conf.AWS_S3_URL + "/" + *temp.Photo
		temp.Photo = &s
	}
	temp.DoctorScheduleKiosk = make([]*dto.DoctorScheduleKiosk, 0)
	temp.DoctorSkillKiosk = make([]*dto.DoctorSkillKiosk, 0)
	temp.DoctorEducationKiosk = make([]*dto.DoctorEducationKiosk, 0)

	// ------------------- SLOT

	slotSql, slotArgs, err := squirrel.
		Select("i.id_str, i.day, i.start_time, i.end_time").
		From("doctor_slots as i").
		LeftJoin("doctors as d ON d.id = i.doctor_id").
		Where(squirrel.Eq{"d.id_str": temp.ID}).
		OrderBy("i.day ASC, i.start_time ASC, i.end_time ASC").
		ToSql()

	if err != nil {
		s.log.Error("err generate sql slot str")
		return
	}

	rows, err := s.db.Raw(slotSql, slotArgs...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	for rows.Next() {
		tmp := dto.DoctorScheduleKiosk{}
		err = rows.Scan(&tmp.ID, &tmp.Day, &tmp.Start, &tmp.End)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}
		temp.DoctorScheduleKiosk = append(temp.DoctorScheduleKiosk, &tmp)
	}

	// ----------------- SKILLS

	skillSql, skillArgs, err := squirrel.
		Select("i.id_str, i.name").
		From("doctor_skills as i").
		LeftJoin("doctors as d ON d.id = i.doctor_id").
		Where(squirrel.Eq{"d.id_str": temp.ID}).
		OrderBy("i.id ASC").
		ToSql()

	if err != nil {
		s.log.Error("err generate sql slot str")
		return
	}

	rows, err = s.db.Raw(skillSql, skillArgs...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	for rows.Next() {
		tmp := dto.DoctorSkillKiosk{}
		err = rows.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}

		temp.DoctorSkillKiosk = append(temp.DoctorSkillKiosk, &tmp)
	}

	// --------------------- EDUCATIONS

	educationSql, educationArgs, err := squirrel.
		Select("i.id_str, i.grade, i.major, i.name").
		From("doctor_educations as i").
		LeftJoin("doctors as d ON d.id = i.doctor_id").
		Where(squirrel.Eq{"d.id_str": temp.ID}).
		OrderBy("i.id ASC").
		ToSql()

	if err != nil {
		s.log.Error("err generate sql slot str")
		return
	}

	rows, err = s.db.Raw(educationSql, educationArgs...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	for rows.Next() {
		tmp := dto.DoctorEducationKiosk{}
		err = rows.Scan(&tmp.ID, &tmp.Grade, &tmp.Major, &tmp.Name)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}

		temp.DoctorEducationKiosk = append(temp.DoctorEducationKiosk, &tmp)
	}

	resp = &temp
	return
}

func (s *GlobalService) GetDashboardKiosk(ctx context.Context, merchantIdStr string) (resp *dto.WrapperDashboardKiosk, err error) {
	resp = &dto.WrapperDashboardKiosk{
		RandomDoctors:  make([]*dto.DoctorDashboarKiosk, 0),
		RandomProducts: make([]*dto.ProductDashboardKiosk, 0),
	}
	cond := squirrel.And{
		squirrel.Eq{
			"i.deleted_at": nil,
			"m.id_str":     merchantIdStr,
		},
	}

	rootSql, argsRoot, err := squirrel.
		Select("i.id_str, i.name, i.photo, pc.name, og.name").
		From("doctors as i").
		LeftJoin("specializations as pc ON pc.id = i.specialization_id").
		// LeftJoin("organs as og ON og.id = pc.organ_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(cond).
		OrderBy("RAND()").
		Limit(5).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql parent")
		return
	}

	rows, err := s.db.Raw(rootSql, argsRoot...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	randDoctor := make([]*dto.DoctorDashboarKiosk, 0)

	for rows.Next() {
		temp := dto.DoctorDashboarKiosk{}
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Photo, &temp.SpecializationName)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}
		randDoctor = append(randDoctor, &temp)
	}

	rootSql, argsRoot, err = squirrel.
		Select("i.id_str, i.id, i.name, i.price, i.is_discount, i.amount_discount, i.discount_start_date, i.discount_end_date, i.photo, pc.name").
		From("products as i").
		LeftJoin("product_categories as pc ON pc.id = i.product_category_id").
		LeftJoin("merchants as m ON m.id = i.merchant_id").
		Where(cond).
		OrderBy("RAND()").
		Limit(5).
		ToSql()

	if err != nil {
		s.log.Error("err generate sql parent")
		return
	}

	rows, err = s.db.Raw(rootSql, argsRoot...).Rows()
	if err != nil {
		s.log.Error("err raw rows")
		return
	}

	randProduct := make([]*dto.ProductDashboardKiosk, 0)

	for rows.Next() {
		temp := dto.ProductDashboardKiosk{}
		var tempId int64
		var tempStartDate, tempEndDate *time.Time
		err = rows.Scan(&temp.ID, &tempId, &temp.Name, &temp.Price, &temp.IsDiscount, &temp.AmountDiscount, &tempStartDate, &tempEndDate, &temp.Photo, &temp.CategoryName)
		if err != nil {
			s.log.Errorf("err scan service : %v", err)
			return
		}

		if temp.Photo != nil {
			s := s.conf.AWS_S3_URL + "/" + *temp.Photo
			temp.Photo = &s
		}

		if temp.IsDiscount {
			s := timeutil.ToString(*tempStartDate)
			e := timeutil.ToString(*tempEndDate)
			temp.DiscountStartDate = &s
			temp.DiscountEndDate = &e
		}

		randProduct = append(randProduct, &temp)
	}

	resp.RandomDoctors = randDoctor
	resp.RandomProducts = randProduct

	return
}
