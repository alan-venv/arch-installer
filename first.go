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
	var commands [6]string
	commands[0] = "parted /dev/sda -s mklabel gpt"
	commands[1] = "parted /dev/sda -s mkpart Boot fat32 0% 1GiB"
	commands[2] = "parted /dev/sda -s set 1 bios_grub on"
	commands[3] = "parted /dev/sda -s mkpart Swap linux-swap 1GiB 2GiB"
	commands[4] = "parted /dev/sda -s set 2 swap on"
	commands[5] = "parted /dev/sda -s mkpart Home ext4 2GiB 100%"

	for _, c := range commands {
		system(c)
	}

	sleep(2)
}

func (x First) run() {
	x.start()
	x.formatting()
	x.keyboard_config()
	x.build_hd()
}

//     def __mountPartitions(self):
//         os.system("clear")
//         print("[+] Mounting partitions")

//         os.system("mkfs.fat -F 32 /dev/sda1 > /dev/null 2>&1")
//         os.system("mkswap /dev/sda2 > /dev/null 2>&1")
//         os.system("mkfs.ext4 /dev/sda3 > /dev/null 2>&1")

//         os.system("mount /dev/sda3 /mnt > /dev/null 2>&1")
//         os.system("swapon /dev/sda2 > /dev/null 2>&1")
//         os.system("clear")

//         print("[+] Mounted partitions")
//         time.sleep(1)

//     def __installPackages(self):
//         os.system("clear")
//         print("[+] Essential packages will be installed")
//         print("[+] This step may take a while")
//         print("[+]")
//         input("[-] Press ENTER to continue...")
//         os.system("clear")
//         os.system("pacstrap /mnt base linux linux-firmware base-devel python3")
//         os.system("clear")
//         print("[+] Installed packages")
//         time.sleep(1)

//     def __generateFstab(self):
//         os.system("clear")
//         os.system("genfstab -U /mnt >> /mnt/etc/fstab")
//         os.system("clear")
//         print("[+] Generated fstab file")
//         time.sleep(1)

//     def __switchSystem(self):
//         os.system("clear")
//         print("[-] Manual step")
//         print("[-]")
//         print("[-] Access the directory '/root/Installer'")
//         print("[-] And run './installer.py --second' to continue the installation")
//         print("[-]")
//         input("[-] Press ENTER to continue...")
//         os.system("mkdir -p /mnt/root/Installer")
//         os.system("cp -r . /mnt/root/Installer")
//         os.system("arch-chroot /mnt")

//     def run(self):
//         self.__start()
//         self.__formatting()
//         self.__keyboardConfig()
//         self.__buildHD()
//         self.__mountPartitions()
//         self.__installPackages()
//         self.__generateFstab()
//         self.__switchSystem()
