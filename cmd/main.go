package main

import (
	"bufio"
	"fmt"
	"os"

	"parking-app/internal/parser"
	"parking-app/internal/parking"
)

func main() {
	var lot *parking.ParkingLot

	if len(os.Args) > 1 {
		// 🔹 Mode baca dari file
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			output := parser.ProcessCommand(line, &lot)
			if output != "" {
				fmt.Println(output)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	} else {
		// 🔹 Mode interaktif (stdin)
		fmt.Println("Parking Lot CLI (type 'exit' to quit)")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			if line == "exit" {
				break
			}
			output := parser.ProcessCommand(line, &lot)
			if output != "" {
				fmt.Println(output)
			}
		}
	}
}
