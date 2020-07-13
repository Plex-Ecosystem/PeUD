package v1

type PlexResponse struct {
	PlexUsers []PlexUser `xml:"User"`
}

type TautulliResponse struct {
	Response struct {
		Data []TautulliUser `json:"data"`
	} `json:"response"`
}

type OrganizrResponse struct {
	Data struct {
		Users []OrganizrUser `json:"users"`
	} `json:"data"`
}

type OmbiResponseUserClaim struct {
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}
