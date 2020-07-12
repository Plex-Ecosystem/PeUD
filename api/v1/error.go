package v1

import (
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type InternalServerError struct {
	Code     int
	Function string
	Error    string
}

// from https://stackoverflow.com/a/37214476/12204515
type ConvertibleBoolean bool

func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" || asString == "null" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

func (bit *ConvertibleBoolean) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		switch src.(int64) {
		case 0:
			*bit = false
		case 1:
			*bit = true
		default:
			return errors.New("incompatible type for ConvertibleBoolean")
		}
	}
	return nil
}

// Convert timestamp string to time.Time
// Needs: One way xml unmarshal
//        String convenience function
//        Value function for sqlx read
type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	n, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(n, 0))
	return nil
}

func (t *Time) String() string {
	return time.Time(*t).Format(time.RFC3339)
}

//func (t *Time) Scan(src interface{}) error {
//	switch src.(type) {
//	case time.Time:
//		*t = Time(src.(time.Time))
//	default:
//		return errors.New("incompatible type for PlexUserServerSlice")
//	}
//	return nil
//}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
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
	return intSliceToString(serverIDs, ","), nil
}

func intSliceToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}
