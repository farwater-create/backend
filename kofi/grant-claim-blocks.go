package kofi

import "github.com/farwater-create/backend/models"

func GrantClaimBlocks(order *models.KofiShopOrder) {

}

func init() {
	AddEventListener(EventSubscription, GrantClaimBlocks)
}
