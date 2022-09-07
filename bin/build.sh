#!/bin/bash

# Usage: sh build.sh
# 因为在开发服务器上会经常更新代码，所以每次都用持续集成构建就不太方便
# 一个解决方案是增量更新代码到服务器上，使用此脚本在服务器上构建

set -e

# 构建目录
BUILD_DIR='project'
# 二进制文件名
BIN_NAME=beego-starter

main(){
  # change to build directory
  echo "sh: cd ${BUILD_DIR}"
  cd ${BUILD_DIR}

  # execute script
  binName="${BIN_NAME}-$(date +'%Y.%m%d.%H%M')"
  echo "sh: go build -o ${binName}"
  go build -o ${binName}

  # move to ../libs
  echo "sh: mv ${binName} ../libs"
  mv ${binName} ../libs

  # change to root directory
  echo "sh: cd ../"
  cd ../

  echo 'sh: sh select.sh'
  sh select.sh
}

main
