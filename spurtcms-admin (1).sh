#!/bin/bash

if [ ! -e /etc/systemd/system/spurtcms-admin.service ]; then
    yes | cp -i spurtcms-admin.service /etc/systemd/system/
fi

if [ $1 ]; then
    if [ $1 = "start" ]; then
        echo ""
        echo "Application Started Successfully - Visit http://localhost:8080"
        echo "To Stop Application : 'sudo sh ./spurtcms-admin.sh stop' "
        echo ""
        exec systemctl start spurtcms-admin.service
    fi

    if [ $1 = "stop" ]; then
        echo ""
        echo "-- Application Stopped Successfully --"
        echo "-- To Start Application Again : 'sudo sh ./spurtcms-admin.sh start' "
        echo ""
        exec systemctl stop spurtcms-admin.service
    fi
fi
