package utils

import (
	"log"
	"net"
)

func GetInternalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Get Internal Ip Fail.")
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	log.Println("Get Internal Ip Fail.")
	return ""
}