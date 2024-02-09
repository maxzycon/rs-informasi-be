package dto

type UserRow struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	NIK         string  `json:"nik"`
	Password    string  `json:"password"`
	ProfilePath *string `json:"profile_path"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Role        uint    `json:"role"`
	MerchantID  *uint   `json:"merchant_id"`
}

type UserRowDetail struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	NIK         string  `json:"nik"`
	ProfilePath *string `json:"profile_path"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Role        uint    `json:"role"`
	MerchantID  *uint   `json:"merchant_id"`
}

type PayloadUpdateUser struct {
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	NIK         string  `json:"nik"`
	Password    *string `json:"password"`
	ProfilePath *string `json:"profile_path"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Role        uint    `json:"role"`
	MerchantID  *uint   `json:"merchant_id"`
}

type PayloadLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type PayloadUpdateProfile struct {
	Password string `json:"password"`
}
type PayloadCreateUser struct {
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	NIK         string  `json:"nik"`
	Password    string  `json:"password"`
	ProfilePath *string `json:"profile_path"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Role        uint    `json:"role"`
	MerchantID  *uint   `json:"merchant_id"`
}

type LoginRes struct {
	ID          uint    `json:"id"`
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Photo       *string `json:"photo"`
	AccessToken string  `json:"access_token"`
	Role        uint    `json:"role"`
	Exp         int64   `json:"exp"`
}

type UserPaginatedRow struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Merchant string `json:"merchant"`
	Role     uint   `json:"role"`
}
