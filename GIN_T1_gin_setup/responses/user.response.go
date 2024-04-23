package responses

import "time"

type UserReponse struct {
	Id        *int       `json:"id"`
	Username  *string    `json:"name"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}