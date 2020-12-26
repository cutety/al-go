package loadbalancer

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadBalancer(t *testing.T) {
	client := &Client{}
	client.Add("192.168.1.1", 1)
	client.Add("192.168.1.2", 3)
	client.Add("192.168.1.3", 1)
	for i := 0; i < 10; i++ {
		server, err := client.LoadBalancer()
		if err != nil {
			log.Println("error", err)
		}
		fmt.Println(server.IP)
	}
}

// output
// 192.168.1.2
// 192.168.1.1
// 192.168.1.2
// 192.168.1.3
// 192.168.1.2
// 192.168.1.2
// 192.168.1.1
// 192.168.1.2
// 192.168.1.3
