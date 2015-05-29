#!/bin/bash
LANGUAGE="java"			#Programation language of the native program
NATIVE="atzerapena"		#native program name
MPORT=2101     			#Main Port
CPORT=2102				#Control Port
clear
serveIP_ctrl

#if there is not a server running (with or without a running native instance)
if [ $? == "3" ]; then
	n=`ps aux | grep $NATIVE | grep $LANGUAGE | wc -l`
	if (($n > "0")); then
		nativeProcesses=(`ps aux | grep $NATIVE | pidof $LANGUAGE`)
		
		for (( i=0; i<$n; i++ ))
		do
			kill ${nativeProcesses[$i]}
		done
	fi
	serveIP_server --mainPort=$MPORT --ctrlPort=$CPORT $LANGUAGE $NATIVE &
fi

serveIP_ctrl
STATE=$?
#if there is not a native instance running
if [ $STATE == "2" ]; then
	echo "restart serveIP_server"
	PID=`ps aux | grep $NATIVE | pidof serveIP_server`
	kill $PID
	serveIP_server --mainPort=$MPORT --ctrlPort=$CPORT $LANGUAGE $NATIVE &
	echo "started process"

fi

# Waiting for the native instance to initializate
while [ $STATE -ne "0" ]; do
	echo -n "..."
	sleep 1
	serveIP_ctrl --ctrlPort=$CPORT
	STATE=$?
done

#everything runs OK
echo
echo "running client"

serveIP_client --port=$MPORT