package utils

import (
	"github.com/mssola/useragent"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func GetMD(md []string) string {
	if len(md) > 0 {
		return md[0]
	}
	return ""
}

func DeviceFromUserAgent(userAgentString string) string {
	ua := useragent.New(userAgentString)
	return ua.Model()
}

func OSFromUserAgent(userAgentString string) string {
	ua := useragent.New(userAgentString)
	return ua.OS()
}

func BrowserFromUserAgent(userAgentString string) string {
	ua := useragent.New(userAgentString)
	name, _ := ua.Browser()
	return name
}

func GetIPInfo(ipAddress string) (*geoip2.City, error) {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ip := net.ParseIP(ipAddress)
	record, err := db.City(ip)
	if err != nil {
		return nil, err
	}

	return record, nil
}
