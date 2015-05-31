package uwsgi

import "fmt"
import "testing"

func TestCreatePing(t *testing.T) {
	ping := createPing()
	if ping.Modifier1 != 100 {
		fmt.Printf("Expected Modifier1 of 100, got %s\n", ping.Modifier1)
		t.Fail()
	}
	if ping.Datasize != 0 {
		fmt.Printf("Expected Datasize of 0, got %s\n", ping.Datasize)
		t.Fail()
	}
	if ping.Modifier2 != 0 {
		fmt.Printf("Expected Modifier2 of 0, got %s\n", ping.Modifier2)
		t.Fail()
	}

}

func TestPing(t *testing.T) {

}
