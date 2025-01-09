#!/bin/bash

if [ ! -e /etc/systemd/system/spurtcms.service ]; then
    yes | cp -i spurtcms.service /etc/systemd/system/
fi

if [ $1 ]; then
    if [ $1 = "start" ]; then
        echo ""
        echo "Application Started Successfully"
        echo " --- Admin Running On http://localhost:8082"
        echo " --- Graphql Running On http://localhost:8084"
        echo " --- Preview Running On http://localhost:8083"
        echo "To Stop Application : 'sudo sh ./spurtcms.sh stop' "
        echo ""
        exec systemctl start spurtcms.service
    fi

    if [ $1 = "stop" ]; then
        echo ""
        echo "-- Application Stopped Successfully --"
        echo "-- To Start Application Again : 'sudo sh ./spurtcms.sh start' "
        echo ""
        exec systemctl stop spurtcms.service
    fi
fi
