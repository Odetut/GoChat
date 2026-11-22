package networking

import (
	"fmt"
	"net"
	
	
	
)

// Init
const MaximumWaitTime = 15 // seconds
const Address = "255.255.255.255:8080"// Temporary port



func searchForServer() bool {
	fmt.Println("Looking for server...")
	Connection, Err := net.Dial("udp", Address)

	if Err != nil {
		fmt.Println("No server found.")
		return false
	}

	Connection.Write([]byte("GOCHAT"))
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

