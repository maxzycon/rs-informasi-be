package dto

type InformationKioskList struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Photo        *string `json:"photo"`
	CategoryName string  `json:"category_name"`
	CreatedAt    string  `json:"created_at"`
}

type InformationKiosk struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Photo        *string `json:"photo"`
	CategoryName string  `json:"category_name"`
	Description  *string `json:"description"`
	CreatedAt    string  `json:"created_at"`
}

type FacilitiesListKiosk struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Photo *string `json:"photo"`
}

type FacilitiesKiosk struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Photo       *string `json:"photo"`
	Description *string `json:"description"`
}

type RoomsListKiosk struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Photo *string `json:"photo"`
}

type RoomsKiosk struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Photo       *string `json:"photo"`
	Description *string `json:"description"`
}

type ServicesListKiosk struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Photo *string `json:"photo"`
}

type ServiceKiosk struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Photo       *string `json:"photo"`
	Description *string `json:"description"`
}

type ProductListKiosk struct {
	ID                string                    `json:"id"`
	Name              string                    `json:"name"`
	CategoryName      string                    `json:"category_name"`
	Photo             *string                   `json:"photo"`
	Price             float64                   `json:"price"`
	AmountDiscount    *float64                  `json:"amount_discount"`
	IsDiscount        bool                      `json:"is_discount"`
	DiscountStartDate *string                   `json:"discount_start_date"`
	DiscountEndDate   *string                   `json:"discount_end_date"`
	Detail            []*ProductListDetailKiosk `json:"detail"`
}

type ProductListDetailKiosk struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TempProductListDetailKiosk struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
}

type DoctorListKiosk struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	SpecializationName  string                 `json:"specialization_name"`
	OrganName           string                 `json:"organ_name"`
	Photo               *string                `json:"photo"`
	DoctorScheduleKiosk []*DoctorScheduleKiosk `json:"slot_detail"`
}

type DoctorKiosk struct {
	ID                   string                  `json:"id"`
	Name                 string                  `json:"name"`
	SpecializationName   string                  `json:"specialization_name"`
	OrganName            string                  `json:"organ_name"`
	Photo                *string                 `json:"photo"`
	DoctorScheduleKiosk  []*DoctorScheduleKiosk  `json:"slot_detail"`
	DoctorSkillKiosk     []*DoctorSkillKiosk     `json:"skill_detail"`
	DoctorEducationKiosk []*DoctorEducationKiosk `json:"education_detail"`
}

type DoctorScheduleKiosk struct {
	ID    string `json:"id"`
	Day   int    `json:"day"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type DoctorSkillKiosk struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DoctorEducationKiosk struct {
	ID    string `json:"id"`
	Grade string `json:"grade"`
	Major string `json:"major"`
	Name  string `json:"name"`
}

type TempDoctorScheduleKiosk struct {
	ID       string `json:"id"`
	DoctorID string `json:"doctor_id"`
	Day      int    `json:"day"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

type DoctorDashboarKiosk struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	SpecializationName string  `json:"specialization_name"`
	OrganName          string  `json:"organ_name"`
	Photo              *string `json:"photo"`
}

type ProductDashboardKiosk struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	CategoryName      string   `json:"category_name"`
	Photo             *string  `json:"photo"`
	Price             float64  `json:"price"`
	AmountDiscount    *float64 `json:"amount_discount"`
	IsDiscount        bool     `json:"is_discount"`
	DiscountStartDate *string  `json:"discount_start_date"`
	DiscountEndDate   *string  `json:"discount_end_date"`
}

type WrapperDashboardKiosk struct {
	RandomDoctors  []*DoctorDashboarKiosk   `json:"random_doctors"`
	RandomProducts []*ProductDashboardKiosk `json:"random_products"`
}
