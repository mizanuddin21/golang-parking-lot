package main

import (
	"bufio" //untuk membaca input line dari command input
	"fmt" //untuk console log
	"log" //untuk logging error
	"os" //untuk membaca command line dari os
	"strings" //convert string ke integer
	"strconv" //untuk edit string seperti trim, split dll
)

type Car struct {
	RegistrationNumber string
}

type ParkingLot struct {
	Capacity int
	Slots []*Car
}

/*Function untuk create slot*/
func (p *ParkingLot) Create(capacity int) {
	p.Capacity = capacity
	p.Slots = make([]*Car, capacity) //Initiate parking slots for each car number
	fmt.Printf("Created slot for parking lot with slot no. %d\n\n", capacity)
}

/*Function untuk automate park*/
func (p *ParkingLot) Park(carNumber string) {
	for i := 0; i < p.Capacity; i++ {
		if p.Slots[i] == nil {
			p.Slots[i] = &Car{RegistrationNumber: carNumber}
			fmt.Printf("Allocated slot number : %d\n", i+1)
			return
		}
	}
	fmt.Printf("Sorry, parking lot is full\n\n")
}

/*Function untuk menghitung total charge parkir sesuai jam*/
func (p *ParkingLot) leave(carNumber string, hours int) {
	for i, car := range p.Slots {
		if car != nil && car.RegistrationNumber == carNumber {
			charge := 10
			if hours > 2 {
				charge += (hours - 2) * 10
			}
			fmt.Printf("Registration number %s with Slot number %d is free with charge $%d\n\n", carNumber, i+1, charge)
			p.Slots[i] = nil
			return
		}
	}
	fmt.Printf("Registration number %s not found\n\n", carNumber)
}

/*Function untuk list nomor kendaraan mobile pada parkiran slot berapa*/
func (p *ParkingLot) Status() {
	fmt.Printf("Slot no. Registration No.\n")
	for i, car := range p.Slots {
		if car != nil {
			fmt.Printf("%d %s\n", i+1, car.RegistrationNumber)
		}
	}
}

/*Function untuk membaca command melalui file*/
func processFile(filename string, lot *ParkingLot) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		switch parts[0] {
			case "create_parking_lot":
				cap, _ := strconv.Atoi(parts[1])
				lot.Create(cap)
			case "park":
				lot.Park(parts[1])
			case "leave":
				hours, _ := strconv.Atoi(parts[2])
				lot.leave(parts[1], hours)
			case "status":
				lot.Status()
			default:
				fmt.Printf("Command not found: ", line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*Main function*/
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide input file as argument")
	}

	lot := &ParkingLot{}
	processFile(os.Args[1], lot)
}