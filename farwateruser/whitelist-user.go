package farwateruser

import "github.com/farwater-create/backend/models"

func WhitelistUser(user *models.User) {

}

func init() {
	AddEventListener(EventRegister, WhitelistUser)
}
