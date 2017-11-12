package entities

type Alive struct {
	Commit string `json:"commit"`
}

type LocationServiceResponse struct {
	Suppliers []struct {
		ID            string        `json:"id"`
		Status        string        `json:"status"`
		Divisions     []interface{} `json:"divisions"`
		LastLocations []struct {
			Lat      float64 `json:"lat"`
			Lon      float64 `json:"lon"`
			Ts       int64   `json:"ts"`
			Bearing  float32 `json:"bearing"`
			Accuracy int     `json:"accuracy"`
			Speed    float32 `json:"speed"`
		} `json:"last_locations"`
		LastUpdate int64 `json:"last_update"`
	} `json:"drivers"`
	Rc int `json:"rc"`
}
