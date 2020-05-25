# gps_lib_go

Go dili için gps kütüphanesi

## Kurulum:

```shell
go get "github.com/AfatekDevelopers/gps_lib_go/devafatekgps"
```

## Kullanım:

```go
gps, err := devafatekgps.ParseGpsLine(scanner.Text())
if err == nil {
	if gps.fixQuality == "1" || gps.fixQuality == "2" {
		latitude, _ := gps.GetLatitude()
		longitude, _ := gps.GetLongitude()
	} else {
		fmt.Println("gps verisi bulunamadı")
	}
}
```

## Geliştirici Bilgileri:
<img src="https://github.com/AfatekDevelopers/companyfiles/blob/master/afatek-logo.png?raw=true" width="200"/>
Web Site        : www.afatek.com.tr <br />
Developer Groups : https://t.me/Afatek/ <br />
