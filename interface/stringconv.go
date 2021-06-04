package main
	import "fmt"
	type IPAddr [4]byte
	func main() {
		hosts := map[string]IPAddr{
			"loopback": {127, 0, 0, 1},
			"googleDNS": {8, 8, 8, 8},
		}
		for name, ip := range hosts {
			fmt.Printf("%v: %s\n", name, ip)
		}
		
	}

func (ip IPAddr) String() string {
	var a string
	for i := range ip {
		a += fmt.Sprintf("%d.", ip[i])
		}
	a = a[:len(a)-1]
	return fmt.Sprintf(a)
}

