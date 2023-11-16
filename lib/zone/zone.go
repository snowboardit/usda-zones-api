package zone

import (
	"fmt"

	"github.com/snowboardit/usda-zones-api/lib/data"
)

type Row = data.Row

func GetZoneByZip(zip string, data []Row) (Row, error) {
	formattedZip := fmt.Sprintf("%05v", zip)[:5]
	for _, r := range data {
		if r.Zipcode == formattedZip {
			fmt.Printf("Found zip: %s\nData: %s", formattedZip, r)
			return r, nil
		}
	}
	return Row{}, fmt.Errorf("Unable to find data for zip: %s", zip)
}
