package main

import (
	"fmt"
)

type Second struct{}

func (x Second) install_packages() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Some packages will be installed" + RESET)
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "This step may take a while" + RESET)
	press_return()
	system("clear")
	system("pacman -Sy --noconfirm vim grub dosfstools mtools networkmanager sudo os-prober xdg-user-dirs")
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Installed packages" + RESET)
	sleep(2)
}

func (x Second) timezone_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting date and time" + RESET)
	system("ln -sf /usr/share/zoneinfo/America/Sao_paulo /etc/localtime")
	system("hwclock --systohc")
	sleep(2)
}

func (x Second) language_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting language, keyboard and hostname" + RESET)
	system("echo en_US.UTF-8 UTF-8 >> /etc/locale.gen") // pt_BR.UTF-8 UTF-8
	system("locale-gen > /dev/null 2>&1")
	system("echo LANG=en_US.UTF-8 >> /etc/locale.conf") // LANG=pt_BR.UTF-8
	system("echo KEYMAP=br-abnt2 >> /etc/vconsole.conf")
	system("echo arch >> /etc/hostname")
	sleep(2)
}

func (x Second) hosts_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting hosts" + RESET)
	system("echo '127.0.0.1       localhost' > /etc/hosts")
	system("echo '::1             localhost' >> /etc/hosts")
	system("echo '127.0.1.1       arch.localdomain      arch' >> /etc/hosts")
	sleep(2)
}

func (x Second) users_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting users" + RESET)
	system("echo 'root:admin' | chpasswd")
	system("useradd -m -g users -G wheel administrator")
	system("echo '%wheel ALL=(ALL:ALL) ALL' >> /etc/sudoers")
	system("echo 'administrator:admin' | chpasswd")
	sleep(2)
}

func (x Second) users_dirs_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting users dirs" + RESET)
	system("xdg-user-dirs-update")
	sleep(2)
}

func (x Second) grub_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting grub" + RESET)
	system("grub-install --target=i386-pc --recheck /dev/sda > /dev/null 2>&1")
	system("grub-mkconfig -o /boot/grub/grub.cfg > /dev/null 2>&1")
	sleep(2)
}

func (x Second) network_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Setting NetworkManager" + RESET)
	system("systemctl enable NetworkManager > /dev/null 2>&1")
	sleep(2)
}

func (x Second) finish() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Manual step" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "1 - Run 'exit'" + RESET)
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "2 - Run 'umount -R /mnt'" + RESET)
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "3 - Run 'reboot'" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Welcome to Arch linux !!!" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Users:" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "User: root" + RESET)
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Password: admin" + RESET)
	fmt.Println("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "User: administrator" + RESET)
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Password: admin" + RESET)
	fmt.Println("")
	sleep(2)
}

func (x Second) run() {
	x.install_packages()
	x.timezone_config()
	x.language_config()
	x.hosts_config()
	x.users_config()
	x.users_dirs_config()
	x.grub_config()
	x.network_config()
	x.finish()
}
