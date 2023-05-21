package models

import (
	"time"

	"gorm.io/gorm"
)

type KofiShopOrder struct {
	gorm.Model
	VerificationToken          string           `json:"verification_token"`
	MessageID                  string           `json:"message_id"`
	Timestamp                  time.Time        `json:"timestamp"`
	Type                       string           `json:"type"`
	IsPublic                   bool             `json:"is_public"`
	FromName                   string           `json:"from_name"`
	Message                    string           `json:"message"`
	Amount                     string           `json:"amount"`
	URL                        string           `json:"url"`
	Email                      string           `json:"email"`
	Currency                   string           `json:"currency"`
	IsSubscriptionPayment      bool             `json:"is_subscription_payment"`
	IsFirstSubscriptionPayment bool             `json:"is_first_subscription_payment"`
	KofiTransactionID          string           `json:"kofi_transaction_id"`
	ShopItems                  []KofiShopItem   `json:"shop_items"`
	TierName                   string           `json:"tier_name"`
	Shipping                   KofiShippingInfo `json:"shipping"`
}

type KofiShopItem struct {
	DirectLinkCode string `json:"direct_link_code"`
	VariationName  string `json:"variation_name"`
	Quantity       int    `json:"quantity"`
}

type KofiShippingInfo struct {
	FullName        string `json:"full_name"`
	StreetAddress   string `json:"street_address"`
	City            string `json:"city"`
	StateOrProvince string `json:"state_or_province"`
	PostalCode      string `json:"postal_code"`
	Country         string `json:"country"`
	CountryCode     string `json:"country_code"`
	Telephone       string `json:"telephone"`
}
