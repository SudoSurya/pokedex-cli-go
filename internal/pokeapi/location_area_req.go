package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURl + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, nil
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	locAreaResp := LocationAreasResp{}

	err = json.Unmarshal(data, &locAreaResp)
	if err!=nil{
		return LocationAreasResp{},err
	}
	return locAreaResp,nil

}
