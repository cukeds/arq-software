// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sales struct {
	Period    string `json:"period"`
	Completed int    `json:"completed"`
}

type Delay struct {
	Period string `json:"period"`
	Rate   int    `json:"rate"`
	Value  int    `json:"value"`
}

type Claim struct {
	Period string  `json:"period"`
	Rate   float32 `json:"rate"`
	Value  int     `json:"value"`
}

type Cancellation struct {
	Period string `json:"period"`
	Rate   int    `json:"rate"`
	Value  int    `json:"value"`
}

type Metric struct {
	Cancellations       Cancellation `json:"cancellations"`
	Claims              Claim        `json:"claims"`
	DelayedHandlingTime Delay        `json:"delayed_handling_time"`
	Sales               Sales        `json:"sales"`
}

type Ratings struct {
	Negative float32 `json:"negative"`
	Neutral  float32 `json:"neutral"`
	Positive float32 `json:"positive"`
}

type Transactions struct {
	Cancelled int     `json:"cancelled"`
	Period    string  `json:"period"`
	Total     int     `json:"total"`
	Ratings   Ratings `json:"ratings"`
	Completed int     `json:"completed"`
}

type Reputation struct {
	PowerSellerStatus string       `json:"power_seller_status"`
	LevelId           string       `json:"level_id"`
	Metrics           Metric       `json:"metrics"`
	Transactions      Transactions `json:"transactions"`
}

type Seller struct {
	Id               int        `json:"id"`
	Permalink        string     `json:"permalink"`
	RegistrationDate string     `json:"registration_date"`
	CarDealer        bool       `json:"car_dealer"`
	RealEstateAgency bool       `json:"real_estate_agency"`
	Tags             []string   `json:"tags"`
	SellerReputaton  Reputation `json:"seller_reputation"`
}

type ConditionsDos struct {
	ContextRestrictions []string `json:"context_restrictions"`
	StartTime           string   `json:"start_time"`
	EndTime             string   `json:"end_time"`
	Eligible            bool     `json:"eligible"`
}

type Price struct {
	Id            string        `json:"id"`
	Tipo          string        `json:"type"`
	Amount        int           `json:"amount"`
	RegularAmount int           `json:"regular_amount"`
	CurrencyId    string        `json:"currency_id"`
	LastUpdated   string        `json:"last_updated"`
	Conditions    ConditionsDos `json:"conditions"`
	ExchangedRate string        `json:"exchange_rate"`
	Metadata      string        `json:"metadata"`
}

type ReferencePrice struct {
	Id            string        `json:"id"`
	Tipo          string        `json:"type"`
	Conditions    ConditionsDos `json:"conditions"`
	Amount        int           `json:"amount"`
	CurrencyId    string        `json:"currency_id"`
	ExchangedRate string        `json:"exchange_rate"`
	Tags          []string      `json:"tags"`
	LastUpdated   string        `json:"last_updated"`
}

type Presentation struct {
	DisplayCurrency     string           `json:"display_currency"`
	PaymentMethodPrices []string         `json:"payment_method_prices"`
	ReferencePrices     []ReferencePrice `json:"reference_prices"`
	PurchaseDiscounts   []string         `json:"purchase_discounts"`
}

type PricesDos struct {
	Id           string       `json:"id"`
	Prices       []Price      `json:"prices"`
	Presentation Presentation `json:"presentation"`
}

type Installments struct {
	Quantity   int     `json:"quantity"`
	Amount     float32 `json:"amount"`
	Rate       int     `json:"rate"`
	CurrencyId string  `json:"currency_id"`
}

type Address struct {
	StateId   string `json:"state_id"`
	StateName string `json:"state_name"`
	CityId    string `json:"city_id"`
	CityName  string `json:"city_name"`
}

type Shipping struct {
	FreeShipping bool     `json:"free_shipping"`
	Mode         string   `json:"mode"`
	Tags         []string `json:"tags"`
	LogisticType string   `json:"logistic_type"`
	StorePickUp  bool     `json:"store_pick_up"`
}

