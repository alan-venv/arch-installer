#!/bin/bash

RESOURCES=`dirname $0`

cp $HOME/.xinitrc $RESOURCES/xorg
cp -r $HOME/.config/bspwm $RESOURCES
cp -r $HOME/.config/sxhkd $RESOURCES
cp -r $HOME/.config/polybar $RESOURCES