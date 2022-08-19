#!/bin/bash

# Usage: sh rerun.sh [index]
# [index]: 运行最近的第几个bin文件，默认是1
BIN_NAME=beego-starter
# 版本文件
VERSION_FILE='version.txt'

index=0
if [ -z $1 ]; then
  index=1
elif [ "`echo $1|sed 's/[^0-9]//g'`" != $1 ]; then
  echo '[index] need number'
else
  index=$1
fi

main(){
  if [ $index -lt 1 ]; then
    return 1
  fi
  binFilesCount=`ls libs/${BIN_NAME}-*|wc -l|sed 's/ //g'`
  if [ $binFilesCount -eq 0 ]; then
    return 1;
  fi
  if [ $index -gt $binFilesCount ]; then
    echo "[index] should be 1-$binFilesCount"
    return 1
  fi

  binFile=`ls libs/${BIN_NAME}-*|tail -${index}|head -1`

  echo "sh: cp ${binFile} ${BIN_NAME}"
  cp $binFile $BIN_NAME

  # save current version
  tempStr=${binFile##*-}
  version=${tempStr%%}
  echo $version > $VERSION_FILE

  echo 'sh: sh run.sh restart'
  sh run.sh restart
}

main
