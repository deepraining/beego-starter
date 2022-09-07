#!/bin/bash

# Usage: sh run.sh [start|stop|restart|status|version]
# 运行项目的脚本

# 二进制文件名
BIN_NAME=beego-starter
# Server env
SERVER_ENV=prod
# 版本文件（会将正在运行的版本号放在这个文件中）
VERSION_FILE='version.txt'

usage() {
  echo "Usage: sh $0 [start|stop|restart|status|version]"
  exit 1
}

running(){
  pid=`ps -ef|grep $BIN_NAME|grep -v grep|awk '{print $2}'`
  if [ -z "${pid}" ]; then
    return 1
  else
    return 0
  fi
}

start(){
  running
  if [ $? -eq "0" ]; then
    echo "${BIN_NAME} is running, pid=${pid} ."
  else
    RUN_MODE=$SERVER_ENV nohup ./$BIN_NAME > /dev/null 2>&1 &
    echo "${BIN_NAME} starts succeeded, and view the logs to confirm that program has already been started."
  fi
}

stop(){
  running
  if [ $? -eq "0" ]; then
    kill -9 $pid
    echo "${pid} has already been killed, and program stopped running."
  else
    echo "${BIN_NAME} is not running."
  fi
}

status(){
  running
  if [ $? -eq "0" ]; then
    echo "${BIN_NAME} is running, pid=${pid} ."
  else
    echo "${BIN_NAME} is not running."
  fi
}

restart(){
  stop
  start
}

version(){
  if [ -f $VERSION_FILE ]; then
    cat $VERSION_FILE
  else
    echo "No version found"
  fi
}

case "$1" in
  "start")
    start
    ;;
  "stop")
    stop
    ;;
  "status")
    status
    ;;
  "restart")
    restart
    ;;
  "version")
    version
    ;;
  *)
    usage
    ;;
esac
