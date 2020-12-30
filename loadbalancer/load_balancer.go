package loadBalancer

import (
	"errors"
)

// Server 服务器
type Server struct {
	IP            string
	Weight        int
	CurrentWeight int
}

// Client 客户端
type Client struct {
	Servers []*Server
}

// Add 添加服务器
func (c *Client) Add(IP string, weight int) error {
	if IP == "" {
		return errors.New("ip address is not available")
	}
	if weight < 0 {
		return errors.New("weight can not be negative")
	}
	server := &Server{
		IP:            IP,
		Weight:        weight,
		CurrentWeight: weight,
	}
	c.Servers = append(c.Servers, server)
	return nil
}

// LoadBalancer 负载均衡
func (c *Client) LoadBalancer() (*Server, error) {
	// total 总权重
	var totalWeight int
	maxWeightServer := &Server{}
	for _, server := range c.Servers {
		totalWeight += server.Weight
		if maxWeightServer == nil || server.CurrentWeight > maxWeightServer.CurrentWeight {
			maxWeightServer = server
		}
	}
	maxWeightServer.CurrentWeight = maxWeightServer.CurrentWeight - totalWeight
	for _, server := range c.Servers {
		server.CurrentWeight = server.CurrentWeight + server.Weight
	}
	return maxWeightServer, nil
}
