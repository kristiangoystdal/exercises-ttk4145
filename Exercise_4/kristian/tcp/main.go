package main

import (
	"Exercise_4/pack"
	"fmt"
)

var number int

func becomeMaster() {
	fmt.Printf("Assuming master role with initial number: %d\n", number)
	pack.Broadcast(number)
}

func main1() {
	number = 0

	// The ListenForMaster function is expected to be blocking and return when it's time to switch roles.
	// It also returns the last number received, which should be used if this instance becomes the master.
	slave, receivedNumber := pack.ListenForMaster()

	if slave {
		fmt.Println("Operating as slave, received last number:", receivedNumber)
		// This instance continues as a slave unless ListenForMaster indicated it's time to switch roles.
		// The logic to switch roles would be inside ListenForMaster based on timeout/no messages received.
	} else {
		// If we're here, it means no master was detected or the master stopped broadcasting.
		number = receivedNumber + 1 // Increment to continue from the last received number
		becomeMaster()
	}
}
