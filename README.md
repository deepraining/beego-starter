# beego-starter

用于快速创建 [Beego](https://beego.vip/) 应用的模板脚手架。

项目功能参考 [spring-boot-starter](https://github.com/deepraining/spring-boot-starter) 的 `pro-admin` 子项目

## 创建项目

克隆代码，然后根据需要调整项目与代码

```
git clone https://github.com/deepraining/beego-starter.git yourProName --depth=1

cd yourProName
```

去掉原有的 Git 信息，并重新初始化

```
rm -rf .git

git init
```

创建数据库与表结构（可以自行修改）

```
# 本地环境：默认 flyway.conf 配置文件
flyway migrate

# 线上环境
flyway migrate -configFiles=flyway-prod.conf

# 如果你需要配置更多的环境，可以自己添加
```

## 运行项目

执行本地开发调试 `bee run` 命令，然后在浏览器中打开 `http://127.0.0.1:8080`

## 部署项目

在服务器上，找个合适的地方创建 `serverDirName` 目录

把本地 `beego-starter, bin/run.sh, bin/select.sh` 上传到 `serverDirName` 目录，并按实际需要修改 `run.sh, select.sh` 中 `SERVER_ENV` 与 `BIN_NAME` 变量的值

```
- serverDirName/
  - run.sh              # 运行、停止、重启、查看程序
  - select.sh           # 运行最新版本的二进制文件
  - beego-starter       # 二进制程序
  ...
```

```
cd serverDirName

sh run.sh start        # 运行程序
sh run.sh stop         # 停止程序
sh run.sh restart      # 重启程序
sh run.sh status       # 查看程序状态
sh run.sh version      # 查看程序版本
```

其他的具体部署请参考 Beego 官方文档。

## 前端配套项目

- [sbs-admin-web](https://github.com/deepraining/sbs-admin-web)
