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
	
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "google.com",
				},
			},	
			Action: searchIps,

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