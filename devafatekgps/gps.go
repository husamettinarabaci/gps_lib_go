package devafatekgps

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Gps type
type Gps struct {
	fixTimestamp       string
	latitude           string
	latitudeDirection  string
	longitude          string
	longitudeDirection string
	FixQuality         string
	satellites         string
	horizontalDilution string
	antennaAltitude    string
	antennaHeight      string
	updateAge          string
}

//ParseGpsLine Parse Gps Data
func ParseGpsLine(line string) (Gps, error) {
	tokens := strings.Split(line, ",")
	if tokens[0] == "$GPGGA" {
		return Gps{
			fixTimestamp:       tokens[1],
			latitude:           tokens[2],
			latitudeDirection:  tokens[3],
			longitude:          tokens[4],
			longitudeDirection: tokens[5],
			FixQuality:         tokens[6],
			satellites:         tokens[7],
		}, nil
	}
	return Gps{}, errors.New("unsupported nmea string")
}

//ParseDegrees Parse Degrees
func ParseDegrees(value string, direction string) (string, error) {
	if value == "" || direction == "" {
		return "", errors.New("the location and / or direction value does not exist")
	}
	lat, _ := strconv.ParseFloat(value, 64)
	degrees := math.Floor(lat / 100)
	minutes := ((lat / 100) - math.Floor(lat/100)) * 100 / 60
	decimal := degrees + minutes
	if direction == "W" || direction == "S" {
		decimal *= -1
	}
	return fmt.Sprintf("%.6f", decimal), nil
}

//GetLatitude Get Latitude
func (nmea Gps) GetLatitude() (string, error) {
	return ParseDegrees(nmea.latitude, nmea.latitudeDirection)
}

//GetFixTimestamp Get Fix Timestamp
func (nmea Gps) GetFixTimestamp() string {
	return nmea.fixTimestamp
}

//GetLongitude Get Longitude
func (nmea Gps) GetLongitude() (string, error) {
	return ParseDegrees(nmea.longitude, nmea.longitudeDirection)
}
