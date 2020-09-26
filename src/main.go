package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func send_message(ip string, message string) {
	conn, err := net.Dial("tcp", ip+":55443")
	defer conn.Close()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	conn.Write([]byte(message + "\r\n"))
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {

	// IP of device
	ipPtr := flag.String("ip", "", "Globe IP address (required)")

	// power flag
	powerPtr := flag.String("power", "", "Set Power state")

	// set_hsv
	hsvHuePtr := flag.Int("hsvHue", -1, "HSV Hue (optional, [0-359])")
	hsvSatPtr := flag.Int("hsvSat", -1, "HSV Saturation (optional, [0-100])")

	// set_rgb
	rgbPtr := flag.String("rgb", "", "rgb Hue (optional, [0-16777215])")

	// set_bright
	brightnessPtr := flag.Int("brightness", -1, "brightness (optional, [0-359])")

	// transition effects
	durationPtr := flag.Int("hsvDuration", 500, "Transition Effect Duration (optional, [>30, default:500])")
	effectPtr := flag.String("rgbEffect", "smooth", "Transition Effect (optional, [sudden, smooth (default)])")

	flag.Parse()

	// Check if IP is set, if not, exit
	if *ipPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// manage power
	switch *powerPtr {
	case "on":
		fmt.Println("Powering on globe at " + *ipPtr)
		message := "{\"id\": 1, \"method\": \"set_power\", \"params\":[\"on\", \"smooth\", 500]}"
		send_message(*ipPtr, message)
	case "off":
		fmt.Println("Powering off globe at " + *ipPtr)
		message := "{\"id\": 1, \"method\": \"set_power\", \"params\":[\"off\", \"smooth\", 500]}"
		send_message(*ipPtr, message)
	case "toggle":
		fmt.Println("Toggling power on globe at " + *ipPtr)
		message := "{\"id\":1,\"method\":\"toggle\",\"params\":[]}"
		send_message(*ipPtr, message)
	}

	// set hsv
	switch {
	case (*hsvHuePtr >= 0) && (*hsvSatPtr >= 0):
		fmt.Println("Setting HSV on globe at " + *ipPtr)
		message := "{\"id\":1,\"method\":\"set_hsv\",\"params\":[" + strconv.Itoa(*hsvHuePtr) + "," + strconv.Itoa(*hsvSatPtr) + ",\"" + *effectPtr + "\"," + strconv.Itoa(*durationPtr) + "]}"
		send_message(*ipPtr, message)
	}

	// set rgb
	switch {
	//case (*rgbPtr >= 0) && (*rgbPtr <= 16777215):
	case (*rgbPtr != ""):
		var rgbErr string = "Invalid RGB value, must be between 000000 and FFFFFF"
		rgbDecimal, err := strconv.ParseInt(*rgbPtr, 16, 32)
		if err != nil {
			fmt.Println(rgbErr)
			os.Exit(1)
		}

		if rgbDecimal >= 0 && rgbDecimal <= 16777215 {
			fmt.Println("Setting RGB on globe at " + *ipPtr)
			message := "{\"id\":1,\"method\":\"set_rgb\",\"params\":[" + strconv.Itoa(int(rgbDecimal)) + ",\"" + *effectPtr + "\"," + strconv.Itoa(*durationPtr) + "]}"
			send_message(*ipPtr, message)
		} else {
			fmt.Println(rgbErr)
		}
	}

	// set brightness
	switch {
	case (*brightnessPtr >= 1) && (*brightnessPtr <= 100):
		fmt.Println("Setting RGB on globe at " + *ipPtr)
		message := "{\"id\":1,\"method\":\"set_bright\",\"params\":[" + strconv.Itoa(*brightnessPtr) + ",\"" + *effectPtr + "\"," + strconv.Itoa(*durationPtr) + "]}"
		send_message(*ipPtr, message)
	}

}
