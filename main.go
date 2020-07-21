package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No parameter supplied")
	}

	reg, err := strconv.ParseInt(os.Args[1], 0, 32)
	if err != nil {
		log.Fatal(err)
	}

	printPinctrl(int(reg))
}

func printPinctrl(n int) {
	// SION
	if n&(1<<30) != 0 {
		fmt.Print("SION | ")
	}

	// Pull
	if n&(1<<8) != 0 {
		if n&(1<<6) == 0 {
			fmt.Printf("PULL_DOWN")
		} else {
			fmt.Printf("PULL_UP")
		}
	} else {
		fmt.Printf("(PULL_NONE)")
	}

	fmt.Printf(" | ")

	// Hyst
	if (1<<7)&n != 0 {
		fmt.Print("HYSTERESIS | ")
	}

	// ODE
	if (1<<5)&n != 0 {
		fmt.Printf("OPEN DRAIN")
	} else {
		fmt.Printf("(PUSH PULL)")
	}

	fmt.Printf(" | ")

	// Slew
	switch (n >> 4) & 0x1 {
	case 0:
		fmt.Printf("SLEW SLOW")
	case 1:
		fmt.Printf("SLEW FAST")
	}

	fmt.Printf(" | ")
	// DSE
	switch (n & 0x6) >> 1 {
	case 0:
		fmt.Printf("DSE X1")
	case 1:
		fmt.Printf("DSE X4")
	case 2:
		fmt.Printf("DSE X2")
	case 3:
		fmt.Printf("DSE X6")
	}
}
