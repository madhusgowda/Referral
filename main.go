package main

import (
	"math/rand"
	"random/randomNumber"
	"random/serialNumber"
	"random/urlGenerate"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber.RandomString(6)
	urlGenerate.Link()
	serialNumber.SerialNumber()
}
