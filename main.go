package main

import (
	"fmt"
	"strings"
)

var adminFees = map[string]float32{
	"senin":  0.10,
	"selasa": 0.05,
	"rabu":   0.10,
	"kamis":  0.05,
	"jumat":  0.10,
	"sabtu":  0.05,
}

// Location
var deliverySchedule = map[string][]string{
	"senin":  {"JKT", "DPK"},
	"selasa": {"JKT", "BKS", "DPK"},
	"rabu":   {"JKT", "BDG"},
	"kamis":  {"JKT", "BDG", "BKS"},
	"jumat":  {"JKT", "BKS"},
	"sabtu":  {"JKT", "BDG"},
}

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)
	validLocations := deliverySchedule[day]

	for _, order := range data {
		parts := strings.Split(order, ":")
		if len(parts) != 4 {
			continue
		}

		firstName := parts[0]
		lastName := parts[1]
		price := parts[2]
		location := parts[3]

		priceValue := parsePrice(price)

		if isValidLocation(location, validLocations) {
			totalCost := priceValue + (priceValue * adminFees[day])
			fullName := firstName + "-" + lastName
			result[fullName] = totalCost
		}
	}

	return result
}

func isValidLocation(location string, validLocations []string) bool {
	for _, validLocation := range validLocations {
		if location == validLocation {
			return true
		}
	}
	return false
}

func parsePrice(price string) float32 {
	var priceValue float32
	fmt.Sscanf(price, "%f", &priceValue)
	return priceValue
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
