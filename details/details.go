package details

import (
	"log"
	"net"
	"os"
)

func GetHostName() (name string, err error) {
	hostname, err := os.Hostname()
	return hostname, err
}

func GetOutboundIP() (net.IP, error){
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, err
}