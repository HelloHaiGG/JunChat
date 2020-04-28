package utils

import (
	"log"
	"net"
	"time"
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

func Telnet(ip, port string) bool {
	addr := net.JoinHostPort(ip, port)
	//telnet
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return false
	}
	if conn != nil {
		_ = conn.Close()
		return true
	}
	return false
}
