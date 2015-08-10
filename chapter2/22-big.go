package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792) // km/s
	secondsPerDay := big.NewInt(86400)

	km := new(big.Int)
	km.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", km, "km away.") //#A

	seconds := new(big.Int)
	seconds.Div(km, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed.") //#B
}

// #A Andromeda Galaxy is 24000000000000000000 km away.
// #B That is 926568346 days of travel at light speed.