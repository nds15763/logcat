#!/bin/bash

killall offertrack

sleep 1

nohup bin/offertrack > logs/worker.log 2>&1 &
