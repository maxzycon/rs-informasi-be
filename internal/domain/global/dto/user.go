package dto

type UserRowPluck struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role uint   `json:"role"`
}
