package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
)

// TCP client connectivity
func StartClient(host string, port int, username string, password string) {
	reader := bufio.NewReader(os.Stdin)

	// Connect to TallyDB server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		fmt.Println("❌ Error connecting:", err)
		os.Exit(1)
	}

	// Send handshake
	handshakeRequest := Request{RequestId: "1", Credentials: CredentialsType{Username: username, Password: password}}
	handshakeMessage, _ := json.Marshal(handshakeRequest)
	conn.Write([]byte(handshakeMessage))

	fmt.Println(`🔥 Connected to TallyDB server`)
	fmt.Println(`🔑 Authentication successful`)
	fmt.Println()

	for true {
		fmt.Print(color.HiCyanString("tally> "))
		command, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(command))

		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println("Error receiving response from server:", err)
			os.Exit(1)
		}

		fmt.Println("Server response:", string(response))
	}

	defer conn.Close()
}
