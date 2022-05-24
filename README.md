# gps_lib_go

Go dili için gps kütüphanesi

## Kurulum:

```shell
go get "github.com/HsmTeknoloji/gps_lib_go/devhsmtekgps"
```

## Kullanım:

```go
gps, err := devhsmtekgps.ParseGpsLine(scanner.Text())
if err == nil {
	if gps.GetFixQuality() == "1" || gps.GetFixQuality() == "2" {
		latitude, _ := gps.GetLatitude()
		longitude, _ := gps.GetLongitude()
	} else {
		fmt.Println("gps verisi bulunamadı")
	}
}
```

## Geliştirici Bilgileri:
<img src="https://github.com/HsmTeknoloji/companyfiles/blob/master/hsmtek-logo.png?raw=true" width="200"/>
Web Site        : www.hsmteknoloji.com <br />
Developer Groups : https://t.me/HsmTeknoloji/ <br />
