package methodsandinterface

import "fmt"

type ipAddr [4]byte

//TODO: Add a String() string method to IPAddr

func (ipa ipAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
}

//StringersExersize needs comment
func StringersExersize() {
	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
