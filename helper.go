package helper

import (
	"fmt"     // to print
	"regexp"  // to print sertain filtered string
	"strconv" // to convert string and get date
)

// Function to check if hetu is valid: ddmmyyxnnnt
func IsHetuValid(hetu string) bool {
	hetuRegex := regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])(0[1-9]|1[0-2])[0-9]{2}[-+A][0-9]{3}[0-9A-Z]$`)
	// Regular expression for hetu
	return hetuRegex.MatchString(hetu)
} // IsHetuValid()

// Function to determine gender from hetu
func GetGender(hetu string) string {
	genderChar := hetu[9] // last n of nnn
	if int(genderChar)%2 == 0 {
		return "female." // if even, female
	}
	return "male." // if odd, male
} // getGender()

// Function to print birthday from hetu
func GetBirthDate(hetu string) string {
	day, _ := strconv.Atoi(hetu[0:2])   // dd
	month, _ := strconv.Atoi(hetu[2:4]) // mm
	year, _ := strconv.Atoi(hetu[4:6])  //yy

	// Adjust year if necessary
	separator := hetu[6]
	switch separator {
	case '-':
		// If x is '-', assume the century is 1900
		year += 1900
	case '+':
		// If x is '+', assume the century is 1800
		year += 1800
	case 'A':
		// If x is 'A', assume the century is 2000
		year += 2000
	default:
		// Handle other cases if necessary
	}
	// return birthday: dd.mm.yyyy
	return fmt.Sprintf("%02d.%02d.%d", day, month, year)
} // GetBirthDate()
