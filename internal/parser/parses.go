package parser

import (
	"fmt"
	"strconv"
	"strings"

	"parking-app/internal/parking"
)

// ProcessCommand menjalankan command satu baris
func ProcessCommand(line string, lot **parking.ParkingLot) string {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return ""
	}

	switch parts[0] {
	case "create_parking_lot":
		if len(parts) < 2 {
			return "Invalid command"
		}
		capacity, _ := strconv.Atoi(parts[1])
		*lot = parking.NewParkingLot(capacity)
		return fmt.Sprintf("Created a parking lot with %d slots", capacity)

	case "park":
		if *lot == nil {
			return "Parking lot not created"
		}
		carNumber := parts[1]
		car := &parking.Car{Number: carNumber}
		slot, err := (*lot).Park(car)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Allocated slot number: %d", slot)

	case "leave":
		if *lot == nil {
			return "Parking lot not created"
		}
		if len(parts) < 3 {
			return "Invalid command"
		}
		carNumber := parts[1]
		hours, _ := strconv.Atoi(parts[2])
		slot, charge, err := (*lot).Leave(carNumber, hours)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d",
			carNumber, slot, charge)

	case "status":
		if *lot == nil {
			return "Parking lot not created"
		}
		status := (*lot).Status()
		result := "Slot No. Registration No.\n"
		for i := 1; i <= (*lot).Capacity(); i++ {
			if num, ok := status[i]; ok {
				result += fmt.Sprintf("%d %s\n", i, num)
			}
		}
		return strings.TrimSpace(result)

	default:
		return "Unknown command"
	}
}
