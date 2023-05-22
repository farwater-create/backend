package kofi

import "github.com/farwater-create/backend/models"

type KofiShopOrderProcessor = func(order *models.KofiShopOrder)

var kofiShopOrderEventListeners = map[KofiEvent][]KofiShopOrderProcessor{}
var kofiShopMembershipTierListeners = map[string][]KofiShopOrderProcessor{}

type KofiEvent string

const (
	EventShopOrder    = "Shop Order"
	EventSubscription = "Subscription"
	EventDonation     = "Donation"
)

func init() {
	// event listener to proc kofiShopMembershipTierListeners
	AddEventListener(EventSubscription, func(order *models.KofiShopOrder) {
		for _, listener := range kofiShopMembershipTierListeners[order.TierName] {
			listener(order)
		}
	})
}

func AddEventListener(event KofiEvent, listener KofiShopOrderProcessor) {
	kofiShopOrderEventListeners[event] = append(kofiShopOrderEventListeners[event], listener)
}
