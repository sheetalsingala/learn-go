package main

import (
	"fmt"
	"net"
	"math/big"
)

func NextIP(ip net.IP) net.IP {
	i := ipToInt(ip)
	return intToIP(i.Add(i, big.NewInt(1)))
}

func ipToInt(ip net.IP) *big.Int {
	if v := ip.To4(); v != nil {
		return big.NewInt(0).SetBytes(v)
	}
	return big.NewInt(0).SetBytes(ip.To16())
}

func intToIP(i *big.Int) net.IP {
	return net.IP(i.Bytes())
}

func main(){
        ip := net.ParseIP("1.0.0.0")
        fmt.Println("Could not ", ip.String())

        fmt.Println(NextIP(ip))
}
