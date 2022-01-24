package main

import (
	"flag"
	"fmt"
	"log"
	"main/asset"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

func init() {
	flag.Parse()
}

func main() {
	mmdb, _ := asset.Asset("Country.mmdb")
	db, err := geoip2.FromBytes(mmdb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(os.Args[len(os.Args)-1])
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", ip, record.Country.Names["en"])
}
