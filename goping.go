package main

import (
        "fmt"
        "net"
        "os"
        "runtime"
        "strconv"
        "strings"
        "time"
//	"reflect"
)

var (
        // Colors
        red    = "\x1b[31m"
        green  = "\x1b[32m"
        yellow = "\x1b[33m"
        blue   = "\x1b[34m"
        reset  = "\x1b[0m"
)

func main() {
        // Get the file name
        filenameSplit := strings.Split(os.Args[0], "/")
        filename := filenameSplit[len(filenameSplit)-1]

        // Make sure the user has the correct number of arguments
        if len(os.Args) < 4 {
                fmt.Printf("Usage: ./%v <ip> <port> <count> <delay ex: 5s, 5m, 2s, 2h> <threshold> \n", filename)
                return
        }

        // Get OS args
        var timeout time.Duration
        ip := os.Args[1]
        port := os.Args[2]
        count, _ := strconv.Atoi(os.Args[3])
        threshold, _ := strconv.Atoi(os.Args[5])

        // Make timeout optional
        if len(os.Args) == 6 {
                timeout, _ = time.ParseDuration(os.Args[4])
        }

        // Connect to the host for count times
        for i := 0; ; i++ {
                // If count is -1, loop forever
                // Otherwise, break when count is reached
                if count != -1 && i >= count {
                        break
                }

                beforeTime := time.Now()                                         // Get current time before connection
                conn, err := net.Dial("tcp", ip+":"+port)                        // Connect to host
                roundedEndTime := time.Since(beforeTime).Round(time.Millisecond) // Round time since the connection was initiated to milliseconds
                conn.Close()                                                     // Close connection

                // If GOOS is windows, do not use colors
                if runtime.GOOS == "windows" {
                        red = ""
                        green = ""
                        yellow = ""
                        blue = ""
                        reset = ""
                }

                if err == nil {
                        fmt.Printf("%vConnected to %v:%v time=%v\n%v", green, ip, port, roundedEndTime, reset)
                        //fmt.Println(threshold)
			//fmt.Println(roundedEndTime)
			suf := roundedEndTime.String()
			trimed := suf[:len(suf)-2]
                        //fmt.Println(trimed)
 			intValue, _  := strconv.Atoi(trimed)
			if intValue > threshold {
                            fmt.Println("Over Treshold Alert Raised")
			    f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			    if err != nil {
				    panic(err)
			    }
			    defer f.Close()
			    currentTime := time.Now()
			    formatedTime := currentTime.Format("2 Jan 06 03:04PM")
			    //fmt.Println(reflect.TypeOf(roundedEndTime))
			    strList := []string{formatedTime, trimed, ip, port, "\n"}
			    fileLogs := strings.Join(strList, "\t")
			    if _, err = f.WriteString(fileLogs); err != nil {
				    panic(err)
				}
			}
		} else {
                        fmt.Printf("%vFailed to connect to %v:%v time=%v\n%v", red, ip, port, roundedEndTime, reset)
                }

                time.Sleep(timeout | time.Second) // Sleep for the specified timeout or 1 second
        }
}

