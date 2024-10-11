package utils

import (
	"log"
	"os"
	"time"

	"github.com/sony/sonyflake"
)

var flake *sonyflake.Sonyflake

func init() {
	st := sonyflake.Settings{
		StartTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		MachineID: getMachineID,
	}

	flake = sonyflake.NewSonyflake(st)
	if flake == nil {
		log.Fatalf("Sonyflake not created")
	}
}

func getMachineID() (uint16, error) {

	machineIDStr := os.Getenv("INSTANCE_ID")
	if machineIDStr == "" {
		log.Printf("INSTANCE_ID is not set")
		return 0, nil
	}

	// Convert string to uint16
	return uint16([]byte(machineIDStr)[0]), nil
}

func GenerateID() (uint64, error) {
	id, err := flake.NextID()
	if err != nil {
		log.Printf("Failed to generate ID: %v", err)
		return 0, err
	}

	//Limit ID to 53 bits
	const mask uint64 = 0x1FFFFFFFFFFFFF
	id = id & mask

	return id, nil
}
