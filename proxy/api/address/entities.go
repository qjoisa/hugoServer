package address

import "github.com/ekomobile/dadata/v2/api/model"

type correctAddresser interface {
	toCorrectAddress() []*Address
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}

// TransformData оставляет только информацию о городе
func (r *ResponseAddress) TransformData(model correctAddresser) {
	r.Addresses = model.toCorrectAddress()
}

type Address struct {
	Country  string `json:"country"`
	Region   string `json:"region"`
	City     string `json:"city"`
	Timezone string `json:"timezone"`
	GeoLat   string `json:"geo_lat"`
	GeoLon   string `json:"geo_lon"`
}

type SearchResponse []*model.Address

func (s SearchResponse) toCorrectAddress() []*Address {
	addresses := make([]*Address, 0)
	for _, address := range s {
		addresses = append(addresses, &Address{
			Country:  address.Country,
			Region:   address.Region,
			City:     address.City,
			Timezone: address.Timezone,
			GeoLat:   address.GeoLat,
			GeoLon:   address.GeoLon,
		})
		break
	}
	return addresses
}

type GeocodeResponse struct {
	Suggestions []Suggestions `json:"suggestions"`
}

func (g GeocodeResponse) toCorrectAddress() []*Address {
	addresses := make([]*Address, 0)
	for _, a := range g.Suggestions {
		addresses = append(addresses, &Address{
			Country:  a.Data.Country,
			Region:   a.Data.Region,
			City:     a.Data.City,
			Timezone: a.Data.Timezone,
			GeoLat:   a.Data.GeoLat,
			GeoLon:   a.Data.GeoLon,
		})
		break
	}
	return addresses
}

type Data struct {
	PostalCode           string `json:"postal_code"`
	Country              string `json:"country"`
	CountryIsoCode       string `json:"country_iso_code"`
	FederalDistrict      string `json:"federal_district"`
	RegionFiasID         string `json:"region_fias_id"`
	RegionKladrID        string `json:"region_kladr_id"`
	RegionIsoCode        string `json:"region_iso_code"`
	RegionWithType       string `json:"region_with_type"`
	RegionType           string `json:"region_type"`
	RegionTypeFull       string `json:"region_type_full"`
	Region               string `json:"region"`
	AreaFiasID           any    `json:"area_fias_id"`
	AreaKladrID          any    `json:"area_kladr_id"`
	AreaWithType         any    `json:"area_with_type"`
	AreaType             any    `json:"area_type"`
	AreaTypeFull         any    `json:"area_type_full"`
	Area                 any    `json:"area"`
	CityFiasID           string `json:"city_fias_id"`
	CityKladrID          string `json:"city_kladr_id"`
	CityWithType         string `json:"city_with_type"`
	CityType             string `json:"city_type"`
	CityTypeFull         string `json:"city_type_full"`
	City                 string `json:"city"`
	CityArea             string `json:"city_area"`
	CityDistrictFiasID   any    `json:"city_district_fias_id"`
	CityDistrictKladrID  any    `json:"city_district_kladr_id"`
	CityDistrictWithType any    `json:"city_district_with_type"`
	CityDistrictType     any    `json:"city_district_type"`
	CityDistrictTypeFull any    `json:"city_district_type_full"`
	CityDistrict         any    `json:"city_district"`
	SettlementFiasID     any    `json:"settlement_fias_id"`
	SettlementKladrID    any    `json:"settlement_kladr_id"`
	SettlementWithType   any    `json:"settlement_with_type"`
	SettlementType       any    `json:"settlement_type"`
	SettlementTypeFull   any    `json:"settlement_type_full"`
	Settlement           any    `json:"settlement"`
	StreetFiasID         string `json:"street_fias_id"`
	StreetKladrID        string `json:"street_kladr_id"`
	StreetWithType       string `json:"street_with_type"`
	StreetType           string `json:"street_type"`
	StreetTypeFull       string `json:"street_type_full"`
	Street               string `json:"street"`
	SteadFiasID          any    `json:"stead_fias_id"`
	SteadCadnum          any    `json:"stead_cadnum"`
	SteadType            any    `json:"stead_type"`
	SteadTypeFull        any    `json:"stead_type_full"`
	Stead                any    `json:"stead"`
	HouseFiasID          string `json:"house_fias_id"`
	HouseKladrID         string `json:"house_kladr_id"`
	HouseCadnum          string `json:"house_cadnum"`
	HouseType            string `json:"house_type"`
	HouseTypeFull        string `json:"house_type_full"`
	House                string `json:"house"`
	BlockType            any    `json:"block_type"`
	BlockTypeFull        any    `json:"block_type_full"`
	Block                any    `json:"block"`
	Entrance             any    `json:"entrance"`
	Floor                any    `json:"floor"`
	FlatFiasID           any    `json:"flat_fias_id"`
	FlatCadnum           any    `json:"flat_cadnum"`
	FlatType             any    `json:"flat_type"`
	FlatTypeFull         any    `json:"flat_type_full"`
	Flat                 any    `json:"flat"`
	FlatArea             any    `json:"flat_area"`
	SquareMeterPrice     string `json:"square_meter_price"`
	FlatPrice            any    `json:"flat_price"`
	RoomFiasID           any    `json:"room_fias_id"`
	RoomCadnum           any    `json:"room_cadnum"`
	RoomType             any    `json:"room_type"`
	RoomTypeFull         any    `json:"room_type_full"`
	Room                 any    `json:"room"`
	PostalBox            any    `json:"postal_box"`
	FiasID               string `json:"fias_id"`
	FiasCode             any    `json:"fias_code"`
	FiasLevel            string `json:"fias_level"`
	FiasActualityState   string `json:"fias_actuality_state"`
	KladrID              string `json:"kladr_id"`
	GeonameID            string `json:"geoname_id"`
	CapitalMarker        string `json:"capital_marker"`
	Okato                string `json:"okato"`
	Oktmo                string `json:"oktmo"`
	TaxOffice            string `json:"tax_office"`
	TaxOfficeLegal       string `json:"tax_office_legal"`
	Timezone             string `json:"timezone"`
	GeoLat               string `json:"geo_lat"`
	GeoLon               string `json:"geo_lon"`
	BeltwayHit           string `json:"beltway_hit"`
	BeltwayDistance      any    `json:"beltway_distance"`
	Metro                any    `json:"metro"`
	Divisions            any    `json:"divisions"`
	QcGeo                string `json:"qc_geo"`
	QcComplete           any    `json:"qc_complete"`
	QcHouse              any    `json:"qc_house"`
	HistoryValues        any    `json:"history_values"`
	UnparsedParts        any    `json:"unparsed_parts"`
	Source               any    `json:"source"`
	Qc                   any    `json:"qc"`
}

type Suggestions struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              Data   `json:"data"`
}
