#!/bin/bash

if [ ! -e /etc/systemd/system/spurtcms-admin.service ]; then
    yes | cp -i spurtcms-admin.service /etc/systemd/system/
fi

if [ $1 ]; then
    if [ $1 == "start" ]; then
        echo "Application Started Successfully - Visit http://localhost:8080"
        echo "To Stop Application  - "sudo ./spurtcms-admin.sh stop" "
        exec systemctl start spurtcms-admin.service
    fi

    if [ $1 == "stop" ]; then
        echo "-- Application Stopped Successfully --"
        echo "-- To Start Application Again - "sudo ./spurtcms-admin.sh start" "
        exec systemctl stop spurtcms-admin.service
    fi
fi
