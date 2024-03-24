package dto

type MerchantSpecializationRow struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	OrganID   uint   `json:"organ_id"`
	OrganName string `json:"organ_name"`
}

type PayloadMerchantSpecialization struct {
	Name    string `json:"name"`
	OrganID uint   `json:"organ_id"`
}
