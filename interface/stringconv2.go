package main
import ("fmt"
 "strconv"
 "strings")
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
	res := []string{}
	for _,val := range ip {
   		res = append(res,strconv.Itoa(int(val)))
	}
	strings.Join(res,".")
	return fmt.Sprintf("%v", res)
}

