#!/bin/bash
LANGUAGE="java"
NATIVE="atzerapena"
clear
serveIP_ctrl
if [ $? == "3" ]; then
	n=`ps aux | grep atzerapena | grep java | wc -l`
	if (($n > "0")); then
		nativeProcesses=(`ps aux | grep atzerapena | pidof java`)
		
		for (( i=0; i<$n; i++ ))
		do
			kill ${nativeProcesses[$i]}
		done
	fi
	serveIP_server $LANGUAGE $NATIVE &
fi

serveIP_ctrl
STATE=$?

if [ $STATE == "2" ]; then
	echo "restart serveIP_server"
	PID=`ps aux | grep $NATIVE | pidof serveIP_server`
	kill $PID
	serveIP_server $LANGUAGE $NATIVE &
	echo "started process"

fi

while [ $STATE -ne "0" ]; do
	echo -n "..."
	sleep 1
	serveIP_ctrl
	STATE=$?
done

echo
echo "running client"

serveIP_client