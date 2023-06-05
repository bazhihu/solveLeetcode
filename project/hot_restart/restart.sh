#!/bin/sh

ps aux | grep wingo
count=`ps -ef |grep "wingo" | grep -v "grep" | wc -l`
echo ""

if [0==$count]; then
    echo "Wingo starting..."
    sudo ./wingo &
    echo "Wingo started"
else 
    echo "Wingo Restarting..."
    sudo kill -USR2 $(ps -ef | grep "wingo" | grep -v grep | awk '{print $2}')
    echo "Wingo Restarted"
fi

sleep 1

ps aux | grep wingo