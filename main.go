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

// colorRed := "\033[31m"
// colorYellow := "\033[33m"
// colorBlue := "\033[34m"
// colorPurple := "\033[35m"
// colorCyan := "\033[36m"
// colorWhite := "\033[37m"
// fmt.Println("desired formatted output")
// fmt.Printf("<%-10s>  <%-10s>  <%-10s>\n\n", data[0],data[1],data[2])

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

	// prg := "sh"
	// arg1 := "-c"

	// cmd := exec.Command(prg, arg1, command)
	// stdout, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println("\"" + command + "\" command failed !!!")
	// 	fmt.Println(err.Error())
	// 	os.Exit(0)
	// }

	// return string(stdout)
}

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

func install_packages() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Some packages will be installed" + RESET)
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "This step may take a while" + RESET)
	press_return()
	system("clear")
	//system("pacman -Sy --noconfirm xorg-server xorg-xinit xorg-apps xfce4-terminal nitrogen picom pulseaudio alsa-utils bspwm sxhkd polybar rofi i3lock")
	system("pacman -Sy --noconfirm vim")
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Installed packages" + RESET)
	sleep(1)
}

func xinitrc_config() {
	system("clear")
	system("cp " + get_executable_path() + "/resources/xorg/.xinitrc /home/administrator")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + ".xinitrc configured" + RESET)
	sleep(1)
}

//	   def __xinitrc_config(self):
//         os.system("clear")
//         os.system(f"cp {self.project_path}/resources/.xinitrc /home/administrator")
//         print("[+] .xinitrc configured")
//         time.sleep(1)

//     def __qtile_config(self):
//         os.system("clear")
//         os.makedirs("/home/administrator/.config/qtile/")
//         os.system(f"cp {self.project_path}/resources/config.py /home/administrator/.config/qtile")
//         print("[+] Qtile configured")
//         time.sleep(1)

//     def __finish(self):
//         os.system("clear")
//         print("[-] Success !!!")
//         print("[-]")
//         print("[-] Use 'startx' with the administrator to start the qtile")
//         print("[-]")
//         time.sleep(1)

func main() {
	id := os.Geteuid()
	if id == 0 {
		// install_packages()
		xinitrc_config()
	} else {
		fmt.Println("Run this script with root permissions")
	}
}
