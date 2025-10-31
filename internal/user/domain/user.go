package user

import (
	"fmt"
	"time"

	"github.com/andrewronscki/lib-golang-teste/internal/shared/utils"
	"github.com/andrewronscki/lib-golang-teste/pkg/commons/config"
)

type User struct {
	Name      string    `json:"name,omitempty"`
	SiteID    string    `json:"site_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(firstName, lastName string) *User {
	date := time.Now().UTC()

	user := &User{
		SiteID:    config.GetConfig().SiteID,
		CreatedAt: date,
		UpdatedAt: date,
	}

	user.GetName(firstName, lastName)

	return user
}

func (u *User) GetName(firstName, lastName string) {
	u.Name = fmt.Sprintf("%s %s", firstName, lastName)
}

func (u *User) Marshal(dest any) {
	utils.DeepCopy(u, dest)
}
