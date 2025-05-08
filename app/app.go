package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application of comand line"
	app.Usage = "Search IP'S and names of server in the internet"
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP",
			Flags: flags,	
			Action: searchIps,
		},
		{
			Name:  "servers",
			Usage: "Search servers",
			Flags: flags,	
			Action: searchServers,
		},
	}
	
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, serve := range servers {
		fmt.Println(serve.Host)
	}
}