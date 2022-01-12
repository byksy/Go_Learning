package main

import "fmt"

func main() {
	const (
		ETH = 9 - iota
		WAN
		ICX
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 20,
	}

	fmt.Printf("1 BTC şs %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC şs %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}
