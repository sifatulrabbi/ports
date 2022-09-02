package validators

import (
	"github.com/sifatulrabbi/ports/pkg/models"
)

func RegisterPayload(u *models.User) bool {
	if u.Username != "" && u.Fullname != "" && u.Password != "" && u.Bio != "" {
		return true
	}
	return false
}
