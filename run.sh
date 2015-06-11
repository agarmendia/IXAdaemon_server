#!/bin/bash
LANGUAGE="java"			#Programation language of the native program
NATIVE="atzerapena"		#native program name
MPORT=2101     			#Main Port
CPORT=2102				#Control Port

clear
IXAdaemon_ctrl
STATE=$?
echo
echo $STATE

if [ $STATE == "1" ]; then
	echo "erroraso"
	exit 1
fi

IXAdaemon_client