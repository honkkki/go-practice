package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "new app"
	app.Usage = "test"
	app.Version = "1.0"

	app.Action = func(context *cli.Context) error {
		fmt.Println("hello golang")
		fmt.Println(context.App.Name)
		fmt.Println(context.Args().Get(0))
		fmt.Println(context.NArg(), context.Args().Slice())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
