package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {

		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("qrchaka.com:%d", j)

			con, err := net.Dial("tcp", address)
			if err != nil {
				return
			}

			fmt.Printf("open port : %d\n", j)
			con.Close()
		}(i)

	}

	wg.Wait()
}
