package helper

import "testing"

/*
import (
	"testing"
	"fmt"
)	*/

func TestIsHetuValid(t *testing.T) {
	// Test valid hetu
	validHetu := "010190-123A"
	if !IsHetuValid(validHetu) {
		t.Errorf("Expected %s to be valid, but it was invalid", validHetu)
		//} else {
		//	fmt.Println("Something odd in valid testing.")
	}

	// Test invalid hetu
	invalidHetu := "1234567890"
	if IsHetuValid(invalidHetu) {
		t.Errorf("Expected %s to be invalid, but it was valid", invalidHetu)
		//} else {
		//	fmt.Println("Something odd in invalid testing.")
	}
}
