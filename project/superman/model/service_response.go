package model

type ServerListResponse struct {
	//GlobalVersion  string
	Items          []*ServiceItem
	ServiceVersion int64
	PSM            string
}

type ServerVersionResponse struct {
	GlobalVersion int64
	VersionMap    map[string]interface{}
}
