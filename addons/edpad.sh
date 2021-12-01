#! /bin/bash

if killall -q -s 0 edpad 2>/dev/null ; then
	killall -q edpad
else
	$HOME/.local/bin/edpad -d -x game-host:0.0 2>$HOME/edpad.log
fi

