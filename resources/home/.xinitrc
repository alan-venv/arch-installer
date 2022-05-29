setxkbmap -model abnt2 -layout br
sxhkd &
picom &
nitrogen --restore &
xsetroot -cursor_name left_ptr
exec bspwm
