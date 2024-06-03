package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh meow",
		Short: "gh meow is a command to print cats in the terminal",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return printMeow(
				args...,
			)
		},
		Long: `gh meow is a command to print cats in the terminal
		You can print different types of cats by passing the following arguments:
		- xl: Chonk XL Meow!
		- nyan: Nyan Cat Meow!
		- hey: Hey Cutie!
		- cute: Cute Meow!
		- sleep: Sleeping Meow!
		- no argument: Meow!
		`,
	}

	return cmd
}

func printMeow(
	splArgs ...string,
) error {
	if len(splArgs) == 0 {
		fmt.Println(` /\_/\`)
		fmt.Println(`( o.o )`)
		fmt.Println(` > ^ <`)
		fmt.Println("> Meow!")

		return nil
	}
	for _, arg := range splArgs {
		if arg == "xl" {
			fmt.Println(`───────────────────────────────────────`)
			fmt.Println(`───▐▀▄───────▄▀▌───▄▄▄▄▄▄▄─────────────`)
			fmt.Println(`───▌▒▒▀▄▄▄▄▄▀▒▒▐▄▀▀▒██▒██▒▀▀▄──────────`)
			fmt.Println(`──▐▒▒▒▒▀▒▀▒▀▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▀▄────────`)
			fmt.Println(`▀█▒▒▒█▌▒▒█▒▒▐█▒▒▒▀▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌─────`)
			fmt.Println(`▀▌▒▒▒▒▒▒▀▒▀▒▒▒▒▒▒▀▀▒▒▒▒▒▒▒▒▒▒▒▒▒▒▐───▄▄`)
			fmt.Println(`▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌▄█▒█`)
			fmt.Println(`▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█▒█▀─`)
			fmt.Println(`▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█▀───`)
			fmt.Println(`▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌────`)
			fmt.Println(`─▌▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▐─────`)
			fmt.Println(`─▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌─────`)
			fmt.Println(`──▌▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▐──────`)
			fmt.Println(`──▐▄▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▄▌──────`)
			fmt.Println(`────▀▄▄▀▀▀▀▀▄▄▀▀▀▀▀▀▀▄▄▀▀▀▀▀▄▄▀────────`)
			fmt.Println(`──────── Chonk XL Meow! ───────────────`)
		}

		if arg == "nyan" {

			fmt.Println(`░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░`)
			fmt.Println(`░░░░░░░░░░▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄░░░░░░░░░`)
			fmt.Println(`░░░░░░░░▄▀░░░░░░░░░░░░▄░░░░░░░▀▄░░░░░░░`)
			fmt.Println(`░░░░░░░░█░░▄░░░░▄░░░░░░░░░░░░░░█░░░░░░░`)
			fmt.Println(`░░░░░░░░█░░░░░░░░░░░░▄█▄▄░░▄░░░█░▄▄▄░░░`)
			fmt.Println(`░▄▄▄▄▄░░█░░░░░░▀░░░░▀█░░▀▄░░░░░█▀▀░██░░`)
			fmt.Println(`░██▄▀██▄█░░░▄░░░░░░░██░░░░▀▀▀▀▀░░░░██░░`)
			fmt.Println(`░░▀██▄▀██░░░░░░░░▀░██▀░░░░░░░░░░░░░▀██░`)
			fmt.Println(`░░░░▀████░▀░░░░▄░░░██░░░▄█░░░░▄░▄█░░██░`)
			fmt.Println(`░░░░░░░▀█░░░░▄░░░░░██░░░░▄░░░▄░░▄░░░██░`)
			fmt.Println(`░░░░░░░▄█▄░░░░░░░░░░░▀▄░░▀▀▀▀▀▀▀▀░░▄▀░░`)
			fmt.Println(`░░░░░░█▀▀█████████▀▀▀▀████████████▀░░░░`)
			fmt.Println(`░░░░░░████▀░░███▀░░░░░░▀███░░▀██▀░░░░░░`)
			fmt.Println(`░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░`)
			fmt.Println(`──────── Nyan Cat Meow! ───────────────`)
		}

		if arg == "hey" {
			fmt.Println(`───────────────────────────────────────`)
			fmt.Println(`──────────────────────────────▓▓█───────`)
			fmt.Println(`────────────────────────────▒██▒▒█──────`)
			fmt.Println(`───────────────────────────█▓▓▓░▒▓▓─────`)
			fmt.Println(`─────────────────────────▒█▓▒█░▒▒▒█─────`)
			fmt.Println(`────────────────────────▒█▒▒▒█▒▒▒▒▓▒────`)
			fmt.Println(`─▓▓▒░──────────────────▓█▒▒▒▓██▓▒░▒█────`)
			fmt.Println(`─█▓▓██▓░──────────────▓█▒▒▒▒████▒▒▒█────`)
			fmt.Println(`─▓█▓▒▒▓██▓░──────────▒█▒▒▒▒▒██▓█▓░░▓▒───`)
			fmt.Println(`─▓▒▓▒▒▒▒▒▓█▓░──░▒▒▓▓██▒▒▒▒▒▒█████▒▒▒▓───`)
			fmt.Println(`─▓░█▒▒▒▒▒▒▒▓▓█▓█▓▓▓▓▒▒▒▒▒▒▒▒██▓██▒░▒█───`)
			fmt.Println(`─▓░▓█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓████▒▒▒█───`)
			fmt.Println(`─▓░▓██▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒▒▓██░░░█───`)
			fmt.Println(`─▓░▓███▒▒▒▒▒▒▒▒▒▒▒▓█▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▒▓▓──`)
			fmt.Println(`─▒▒▒██▓▒▓█▓▒▒▒▒▒▒▒▓▒▒▒▒▒▒▓▓▓▒▒▒▒▒▒▒▓▒█──`)
			fmt.Println(`──▓▒█▓▒▒▒▒▓▒▒▒▒▒▒▒▒▒▒▒▓█▓▓▓▓█▓▒▒▒▒▒▒▒▓▒─`)
			fmt.Println(`──▓▒█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓──────▓█▓▒▒▒▒▒▓█─`)
			fmt.Println(`──▒▒▓▒▒▒▓▓▓▒▒▒▒▒▒▒▒▒▓▓───░▓▓───█▓▒▒▒▒▒█─`)
			fmt.Println(`───█▒▒▓▓▓▒▒▓▓▒▒▒▒▒▒▓▓───█████▓──█▓▒▒▒▒▓▒`)
			fmt.Println(`───▓▓█▒─────▒▓▒▒▒▒▒█───░██████──░█▒▒▒▒▓▓`)
			fmt.Println(`───▓█▒──▒███─▒▓▒▒▒▒█────██████───▓▒▒▒▒▒▓`)
			fmt.Println(`───██───█████─█▒▒▒▒█─────███▓────▓▓▒▒▒▒▓`)
			fmt.Println(`───█▓───█████─▒▓▒▒▒█─────────────█▓▓▓▒▒▓`)
			fmt.Println(`───█▓───░███──░▓▒▒▒▓█──────────░█▓▒▒▒▓▒▓`)
			fmt.Println(`───██─────────▒▓▒▒▒▒▓▓──────░▒▓█▓────░▓▓`)
			fmt.Println(`───▓█░────────█▓██▓▒▒▓█▓▓▓▓██▓▓▒▓▒░░▒▓▒▓`)
			fmt.Println(`───▒██░──────▓▒███▓▒▒▒▒▓▓▓▓▒▒▒▒▒▒▓▓▓▓▒▓─`)
			fmt.Println(`────█▓█▓▓▒▒▓█▓▒░██▒▒▓▓█▓▒▒▒▒▒▒▒▒▒▒▒▒▓▓█▒`)
			fmt.Println(`────▓─░▓▓▓▓▓▒▓▓▓▓▒▓▓▓▒▓▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓`)
			fmt.Println(`────▒▒▒▓▒▒▒▒▒▒▓█░─░░░─▓▓▒▒▒▒▒▒▒▒▒▒▒▓██▓▒`)
			fmt.Println(`─────█▓▒▒▒▒▒▒▒▒▓▓─░░░─▓▓▒▒▒▒▒▒▒▒▒▓▓▓▒▒▓▒`)
			fmt.Println(`──────██▓▓▒▒▒▒▒▒█▒░░░░█▒▒▒▒▒▒▒▒▓█▓▓▒▒▒▒▒`)
			fmt.Println(`─────░─▒██▓▓▒▒▒▒▒█▓▒▒▓▒▒▒▒▒▒▓███▓▒▒▒▒▒▓▓`)
			fmt.Println(`──────────░▒▓▓▓▓▒▒▓▓▓▓▓▓████▓▓█▒▒▒▒▒▓▓█░`)
			fmt.Println(`─────────────Hey Cutie!────────────────`)
		}

		if arg == "cute" {
			fmt.Println(`──────────────`)
			fmt.Println(`  /\_/\  (    `)
			fmt.Println(` ( ^.^ ) _)   `)
			fmt.Println(`   \"/  (     `)
			fmt.Println(` ( | | )      `)
			fmt.Println(`(__d b__)     `)
			fmt.Println(`──Cute Meow!──`)
		}

		if arg == "sleep" {
			fmt.Println(`───────────────`)
			fmt.Println(` |\__/,|   ('\'`)
			fmt.Println(` |_ _  |.--.) )`)
			fmt.Println(` ( T   )      /`)
			fmt.Println(`(((^_(((/(((_/ `)
			fmt.Println(`────────────────`)
		}
	}
	return nil
}

func main() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
