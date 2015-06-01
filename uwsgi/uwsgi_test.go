package uwsgi

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCreatePing(t *testing.T) {
	ping := createPing()
	if ping.Modifier1 != 100 {
		fmt.Printf("Expected Modifier1 of 100, got %d\n", ping.Modifier1)
		t.Fail()
	}
	if ping.Datasize != 0 {
		fmt.Printf("Expected Datasize of 0, got %d\n", ping.Datasize)
		t.Fail()
	}
	if ping.Modifier2 != 0 {
		fmt.Printf("Expected Modifier2 of 0, got %d\n", ping.Modifier2)
		t.Fail()
	}

}

func TestToBytes(t *testing.T) {
	ping := createPing()
	result := ping.ToBytes()

	pong, err := ping.ToHeader(result)

	if err != nil {
		t.Fail()
	}

	if pong.Modifier1 != 100 {
		t.Fail()
	}
	if pong.Modifier2 != 0 {
		t.Fail()
	}
	if pong.Datasize != 0 {
		t.Fail()
	}
}

func TestToHeaderWithErr(t *testing.T) {
	result := []byte{7}
	ping := createPing()

	header, err := ping.ToHeader(result)
	if err == nil {
		t.Fail()
	}
	if header != nil {
		t.Fail()
	}
}

func TestIsPong(t *testing.T) {
	pong := createPong()

	if isPong := pong.IsPong(); isPong == false {
		t.Fail()
	}
	if isPong := pong.IsPing(); isPong == true {
		t.Fail()
	}
}

func TestIsPing(t *testing.T) {
	ping := createPing()

	if isPing := ping.IsPing(); isPing == false {
		t.Fail()
	}
	if isPing := ping.IsPong(); isPing == true {
		t.Fail()
	}
}

func TestWrite(t *testing.T) {
	ping := createPing()

	buffer := bytes.NewBuffer([]byte{})

	_, err := ping.Write(buffer)
	if err != nil {
		t.Fail()
	}

	response, err := ping.Read(buffer)
	if response.Modifier1 != ping.Modifier1 {
		t.Fail()
	}
	if response.Modifier2 != ping.Modifier2 {
		t.Fail()
	}
	if response.Datasize != ping.Datasize {
		t.Fail()
	}
}

func TestRead(t *testing.T) {
	ping := createPing()

	buffer := bytes.NewBuffer([]byte{})

	ping.Write(buffer)

	response, err := ping.Read(buffer)
	if err != nil {
		t.Fail()
	}
	if response.Modifier1 != ping.Modifier1 {
		t.Fail()
	}
	if response.Modifier2 != ping.Modifier2 {
		t.Fail()
	}
	if response.Datasize != ping.Datasize {
		t.Fail()
	}
}

func TestReadFail(t *testing.T) {
	ping := createPing()

	buffer := bytes.NewBuffer([]byte{})

	_, err := ping.Read(buffer)
	if err == nil {
		t.Fail()
	}
}
