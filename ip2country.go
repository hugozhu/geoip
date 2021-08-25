package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"os"
	"flag"
)

var mmdbPath = flag.String("mmdb","./Country.mmdb", "Path to geoip mmdb file")

func init() {
	flag.Parse()
}

func main() {
	db, err := geoip2.Open(*mmdbPath)
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