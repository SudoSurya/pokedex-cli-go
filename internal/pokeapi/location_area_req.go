package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURl + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache Hit!!!!")
		locAreaResp := LocationAreasResp{}

		err := json.Unmarshal(data, &locAreaResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locAreaResp, nil
	}
	fmt.Println("Cache Miss!!!")

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
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	locAreaResp := LocationAreasResp{}

	err = json.Unmarshal(data, &locAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	c.cache.Add(fullUrl, data)
	return locAreaResp, nil

}

func (c *Client) GetLocationAres(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseURl + endpoint

	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache Hit!!!!")
		locAreaResp := LocationArea{}

		err := json.Unmarshal(data, &locAreaResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locAreaResp, nil
	}
	fmt.Println("Cache Miss!!!")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, nil
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locAreaResp := LocationArea{}

	err = json.Unmarshal(data, &locAreaResp)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullUrl, data)
	return locAreaResp, nil

}
