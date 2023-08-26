package main

import (
	"crypto/rand"
)

type AddressInformation struct {
	IP            string                       `json:"ip"`
	Type          string                       `json:"type"`
	Continent     string                       `json:"continent"`
	ContinentCode string                       `json:"continent_code"`
	Country       string                       `json:"country"`
	CountryCode   string                       `json:"country_code"`
	Region        string                       `json:"region"`
	RegionCode    string                       `json:"region_code"`
	City          string                       `json:"city"`
	Latitude      float64                      `json:"latitude"`
	Longitude     float64                      `json:"longitude"`
	IsEU          bool                         `json:"is_eu"`
	PostalCode    string                       `json:"postal"`
	CallingCode   string                       `json:"calling_code"`
	Capital       string                       `json:"capital"`
	Connection    AddressConnectionInformation `json:"connection"`
	Timezone      AddressTimezoneInformation   `json:"timezone"`
}

type AddressConnectionInformation struct {
	SystemNumber    int    `json:"asn"`
	Organisation    string `json:"org"`
	ServiceProvider string `json:"isp"`
	ISPDomain       string `json:"domain"`
}

type AddressTimezoneInformation struct {
	ID           string `json:"id"`
	Abbreviation string `json:"abbr"`
	UTC          string `json:"utc"`
}

func RandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func RemoveElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func Input(msg string) string {
	PrintR(msg)
	scanner.Scan()
	return scanner.Text()
}

func RInput(msg string) string {
	PrintR(msg)
	scanner.Scan()
	return scanner.Text()
}
