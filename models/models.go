package models

type Embassy struct {
	HomeCountry  string       `bson:"home_country"`
	HostCountry  string       `bson:"host_country"`
	Consulate    bool         `bson:"consulate"`
	MapLink      string       `bson:"map_link"`
	City         string       `bson:"city"`
	GoogleID     string       `bson:"google_id"`
	PlaceDetails PlaceDetails `bson:"place_details"`
	//TODO: implement methods to get the place details and the google id
}

type PlaceDetails struct {
	HtmlAttributions []interface{} `json:"html_attributions" bson:"html_attributions"`
	Result           Result        `json:"result" bson:"result"`
	Status           string        `json:"status" bson:"status"`
}

type Result struct {
	AddressComponents        []AddressComponent `json:"address_components" bson:"address_components"`
	AdrAddress               string             `json:"adr_address" bson:"adr_address"`
	BusinessStatus           string             `json:"business_status" bson:"business_status"`
	CurrentOpeningHours      OpeningHours       `json:"current_opening_hours" bson:"current_opening_hours"`
	FormattedAddress         string             `json:"formatted_address" bson:"formatted_address"`
	FormattedPhoneNumber     string             `json:"formatted_phone_number" bson:"formatted_phone_number"`
	Geometry                 Geometry           `json:"geometry" bson:"geometry"`
	Icon                     string             `json:"icon" bson:"icon"`
	IconBackgroundColor      string             `json:"icon_background_color" bson:"icon_background_color"`
	InternationalPhoneNumber string             `json:"international_phone_number" bson:"international_phone_number"`
	Name                     string             `json:"name" bson:"name"`
	OpeningHours             OpeningHours       `json:"opening_hours" bson:"opening_hours"`
	Photos                   []Photo            `json:"photos" bson:"photos"`
	PlaceID                  string             `json:"place_id" bson:"place_id"`
	PlusCode                 PlusCode           `json:"plus_code" bson:"plus_code"`
	Rating                   float64            `json:"rating" bson:"rating"`
	//Reviews              []Review           `json:"reviews" bson:"reviews"`
	Types            []string `json:"types" bson:"types"`
	URL              string   `json:"url" bson:"url"`
	UserRatingsTotal int      `json:"user_ratings_total" bson:"user_ratings_total"`
	Website          string   `json:"website" bson:"website"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name" bson:"long_name"`
	ShortName string   `json:"short_name" bson:"short_name"`
	Types     []string `json:"types" bson:"types"`
}

type OpeningHours struct {
	OpenNow     bool     `json:"open_now" bson:"open_now"`
	Periods     []Period `json:"periods" bson:"periods"`
	WeekdayText []string `json:"weekday_text" bson:"weekday_text"`
}

type Period struct {
	Close TimeInfo `json:"close" bson:"close"`
	Open  TimeInfo `json:"open" bson:"open"`
}

type TimeInfo struct {
	Date string `json:"date" bson:"date"`
	Day  int    `json:"day" bson:"day"`
	Time string `json:"time" bson:"time"`
}

type Geometry struct {
	Location Location `json:"location" bson:"location"`
	Viewport Viewport `json:"viewport" bson:"viewport"`
}

type Location struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

type Viewport struct {
	Northeast Coordinates `json:"northeast" bson:"northeast"`
	Southwest Coordinates `json:"southwest" bson:"southwest"`
}

type Coordinates struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

type Photo struct {
	Height           int      `json:"height" bson:"height"`
	HtmlAttributions []string `json:"html_attributions" bson:"html_attributions"`
	PhotoReference   string   `json:"photo_reference" bson:"photo_reference"`
	Width            int      `json:"width" bson:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code" bson:"compound_code"`
	GlobalCode   string `json:"global_code" bson:"global_code"`
}

/*type Review struct {
	AuthorName             string `json:"author_name" bson:"author_name"`
	AuthorURL              string `json:"author_url" bson:"author_url"`
	ProfilePhotoURL        string `json:"profile_photo_url" bson:"profile_photo_url"`
	Rating                 int    `json:"rating" bson:"rating"`
	RelativeTimeDescription string `json:"relative_time_description" bson:"relative_time_description"`
	Text                   string `json:"text" bson:"text"`
	Time                   int64  `json:"time" bson:"time"`
	Translated             bool   `json:"translated" bson:"translated"`
}
*/
