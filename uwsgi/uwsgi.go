package uwsgi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	types "github.com/micahhausler/consul-uwsgi-healthcheck/types"
	"io"
	"net"
	"strconv"
)

// Create a Ping header
func createPing() *UwsgiPacketHeader {
	return &UwsgiPacketHeader{100, 0, 0}
}

// Create a Pong header
func createPong() *UwsgiPacketHeader {
	return &UwsgiPacketHeader{100, 0, 1}
}

type Header interface {
	ToBytes() []byte
	Write() (int, error)
}

type UwsgiPacketHeader struct {
	Modifier1 uint8
	Datasize  uint16
	Modifier2 uint8
}

// Return a byte slice representation of the header
func (header UwsgiPacketHeader) ToBytes() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, header)
	return buffer.Bytes()
}

// Return a new header with the contents of the byte slice
func (header UwsgiPacketHeader) ToHeader(input []byte) (*UwsgiPacketHeader, error) {
	buffer := bytes.NewBuffer(input)
	pong := new(UwsgiPacketHeader)
	if err := binary.Read(buffer, binary.LittleEndian, pong); err != nil {
		return nil, err
	}
	return pong, nil
}

// Returns a boolean on inspecting whether Modifier2 is that of a Pong
func (header UwsgiPacketHeader) IsPong() bool {
	return header.Modifier2 == uint8(1) && header.Modifier1 == uint8(100)
}

// Returns a boolean on inspecting whether Modifier2 is that of a Ping
func (header UwsgiPacketHeader) IsPing() bool {
	return header.Modifier2 == uint8(0) && header.Modifier1 == uint8(100)
}

func (header UwsgiPacketHeader) Write(writer io.Writer) (int, error) {
	return writer.Write(header.ToBytes())
}

// Returns a new header from the reader
func (header UwsgiPacketHeader) Read(reader io.Reader) (*UwsgiPacketHeader, error) {
	input := make([]byte, 4)
	if _, err := reader.Read(input); err != nil {
		return nil, err
	}
	return header.ToHeader(input)
}

func (header UwsgiPacketHeader) getConnection(config types.Config) (net.Conn, error) {
	return net.Dial("tcp", config.Address+":"+strconv.Itoa(config.Port))
}

func handleErr(err error, message string) bool {
	if err != nil {
		if message != "" {
			fmt.Println(message)
		}
		return false
	}
	return true
}

// This function sends a uwsgi ping and expects a pong back.
// If no pong is returned, the function returns `false`
func Ping(config types.Config) bool {
	ping := createPing()

	conn, err := ping.getConnection(config)
	if result := handleErr(err, fmt.Sprintf("Could not connect to %s:%d\n", config.Address, config.Port)); result == false {
		return false
	}

	// Send Ping
	_, err = ping.Write(conn)
	if result := handleErr(err, fmt.Sprintf("Could not send data to %s:%d\n", config.Address, config.Port)); result == false {
		return false
	}

	// Read Pong
	pong, err := ping.Read(conn)
	if result := handleErr(err, fmt.Sprintf("Could not read data from %s:%d\n", config.Address, config.Port)); result == false {
		return false
	}

	// Assert pong is a pong
	return pong.IsPong()
}
