package codewars

import (
	"strconv"
	"strings"
)

type clientaddresses struct {
	street string
	house  int
}

func SalesmanTravel(addresses, zipcode string) string {
	addressesMap := map[string][]clientaddresses{}
	for _, address := range strings.Split(addresses, ",") {
		address = strings.TrimSpace(address)
		address_chunks := strings.Split(address, " ")
		zip := address_chunks[len(address_chunks)-2] + " " + address_chunks[len(address_chunks)-1]
		house, _ := strconv.Atoi(address_chunks[0])
		street := strings.Join(address_chunks[1:len(address_chunks)-2], " ")
		addressesMap[zip] = append(addressesMap[zip], clientaddresses{
			house:  house,
			street: street,
		})
	}
	items, ok := addressesMap[zipcode]
	if ok {
		houses, streets := []string{}, []string{}
		for _, item := range items {
			house := strconv.Itoa(item.house)
			houses = append(houses, house)
			streets = append(streets, item.street)
		}
		return zipcode + ":" + strings.Join(streets, ",") + "/" + strings.Join(houses, ",")
	} else {
		return zipcode + ":/"
	}
}
