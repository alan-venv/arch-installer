package main

import (
	"fmt"
)

type First struct{}

func (x First) start() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Initializing" + RESET)
	sleep(2)
}

func (x First) formatting() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Formatting the primary HD" + RESET)
	system("wipefs -a /dev/sda > /dev/null 2>&1")
	sleep(2)
}

func (x First) keyboard_config() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Configuring the keyboard" + RESET)
	system("loadkeys br-abnt2 > /dev/null 2>&1")
	sleep(2)
}

func (x First) build_hd() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Configuring the drive" + RESET)
	system("parted /dev/sda -s mklabel gpt")
	system("parted /dev/sda -s mkpart Boot fat32 0% 1GiB")
	system("parted /dev/sda -s set 1 bios_grub on")
	system("parted /dev/sda -s mkpart Swap linux-swap 1GiB 2GiB")
	system("parted /dev/sda -s set 2 swap on")
	system("parted /dev/sda -s mkpart Home ext4 2GiB 100%")
	sleep(2)
}

func (x First) mount_partitions() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Mounting partitions" + RESET)
	system("mkfs.fat -F 32 /dev/sda1 > /dev/null 2>&1")
	system("mkswap /dev/sda2 > /dev/null 2>&1")
	system("mkfs.ext4 /dev/sda3 > /dev/null 2>&1")
	system("mount /dev/sda3 /mnt > /dev/null 2>&1")
	system("swapon /dev/sda2 > /dev/null 2>&1")
	sleep(2)
}

func (x First) install_packages() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Essential packages will be installed" + RESET)
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "This step may take a while" + RESET)
	press_return()
	system("clear")
	system("pacstrap /mnt base linux linux-firmware base-devel")
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Installed packages" + RESET)
	sleep(2)
}

func (x First) generate_fstab() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[+] " + RESET + GREEN + "Creating fstab file" + RESET)
	system("genfstab -U /mnt >> /mnt/etc/fstab")
	sleep(2)
}

func (x First) switch_system() {
	system("clear")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "Manual step" + RESET)
	print("")
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "1 - Run 'arch-chroot /mnt'" + RESET)
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "2 - Run 'cd /root'" + RESET)
	fmt.Println(BOLD_GREEN + "[-] " + RESET + GREEN + "3 - Run './installer --second' to continue the installation" + RESET)
	press_return()
	system("cp installer /mnt/root")
	// system("arch-chroot /mnt") // TODO: Error, fix this!
}

func (x First) run() {
	x.start()
	x.formatting()
	x.keyboard_config()
	x.build_hd()
	x.mount_partitions()
	x.install_packages()
	x.generate_fstab()
	x.switch_system()
}
