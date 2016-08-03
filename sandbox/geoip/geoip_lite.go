package geoip

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func GeoIPCheck() {
	fmt.Println("Hello 3")
	db, _ := geoip2.Open("/usr/local/share/GeoIP/GeoLite2-Country.mmdb")
	defer db.Close()
	result, _ := db.Country(net.ParseIP("195.175.1.1"))
	fmt.Println(string(result.Country.Names["en"])) // Turkey
}