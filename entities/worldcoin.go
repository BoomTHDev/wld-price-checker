package entities

type WorldCoin struct {
	Data struct {
		WLD struct {
			Quote struct {
				THB struct {
					Price float64 `json:"price"`
				} `json:"THB"`
			} `json:"quote"`
		} `json:"WLD"`
	} `json:"data"`
}
