package main

import (
	"gopkg.in/yaml.v1"  cli
	"fmt"
	"os"
)

func main(){
	app := cli.NewApp();
	app.Name = "hello_cli"
	app.Usage = "Print Hello Word"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)
		return nil
	}

	app.Run(os.Args)
}
