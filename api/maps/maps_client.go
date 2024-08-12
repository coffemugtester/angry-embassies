package maps

import (
	"angry-embassies/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var _ ClientInterface = (*Client)(nil)

// Client mapzImplementation
type Client struct {
	apiKey string
}

func (c Client) GetGoogleID(placeQuery string) (id string) {

	placeQuery = url.QueryEscape(placeQuery)

	mapsURL := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=" + placeQuery + "&inputtype=textquery&fields=place_id&key=" + c.apiKey

	fmt.Printf("placeQuery: %v\n", placeQuery)

	fmt.Printf("apiKey: %v\n", c.apiKey)

	fmt.Printf("Encoded URL: %v\n", mapsURL)

	response, err := http.Get(mapsURL)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	var item struct {
		Candidates []struct {
			PlaceID string `json:"place_id"`
		}
	}

	if err := json.Unmarshal(body, &item); err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	id = item.Candidates[0].PlaceID

	return id
}

func (c Client) GetPlaceDetails(placeID string) (models.PlaceDetails, error) {

	var details models.PlaceDetails

	mapsURL := "https://maps.googleapis.com/maps/api/place/details/json?fields=name,rating,opening_hours,website,address_components,adr_address,business_status,formatted_address,formatted_phone_number,geometry,rating,user_ratings_total,reviews,opening_hours,photos,current_opening_hours,editorial_summary,icon,icon_background_color,place_id,plus_code,secondary_opening_hours,types,url,website,wheelchair_accessible_entrance,international_phone_number&reviews_sort=newest&reviews_no_translations=true&place_id=" + placeID + "&key=" + c.apiKey

	fmt.Printf("Encoded URL: %v\n", mapsURL)

	response, err := http.Get(mapsURL)
	if err != nil {
		return details, fmt.Errorf("error: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return details, fmt.Errorf("error: %v", err)
	}

	if response.StatusCode != 200 {
		return details, fmt.Errorf("error: %v", err)
	}

	if err := json.Unmarshal(body, &details); err != nil {
		// TODO: use a struct or a pointer so it can be null
		return details, fmt.Errorf("error: %v", err)
	}

	return details, nil
}

// NewMapsClient creates a new mapz client
func NewMapsClient(apiKey string) *Client { // add apikey as a parameter
	return &Client{
		apiKey: apiKey,
	}
}
