```
wget https://github.com/Dreamacro/maxmind-geoip/releases/latest/download/Country.mmdb

go run ip2country.go `curl http://ifconfig.io`

go get -u github.com/jteeuwen/go-bindata/...

go-bindata -o=asset/asset.go -pkg=asset Country.mmdb

```
