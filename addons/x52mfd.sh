#! /bin/bash

EDHOST='user@game-host'

IS_RUNNING=`ssh -XC $EDHOST 'systemctl --user is-active x52mfd'`

case $IS_RUNNING in
	active)
		ssh -XC $EDHOST 'systemctl --user stop x52mfd' |\
			zenity --progress --no-cancel --text 'STOPPING x52mfd' --timeout 3
		;;
	*) 
		ssh -XC $EDHOST 'systemctl --user start x52mfd' |\
			zenity --progress --no-cancel --text 'STARTING x52mfd' --timeout 3
		;;
esac


