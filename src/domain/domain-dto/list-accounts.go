package domaindto

import "time"

type ListAccountsOutputDTO struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt" valid:"-"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
}
