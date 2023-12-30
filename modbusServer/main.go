package main

import (
	"fmt"
	"log"
	"time"

	"github.com/simonvetter/modbus"
	"github.com/tbrandon/mbserver"
)

func main() {
	// Start the server.
	serv := mbserver.NewServer()
	err := serv.ListenTCP("127.0.0.1:1502")
	fmt.Println("Serving")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	defer serv.Close()
	serv.Coils[1] = 1
	serv.HoldingRegisters[5] = 200
	// Wait for the server to start
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Creating client")
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://localhost:1502",
		Timeout: 1 * time.Second,
	})
	if err != nil {
		// error out if client creation failed
		fmt.Println("Failed to connect")

	}
	fmt.Println("Connecting Client")
	err = client.Open()
	defer client.Close()
	if err != nil {
		// error out if we failed to connect/open the device
		// note: multiple Open() attempts can be made on the same client until
		// the connection succeeds (i.e. err == nil), calling the constructor again
		// is unnecessary.
		// likewise, a client can be opened and closed as many times as needed.
		fmt.Println("failed to open.")
	}

	// read a single 16-bit holding register at address 100
	var reg16 uint16
	reg16, _ = client.ReadRegister(5, modbus.HOLDING_REGISTER)

	fmt.Println(reg16)

	client.WriteCoil(3, true)
	fmt.Println(serv.Coils[3])
	coil3, _ := client.ReadCoils(0, 5)
	fmt.Println(coil3)
}
