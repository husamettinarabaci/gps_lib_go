# ping_lib_go

Gps library for Go

## Installation:

```shell
go get "github.com/AfatekDevelopers/gps_lib_go/devafatekgps"
```

## Usage:

```go
gps, err := devafatekgps.ParseGpsLine(scanner.Text())
if err == nil {
	if gps.GetFixQuality() == "1" || gps.GetFixQuality() == "2" {
		latitude, _ := gps.GetLatitude()
		longitude, _ := gps.GetLongitude()
	} else {
		fmt.Println("no gps fix available")
	}
}
```

## Developers:
<img src="https://github.com/AfatekDevelopers/companyfiles/blob/master/afatek-logo.png?raw=true" width="200"/>
Web Site        : www.afatek.com.tr <br />
Developer Groups : https://t.me/Afatek/ <br />