package request

type Criminal_Request struct {
	Full_Name           *string  `json:"full_name"`
	Crime_Description   *string  `json:"crime_description"`
	Danger_Level        *int     `json:"danger_level"`
	Last_Known_Location *string  `json:"last_known_location"`
	Is_Captured         *bool    `json:"is_captured"`
	Reward_Amount       *float64 `json:"reward_amount"`
}
