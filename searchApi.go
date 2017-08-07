package main

type Search struct {
	Site_id           string
	Query             string
	Paging            Paging
	Results           []Result
	Secondary_results []SecondaryResult
	Related_results   []RelatedResult
	Sort              Sort
	Available_Sorts   []Sort
	Filters           []Filter
	Available_filters []Filter
}

type Paging struct {
	Total           int64
	Offset          int64
	Limit           int64
	Primary_results int64
}

type Result struct {
	Id                   string
	Site_id              string
	Title                string
	Seller               Seller
	Price                int64
	Currency_id          string
	Available_quantity   int32
	Sold_quantity        int32
	Buying_mode          string
	Listing_type_id      string
	Stop_time            string
	Condition            string
	Permalink            string
	Thumbnail            string
	Accepts_mercadopago  bool
	Installments         Installment
	Address              Address
	Shipping             Shipping
	Seller_Address       SellerAddress
	Attributes           []Attribute
	Differential_pricing DifferentialPricing
	Original_price       int32
	Category_id          string
	Official_store_id    int32
	Catalog_product_id   string
	Reviews              Review
}

type Installment struct {
	Quantity    int32
	Amount      float64
	Rate        float64
	Currency_id string
}

type Seller struct {
	Id                  int64
	Power_seller_status string
	Car_dealer          bool
	Real_estate_agency  bool
	Tags                []Tag
}

type Tag struct {
}

type Address struct {
	State_id   string
	State_name string
	City_id    string
	City_name  string
}

type Shipping struct {
	Free_shipping bool
	Mode          string
	Tags          []Tag
}

type SellerAddress struct {
	Id           int64
	Comment      string
	Address_line string
	Zip_code     string
	Country      Location
	State        Location
	City         Location
	Latitude     float64
	Longitude    float64
}

type Location struct {
	Id   string
	Name string
}

type Attribute struct {
	Id                   string
	Name                 string
	Value_id             string
	Value_name           string
	Attribute_group_id   string
	Attribute_group_name string
}

type DifferentialPricing struct {
	Id int64
}

type Review struct {
	Rating_average float32
	Total          int64
}

type SecondaryResult struct {
}

type RelatedResult struct {
}

type Sort struct {
	Id   string
	Name string
}

type Filter struct {
	ID     string
	Name   string
	Type   string
	Values []Value
}

type Value struct {
	ID      string
	Name    string
	Results int64
}
