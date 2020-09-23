package main

import (
    "fmt"
    "flag"
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
    ipPtr :=  flag.String("ip", "", "Globe IP address (required)")

    // power flag
    powerPtr := flag.String("power", "", "Set Power state")

    // set_hsv
    hsvHuePtr := flag.Int("hsvHue", -1, "HSV Hue (optional, [0-359])")
    hsvSatPtr := flag.Int("hsvSat", -1, "HSV Sat (optional, [0-100])")
    hsvEffectPtr := flag.String("hsvEffect","smooth", "HSV Effect (optional, [sudden, smooth (default)])")
    hsvDurationPtr := flag.Int("hsvDuration",500, "HSV Duration (optional, [>30, default:500])")

    // set_rgb
    rgbPtr := flag.Int("rgb", -1, "rgb Hue (optional, [0-16777215])")
    rgbEffectPtr := flag.String("rgbEffect","smooth", "rgb Effect (optional, [sudden, smooth (default)])")
    rgbDurationPtr := flag.Int("rgbDuration",500, "rgb Duration (optional, [>30, default:500])")

    // set_bright
    brightnessPtr := flag.Int("brightness", -1, "brightness Hue (optional, [0-359])")
    brightnessEffectPtr := flag.String("brightnessEffect","smooth", "brightness Effect (optional, [sudden, smooth (default)])")
    brightnessDurationPtr := flag.Int("brightnessDuration",500, "brightness Duration (optional, [>30, default:500])")

    
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
            message := "{\"id\":1,\"method\":\"set_hsv\",\"params\":[" + strconv.Itoa(*hsvHuePtr) +","+ strconv.Itoa(*hsvSatPtr) +",\""+ *hsvEffectPtr +"\","+ strconv.Itoa(*hsvDurationPtr) + "]}"
            send_message(*ipPtr, message)
    }

    // set rgb
    switch {
        case (*rgbPtr >= 0) && (*rgbPtr <= 16777215):
            fmt.Println("Setting RGB on globe at " + *ipPtr)
            message := "{\"id\":1,\"method\":\"set_rgb\",\"params\":[" + strconv.Itoa(*rgbPtr) +",\""+ *rgbEffectPtr +"\","+ strconv.Itoa(*rgbDurationPtr) + "]}"
            send_message(*ipPtr, message)
    }

    // set brightness
    switch {
        case (*brightnessPtr >= 1) && (*brightnessPtr <= 100):
            fmt.Println("Setting RGB on globe at " + *ipPtr)
            message := "{\"id\":1,\"method\":\"set_bright\",\"params\":[" + strconv.Itoa(*brightnessPtr) +",\""+ *brightnessEffectPtr +"\","+ strconv.Itoa(*brightnessDurationPtr) + "]}"
            send_message(*ipPtr, message)
    }

}
