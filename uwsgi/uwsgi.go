package uwsgi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	types "github.com/micahhausler/consul-uwsgi-healthcheck/types"
	"net"
	"strconv"
)

type UwsgiPacketHeader struct {
	Modifier1 uint8
	Datasize  uint16
	Modifier2 uint8
}

func createPing() *UwsgiPacketHeader {
	return &UwsgiPacketHeader{100, 0, 0}
}

// This function sends a uwsgi ping and expects a pong back.
// If no pong is returned, the function returns `false`
func Ping(config types.Config) bool {
	ping := createPing()
	buffer := new(bytes.Buffer)

	if err := binary.Write(buffer, binary.LittleEndian, ping); err != nil {
		return false
	}

	conn, err := net.Dial("tcp", config.Address+":"+strconv.Itoa(config.Port))
	if err != nil {
		if config.Verbose == true {
			fmt.Printf("Could not connect to %s:%d\n", config.Address, config.Port)
		}
		return false
	}

	// Send Ping
	if _, err = conn.Write(buffer.Bytes()); err != nil {
		if config.Verbose == true {
			fmt.Printf("Could not send bytes to %s:%d\n", config.Address, config.Port)
		}
		return false
	}

	input := make([]byte, 4)

	// Read Pong
	if _, err = conn.Read(input); err != nil {
		if config.Verbose == true {
			fmt.Printf("Error reading bytes from %s:%d\n", config.Address, config.Port)
		}
		return false
	}

	buffer = bytes.NewBuffer(input)
	pong := new(UwsgiPacketHeader)
	if err = binary.Read(buffer, binary.LittleEndian, pong); err != nil {
		if config.Verbose == true {
			fmt.Printf("Error parsing bytes from %s:%d\n", config.Address, config.Port)
		}
		return false
	}

	if pong.Modifier2 != uint8(1) {
		if config.Verbose == true {
			fmt.Printf("Invalid Modifier2 on response! Expected '1', got '%d'\n", &pong.Modifier2)
		}
		return false
	}

	return true
}