type Country struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type State struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SellerAddress struct {
	Id          string  `json:"id"`
	Comment     string  `json:"comment"`
	AddressLine string  `json:"address_line"`
	ZipCode     string  `json:"zip_code"`
	Country     Country `json:"country"`
	State       State   `json:"state"`
	City        City    `json:"city"`
	Latitude    string  `json:"latitude"`
	Longitude   string  `json:"longitude"`
}

type Value struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Structura string `json:"struct"`
	Source    int    `json:"source"`
}

type Attribute struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	ValueName          string  `json:"value_name"`
	ValueStruct        string  `json:"value_struct"`
	Values             []Value `json:"values"`
	AttributeGroupName string  `json:"attribute_group_name"`
	ValueId            string  `json:"value_id"`
	AttributeGroupId   string  `json:"attribute_group_id"`
	Source             int     `json:"source"`
}
type DifferentialPricing struct {
	Id int `json:"id"`
}

type Result struct {
	Id                  string              `json:"id"`
	SiteId              string              `json:"site_id"`
	Title               string              `json:"title"`
	Seller              Seller              `json:"seller"`
	Price               int                 `json:"price"`
	Prices              PricesDos           `json:"prices"`
	SalePrice           int                 `json:"sale_price"`
	CurrencyId          string              `json:"currency_id"`
	AvailableQuantity   int                 `json:"available_quantity"`
	SoldQuantity        int                 `json:"sold_quantity"`
	BuyingMode          string              `json:"buying_mode"`
	ListingTypeId       string              `json:"listing_type_id"`
	StopTime            string              `json:"stop_time"`
	Condition           string              `json:"condition"`
	Permalink           string              `json:"permalink"`
	Thumbnail           string              `json:"thumbnail"`
	ThumbnailId         string              `json:"thumbnail_id"`
	AcceptsMercadopago  bool                `json:"accepts_mercadopago"`
	Installments        Installments        `json:"installments"`
	Address             Address             `json:"address"`
	Shipping            Shipping            `json:"shipping"`
	SellerAddress       SellerAddress       `json:"seller_address"`
	Attributes          []Attribute         `json:"attributes"`
	DifferentialPricing DifferentialPricing `json:"differential_pricing"`
	OriginalPrice       int                 `json:"original_price"`
	CategoryId          string              `json:"category_id"`
	OfficialStoreId     int                 `json:"official_store_id"`
	DomainId            string              `json:"domain_id"`
	CatalogProductId    string              `json:"catalog_product_id"`
	Tags                []string            `json:"tags"`
	CatalogListing      bool                `json:"catalog_listing"`
	UseThumbnailId      bool                `json:"use_thumbnail_id"`
	OfferScore          string              `json:"offer_score"`
	OfferShare          string              `json:"offer_share"`
	MatchScore          string              `json:"match_score"`
	WinnerItemId        string              `json:"winner_item_id"`
	Melicoin            string              `json:"melicoin"`
	OrderBackend        int                 `json:"order_backend"`
}

type Paging struct {
	Total          int `json:"total"`
	PrimaryResults int `json:"primary_results"`
	Offset         int `json:"offset"`
	Limit          int `json:"limit"`
}

type Sort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ValueDos struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	PathFromRoot []Category `json:"pathFromRoot"`
}

type Filter struct {
	Id     string     `json:"id"`
	Name   string     `json:"name"`
	Tipo   string     `json:"type"`
	Values []ValueDos `json:"categories"`
}

type Motorola struct {
	SiteId           string   `json:"site_id"`
	CountryTimeZone  string   `json:"country_time_zone"`
	Query            string   `json:"query"`
	Paging           Paging   `json:"paging"`
	Results          []Result `json:"results"`
	Sort             Sort     `json:"sort"`
	AvailableSorts   []Sort   `json:"available_sorts"`
	Filters          []Filter `json:"filters"`
	AvailableFilters []Filter `json:"available_filters"`
}

func GetCategories(url string) (Motorola, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("NO SE AHHHHH")
	}
	bytes, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("NO SE AHHHHH")
	}

	var moto Motorola
	json.Unmarshal(bytes, &moto)
	return moto, nil
}

func main() {
	var url string = "https://api.mercadolibre.com/sites/MLA/search?q=Motorola"
	cats, _ := GetCategories(url)
	for _, cat := range cats.Results {
		fmt.Println(cat.CurrencyId)
	}

}
