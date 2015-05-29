#!/bin/bash
LANGUAGE="java"
NATIVE="atzerapena"
clear
serveIP_ctrl
if [ $? == "3" ]; then
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

echo -n "waiting for initialization..."
while [ $STATE -ne "0" ]; do
	echo -n "..."
	sleep 1
	serveIP_ctrl
	STATE=$?
done

echo
echo "running client"

serveIP_client