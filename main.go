package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

const RESET string = "\033[0m"
const BOLD_GREEN string = "\033[1;32m"
const GREEN string = "\033[32m"
const RED string = "\033[31m"

func get_home_path() string {
	home := os.Getenv("XDG_CONFIG_HOME")
	if home != "" {
		return home
	} else {
		return os.Getenv("HOME")
	}
}

func get_executable_path() string {
	return path.Dir(os.Args[0])
}

func press_return() string {
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Press ENTER to continue..." + RESET)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return input
}

func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func system(command string) {
	prg := "sh"
	arg1 := "-c"

	cmd := exec.Command(prg, arg1, command)
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("\"" + command + "\" command failed !!!")
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) == 2 {
		if os.Geteuid() == 0 {
			if os.Args[1] == "--bspwm" {
				bspwm := BSPWM{}
				bspwm.run()
			} else if os.Args[1] == "--first" {
				fmt.Println(RED + "In Development" + RESET)
			} else if os.Args[1] == "--second" {
				fmt.Println(RED + "In Development" + RESET)
			} else {
				fmt.Println(RED + "Invalid option" + RESET)
			}
		} else {
			fmt.Println(RED + "Run this script with root permissions" + RESET)
		}
	} else {
		system("clear")
		fmt.Println(GREEN + "ArchLinux Installer" + RESET)
		fmt.Println("")
		fmt.Println("OPTIONS:")
		fmt.Println("")
		fmt.Println("--first      Initializes the first step of the script.")
		fmt.Println("--second     Initializes the second part of the script.")
		fmt.Println("--bspwm      Initializes the bspwm configuration.")
		fmt.Println("")
	}
}
