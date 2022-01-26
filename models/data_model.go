package model

type Data struct {
	Imei             string `json:"imei"`
	Serial           string `json:"serial"`
	Momsn            string `json:"momsn"`
	TransmitTime     string `json:"transmit_time"`
	IridiumLatitude  string `json:"iridium_latitude"`
	IridiumLongitude string `json:"iridium_longitude"`
	IridiumCep       string `json:"iridium_cep"`
	Data             string `json:"data"`
}

type NewSpaceData struct {
	Imei             *string `json:"imei"`
	Serial           *string `json:"serial"`
	Momsn            *string `json:"momsn"`
	TransmitTime     *string `json:"transmit_time"`
	IridiumLatitude  *string `json:"iridium_latitude"`
	IridiumLongitude *string `json:"iridium_longitude"`
	IridiumCep       *string `json:"iridium_cep"`
	Data             *string `json:"data"`
}

type SpaceData struct {
	SpaceID          string  `json:"spaceID"`
	Imei             *string `json:"imei"`
	Serial           *string `json:"serial"`
	Momsn            *string `json:"momsn"`
	TransmitTime     *string `json:"transmit_time"`
	IridiumLatitude  *string `json:"iridium_latitude"`
	IridiumLongitude *string `json:"iridium_longitude"`
	IridiumCep       *string `json:"iridium_cep"`
	Data             *string `json:"data"`
}
