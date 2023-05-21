package kofi

import "github.com/farwater-create/backend/models"

func WhitelistPremium(order *models.KofiShopOrder) {

}

func init() {
	AddEventListener(EventSubscription, WhitelistPremium)
}
