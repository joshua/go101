package main

import "fmt"

// INT_START OMIT
type Communicate interface {
	Call(phoneNumber string)
	Message(phoneNumber string)
}

// INT_END OMIT

// EMBD_START OMIT
type OS struct {
	Type    string
	Version string
}

type Hardware struct {
	Memory int
	Ram    int
	NFC    bool
}

type Phone struct {
	Model string
	Hardware
	OS
}

// EMBD_END OMIT

// INT_IMP_START OMIT
func (p *Phone) Call(phoneNumber string) {
	fmt.Printf("Calling %s", phoneNumber)
}

func (p *Phone) Message(phoneNumber string) {
	fmt.Printf("Messaging %s", phoneNumber)
}

// INT_IMP_END OMIT

// CHOOSE_START OMIT
func chooseCommunication(c Communicate) {
	var i int
	fmt.Print("Choose Communication\n 1. Call\n 2. Message\n\n")
	// _, err := fmt.Scan(&i)
	i = 1

	switch i {
	case 1: // Call
		c.Call("843-123-4567")
	case 2: // Message
		c.Message("843-123-4567")
	default:
		fmt.Println("Failed to choose correctly")
	}
}

// CHOOSE_END OMIT
// MAIN_START OMIT
func main() {
	iphone6 := &Phone{
		Model:    "iPhone 6",
		Hardware: Hardware{Memory: 16, Ram: 1, NFC: false},
		OS:       OS{Type: "iOS", Version: "9.0.0"}}

	chooseCommunication(iphone6)
}

// MAIN_END OMIT
