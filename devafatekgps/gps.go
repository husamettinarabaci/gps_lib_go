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
	fixQuality         string
	satellites         string
	horizontalDilution string
	antennaAltitude    string
	antennaHeight      string
	updateAge          string
}

//GNGGA,095507.00,3953.65517,N,03248.30082,E,2,07,1.47,935.3,M,36.0,M,,0000

//ParseGpsLine Parse Gps Data
func ParseGpsLine(line string) (Gps, error) {
	tokens := strings.Split(line, ",")
	if tokens[0] == "$GPGGA" || tokens[0] == "$GLGGA" || tokens[0] == "$GAGGA" || tokens[0] == "$GNGGA" {
		return Gps{
			fixTimestamp:       tokens[1],
			latitude:           tokens[2],
			latitudeDirection:  tokens[3],
			longitude:          tokens[4],
			longitudeDirection: tokens[5],
			fixQuality:         tokens[6],
			satellites:         tokens[7],
		}, nil
	}
	return Gps{}, errors.New("unsupported gps string")
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
func (gps Gps) GetLatitude() (string, error) {
	return ParseDegrees(gps.latitude, gps.latitudeDirection)
}

//GetFixTimestamp Get Fix Timestamp
func (gps Gps) GetFixTimestamp() string {
	return gps.fixTimestamp
}

//GetLongitude Get Longitude
func (gps Gps) GetLongitude() (string, error) {
	return ParseDegrees(gps.longitude, gps.longitudeDirection)
}

//GetFixQuality Get Fix Quality
func (gps Gps) GetFixQuality() string {
	return gps.fixQuality
}
