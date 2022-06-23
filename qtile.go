package main

import (
	"fmt"
)

type QTILE struct{}

func (x QTILE) install_packages() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Some packages will be installed" + RESET)
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "This step may take a while" + RESET)
	press_return()
	system("clear")
	system("pacman -Sy --noconfirm xorg-server xorg-xinit xorg-apps qtile xfce4-terminal nitrogen picom pulseaudio alsa-utils i3lock archlinux-wallpaper")
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Installed packages" + RESET)
	sleep(2)
}

func (x QTILE) xinitrc_config() {
	system("clear")
	base64_to_file(XINITRC, "/home/administrator", ".xinitrc")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "XORG configured" + RESET)
	sleep(2)
}

func (x QTILE) qtile_config() {
	system("clear")
	base64_to_file(QTILE_CONFIG, "/home/administrator/.config/qtile", "config.py")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "QTILE configured" + RESET)
	sleep(2)
}

func (x QTILE) nitrogen_config() {
	system("clear")
	base64_to_file(NITROGEN_CONFIG, "/home/administrator/.config/nitrogen", "nitrogen.cfg")
	base64_to_file(NITROGEN_SAVE, "/home/administrator/.config/nitrogen", "bg-saved.cfg")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "NITROGEN configured" + RESET)
	sleep(2)
}

func (x QTILE) xfce4_terminal_config() {
	system("clear")
	base64_to_file(XFCE4_TERMINAL_CONFIG, "/home/administrator/.config/xfce4/terminal", "terminalrc")
	base64_to_file(BASHRC, "/home/administrator", ".bashrc")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "XFCE4-TERMINAL configured" + RESET)
	sleep(2)
}

// func (x QTILE) ly_dm_config() {
// 	system("clear")
// 	system("git clone https://aur.archlinux.org/brave-bin.git")
// 	system("chown -R administrator brave-bin/")
// 	system("mv brave-bin/ /home/administrator")
// 	system("runuser -l administrator -c 'cd brave-bin && makepkg --syncdeps --install --needed --noconfirm'")
// 	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "LY-DM configured" + RESET)
// 	sleep(2)
// }

func (x QTILE) finish() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Success !!!" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Use 'reboot' for finish the instalation" + RESET)
	fmt.Println("")
	sleep(2)
}

func (x QTILE) run() {
	x.install_packages()
	x.xinitrc_config()
	x.qtile_config()
	x.nitrogen_config()
	x.xfce4_terminal_config()
	//x.ly_dm_config()
	x.finish()
}
