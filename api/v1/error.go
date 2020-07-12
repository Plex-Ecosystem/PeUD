package v1

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"

	"github.com/DirtyCajunRice/go-utility"
)

type InternalServerError struct {
	Code     int
	Function string
	Error    string
}

// Create plex user slice type for database functions
// Needs: Scan function for sql scan
//        Value function for sqlx read
type PlexUserServerSlice []PlexUserServer

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

func (p PlexUserServerSlice) Value() (driver.Value, error) {
	serverIDs := make([]int, 0)
	for _, server := range p {
		serverIDs = append(serverIDs, server.ID)
	}
	return utility.IntSliceToString(serverIDs, ","), nil
}
