package  main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main()  {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "special the language use",
		},
		cli.BoolFlag{
			Name:   "readonly, ro",
			Hidden: true,
			Usage:  "whether this db is ro",
		},
	}

	app.Commands = []cli.Command{
		{

		 Name: "run",
		 Usage: "this is for run command",

		 Flags: []cli.Flag {

		 	cli.StringFlag{
		      Name: "name,n",
		      Usage: "special the name of the container",
			},

			cli.IntFlag{
				Name: "port, p",
			    Usage: "port invalid",

			},
		 },

		 Action: func(ctx *cli.Context)error {
		 	  fmt.Println(ctx.String("n"))
			 fmt.Println(ctx.Int("port"), ctx.Args().First())
			 return nil
		 },
		},
	}

	app.Action = func(ctx *cli.Context)error {
	   fmt.Println(ctx.String("l"))
	   fmt.Println(ctx.Bool("ro"))
	   return  nil
	}


	if err := app.Run(os.Args);err!= nil {
		panic("error")
	}


}
