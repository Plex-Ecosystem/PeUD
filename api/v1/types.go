package v1

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/DirtyCajunRice/go-utility"
)

// PlexUserServerSlice is an abstract type to re-implement sql methods
type PlexUserServerSlice []PlexUserServer

// Scan re-implements the database/sql Scan() method.
// It will map the strings found in the response to
// to new PlexUserServer structs in a PlexUserServerSlice
func (p *PlexUserServerSlice) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		serverSlice := make(PlexUserServerSlice, 0)
		for _, s := range strings.Split(src.(string), ",") {
			if n, err := strconv.Atoi(s); err != nil {
				return err
			} else {
				serverSlice = append(serverSlice, PlexUserServer{ID: n})
			}
		}
		*p = serverSlice
	default:
		return errors.New("incompatible type for PlexUserServerSlice")
	}
	return nil
}

// Value re-implements the database/sql Value() method.
// It will map the integers found in the response to
// a comma delimited string
func (p PlexUserServerSlice) Value() (driver.Value, error) {
	serverIDs := make([]int, 0)
	for _, server := range p {
		serverIDs = append(serverIDs, server.ID)
	}
	return utility.IntSliceToString(serverIDs, ","), nil
}

// UnmarshalJSON re-implements the encoding/json Unmarshal method.
// It will map the user claim array to a coherent struct and map
// it to the parent struct
func (c *OmbiUserClaim) UnmarshalJSON(data []byte) error {
	var claims []OmbiResponseUserClaim
	if err := json.Unmarshal(data, &claims); err != nil {
		return err
	}
	userClaim := OmbiUserClaim{}
	for _, claim := range claims {
		switch claim.Value {
		case "RequestTv":
			userClaim.RequestTv = claim.Enabled
		case "RequestMovie":
			userClaim.RequestMovie = claim.Enabled
		case "AutoApproveMovie":
			userClaim.AutoApproveMovie = claim.Enabled
		case "Admin":
			userClaim.Admin = claim.Enabled
		case "AutoApproveTv":
			userClaim.AutoApproveTv = claim.Enabled
		case "AutoApproveMusic":
			userClaim.AutoApproveMusic = claim.Enabled
		case "RequestMusic":
			userClaim.RequestMusic = claim.Enabled
		case "PowerUser":
			userClaim.PowerUser = claim.Enabled
		case "Disabled":
			userClaim.Disabled = claim.Enabled
		case "ReceivesNewsletter":
			userClaim.ReceivesNewsletter = claim.Enabled
		case "ManageOwnRequests":
			userClaim.ManageOwnRequests = claim.Enabled
		case "EditCustomPage":
			userClaim.EditCustomPage = claim.Enabled
		}
	}
	*c = userClaim
	return nil
}

// Scan re-implements the database/sql Scan() method.
// It will map the string found in the response to
// to a new OmbiUserClaim struct as the ID
func (c *OmbiUserClaim) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		*c = OmbiUserClaim{UserID: src.(string)}
	default:
		return errors.New("incompatible type for OmbiUserClaim")
	}
	return nil
}

// Value re-implements the database/sql Value() method.
// It will map the Ombi userID found in the response to
// the field.
func (c OmbiUserClaim) Value() (driver.Value, error) {
	return c.UserID, nil
}

// Scan re-implements the database/sql Scan() method.
// It will map the string found in the response to
// to a new OmbiUserQualityProfile struct as the ID
func (c *OmbiUserQualityProfile) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		*c = OmbiUserQualityProfile{UserID: src.(string)}
	default:
		return errors.New("incompatible type for OmbiUserQualityProfile")
	}
	return nil
}

// Value re-implements the database/sql Value() method.
// It will map the Ombi userID found in the response to
// the field.
func (c OmbiUserQualityProfile) Value() (driver.Value, error) {
	return c.UserID, nil
}
