# DnsGo
Make DNS go away with a DOS attack

## The core idea
A lot of networks use a self-hosted DNS server, and, especially in home networks, the server is often-times a potato. One machine sending a bunch of queries is enough to overload the server, and make it impossible for the whole network to access the web and other services that need DNS to work.

## Why I made this?
This project wouldn't be interesting if it was just a bunch of threads sending DNS queries in a while-loop. The fun part was making a DNS client from scratch. Obviously this is for educational purposes only.

## Usage
First, build the program:
```bash
go build -o dnsgo cmd/dnsgo/main.go
chmod u+x ./dnsgo
```

Now you can run it:
```
❯ ./dnsgo -help

	 ____  _   _ ____      ____       _
	|  _ \| \ | / ___|    / ___| ___ | |
	| | | |  \| \___ \   | |  _ / _ \| |
	| |_| | |\  |___) |  | |_| | (_) |_|
	|____/|_| \_|____( )  \____|\___/(_)
					 |/
	
Usage of ./dnsgo:
  -host string
    	Target host
  -length int
    	Length of requested domain (default 10)
  -port int
    	Target port (default 53)
  -threads int
    	Number of threads (default 30)

```

Example:
```
./dnsgo -host 192.168.178.1
```


