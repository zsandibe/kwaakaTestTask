package domain

type ApiResponse struct {
	City string `json:"name"`
	Main struct {
		Kelvin float32 `json:"temp"`
	} `json:"main"`
}
