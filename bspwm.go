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
	system("pacman -Sy --noconfirm xorg-server xorg-xinit xorg-apps xfce4-terminal nitrogen picom pulseaudio alsa-utils bspwm sxhkd polybar rofi i3lock archlinux-wallpaper")
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Installed packages" + RESET)
	sleep(1)
}

func (x BSPWM) xinitrc_config() {
	system("clear")
	base64_to_file(XINITRC, "/home/administrator", ".xinitrc")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "XORG configured" + RESET)
	sleep(2)
}

func (x BSPWM) bspwm_config() {
	system("clear")
	base64_to_file(BSPWMRC, "/home/administrator/.config/bspwm", "bspwmrc")
	base64_to_file(RESIZE_BSPWM, "/home/administrator/.config/bspwm/scripts", "resize.sh")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "BSPWM configured" + RESET)
	sleep(2)
}

func (x BSPWM) sxhkd_config() {
	system("clear")
	base64_to_file(SXHKDRC, "/home/administrator/.config/sxhkd", "sxhkdrc")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "SXHKD configured" + RESET)
	sleep(2)
}

func (x BSPWM) polybar_config() {
	system("clear")
	base64_to_file(POLYBAR_CONFIG, "/home/administrator/.config/polybar", "config.ini")
	base64_to_file(POLYBAR_LAUNCHER, "/home/administrator/.config/polybar", "launch.sh")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "POLYBAR configured" + RESET)
	sleep(2)
}

func (x BSPWM) nitrogen_config() {
	system("clear")
	base64_to_file(NITROGEN_CONFIG, "/home/administrator/.config/nitrogen", "nitrogen.cfg")
	base64_to_file(NITROGEN_SAVE, "/home/administrator/.config/nitrogen", "bg-saved.cfg")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "NITROGEN configured" + RESET)
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
	x.nitrogen_config()
	x.finish()
}
