package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func maian() {
	start := time.Now()

	fmt.Println("Program Başarıyla Başladı")

	var wg sync.WaitGroup // Bir WaitGroup nesnesi oluştur

	for ia := 0; ia < 100; ia++ {
		ip := "drive.google.com"

		wg.Add(1) // Her Go rutini için bir sayaç arttır

		go func(port int) { // Her port için yeni bir Go rutini başlat
			defer wg.Done() // Go rutini bittiğinde sayaç azalt

			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 10*time.Second) // 10 saniye içinde bağlantı kurmaya çalış
			if err == nil {
				fmt.Println("Port ", port, " açık")
				conn.Close()
			} else {
				fmt.Println("Port ", port, " kapalı veya zaman aşımına uğradı")
			}

		}(ia) // Go rutinine port numarasını parametre olarak geçir
	}

	wg.Wait() // Tüm Go rutinleri bitene kadar bekle

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Geçen süre:", elapsed)

}
