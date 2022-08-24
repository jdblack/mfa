package main

import (
	"fmt"
	"time"
)

func printarr(in map[string]string) {
	fmt.Printf("----------------\n")
	for k, v := range in {
		fmt.Printf("%s : %s\n", v, k)
	}
	fmt.Printf("----------------\n")
}

func main() {
	otps := OtpConfig{}
	otps.init("/Users/jblack/.otps")
	for {
		fmt.Println("---------")
		for _, o := range otps.get() {
			fmt.Printf("%s\n    %s\n", o.name, o.token)
		}
		time.Sleep(1 * time.Second)
	}

}
