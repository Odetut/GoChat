package networking

import (
	"fmt"
	"net"
	"time"
	"GoChat/main/src/modules/utilities"
	"encoding/json"
)






// Init
const MaximumWaitTime = 15 // seconds
const MaxRetries = 3
const Address = "255.255.255.255:8080" // Temporary port


func searchForServer() bool {
	var TargetServerInfo *net.UDPAddr
	var BytesRead int
	RemoteAddress, err := net.ResolveUDPAddr("udp", Address)
	if err != nil {
		panic(err)
	}

	fmt.Println("Looking for server...")
	Connection, err := net.DialUDP("udp", nil, RemoteAddress)
	buffer := make([]byte, 1024)

	Connection.Write([]byte("GOCHAT"))
	Connection.SetReadDeadline(time.Now().Add(MaximumWaitTime * time.Second))
	BytesRead, TargetServerInfo, err = Connection.ReadFromUDP(buffer)

	if err != nil {
		Connection.Close()

		for Attempt := 1; Attempt != MaxRetries; Attempt++ {
			Connection, err := net.DialUDP("udp", nil, RemoteAddress)
			Connection.SetReadDeadline(time.Now().Add(MaximumWaitTime * time.Second))
			Connection.Write([]byte("GOCHAT"))
			BytesRead, TargetServerInfo, err = Connection.ReadFromUDP(buffer)
				

			if err != nil {
				//Error Handle

			} else {
				Connection.Close()
				break
			}

		}
	}

	Response := buffer[:BytesRead] // Check Server info to see compatibility
	var ResponseData utilities.AppData

	if err := json.Unmarshal(Response,&ResponseData) ; err != nil{
		//Do stuff aka We failed now panic and say server is incompatible
	}

	
	
	

	ServerAddr := TargetServerInfo.String()


	// Attempt to connect via TCP if server is compatible 

	return true
}

// func Disconnect() {
// 	fmt.Println("Disconnecting from server...")
// 	Connection.Close()
// }

// Starts looking for a server. Returns false if no server
func StartClient() bool {
	if !searchForServer() {
		return false
	}

	return true
}
