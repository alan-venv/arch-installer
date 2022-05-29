package main

import (
	"fmt"
)

type BSPWM struct{}

func (x BSPWM) install_packages() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Some packages will be installed" + RESET)
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "This step may take a while" + RESET)
	press_return()
	system("clear")
	system("pacman -Sy --noconfirm xorg-server xorg-xinit xorg-apps xfce4-terminal nitrogen picom pulseaudio alsa-utils bspwm sxhkd polybar rofi i3lock")
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Installed packages" + RESET)
	sleep(1)
}

func (x BSPWM) xinitrc_config() {
	system("clear")
	system("cp " + get_executable_path() + "/resources/xorg/.xinitrc /home/administrator")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + ".xinitrc configured" + RESET)
	sleep(2)
}

func (x BSPWM) bspwm_config() {
	system("clear")
	system("cp -r " + get_executable_path() + "/resources/bspwm /home/administrator/.config")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "BSPWM configured" + RESET)
	sleep(2)
}

func (x BSPWM) sxhkd_config() {
	system("clear")
	system("cp -r " + get_executable_path() + "/resources/sxhkd /home/administrator/.config")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "SXHKD configured" + RESET)
	sleep(2)
}

func (x BSPWM) polybar_config() {
	system("clear")
	system("cp -r " + get_executable_path() + "/resources/polybar /home/administrator/.config")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "POLYBAR configured" + RESET)
	sleep(2)
}

func (x BSPWM) finish() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Success !!!" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Use 'startx' with the administrator to start the bspwm" + RESET)
	fmt.Println("")
	sleep(2)
}

func (x BSPWM) run() {
	x.install_packages()
	x.xinitrc_config()
	x.bspwm_config()
	x.sxhkd_config()
	x.polybar_config()
	x.finish()
}
