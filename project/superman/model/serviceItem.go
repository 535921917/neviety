package model

import (
	"fmt"
	"strconv"
)

type ServiceItem struct {
	PSM  string
	IP   string
	Port int64
	TTL  int64
}

func (s *ServiceItem) GetKey() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}

func NewServiceItem(PSM, IP, port string) *ServiceItem {
	portNum, _ := strconv.ParseInt(port, 10, 64)
	return &ServiceItem{
		PSM:  PSM,
		IP:   IP,
		Port: portNum,
	}

}
