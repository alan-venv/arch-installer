Utility script for installing ArchLinux
==========
Follow the instructions for correct use and configuration

Manual installation
------------

1 - Configuring the keyboard::

    loadkeys br-abnt2
    
2 - Formatting the primary HD::

    wipefs -a /dev/sda
    
3 - Create a GPT table::

    fdisk /dev/sda    
    
    [Options: 'g' and 'w']

4 - Create the partitions::

    cfdisk /dev/sda
    
    [1G - BIOS Boot]
    [1G - Linux swap]
    [ALL - File System]
    
5 - Format the partitions::

    mkfs.fat -F 32 [BIOS Boot]
    mkfs.ext4 [File System]
    mkswap [Swap]
    
6 - Mount the partitions::

    mount [File System] /mnt
    swapon [Swap]
    
7 - Install essential packages::

    pacstrap /mnt base linux linux-firmware base-devel
    
8 - Generate FSTAB::

    genfstab -U /mnt >> /mnt/etc/fstab

9 - Enter de system::

    arch-chroot /mnt

10 - Install essentials packages::

    pacman -S vim grub dosfstools mtools networkmanager sudo os-prober
    
11 - Set the time::

    ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

    hwclock --systohc

12 - Adjust the language::

    vim /etc/locale.gen
    
    [Uncomment line "pt_BR.UTF-8 UTF-8"]
    
    locale-gen
    
13 - Language, keyboard and hostname settings::

    echo LANG=pt_BR.UTF-8 >> /etc/locale.conf
    echo KEYMAP=br-abnt2 >> /etc/vconsole.conf
    echo arch >> /etc/hostname

14 - Config hosts::
    
    vim /etc/hosts
    
    127.0.0.1   localhost
    ::1         localhost
    127.0.1.1	arch.localdomain    arch

15 - Change root password::

    passwd
    
16 - Create a new user::

    useradd -m -g users -G wheel administrator
    passwd administrator
    
17 - Sudo condig::

    vim /etc/sudoers
    
    [Uncoment line "%wheel ALL=(ALL:ALL) ALL"]
    
18 - Grub Configs::
    
    grub-install --target=i386-pc --recheck /dev/sda
    grub-mkconfig -o /boot/grub/grub.cfg

19 - Finishing::

    systemctl enable NetworkManager
    exit
    umount -R /mnt
    shutdown now
    
    
    Obs: Remove ISO CD
