package telnet

import (
	"net"
	"time"
)

func TcpGather(ip string, ports []string) map[string]string {
	// 检查 emqx 1883, 8083, 8080, 18083 端口

	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[port] = "failed"
			// todo log handler
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}
