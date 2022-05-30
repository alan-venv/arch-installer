package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"time"
)

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

func base64_to_file(b64 string, path string, file string) {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path + "/" + file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = f.Chmod(0755)
	if err != nil {
		panic(err)
	}

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}

	if err := f.Sync(); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 2 {
		if os.Geteuid() == 0 {
			if os.Args[1] == "--bspwm" {
				bspwm := BSPWM{}
				bspwm.run()
			} else if os.Args[1] == "--first" {
				first := First{}
				first.run()
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
