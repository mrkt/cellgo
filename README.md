# CellGo
==========
 ![image](https://raw.githubusercontent.com/mrkt/cellgo/master/pic/logo.png)
细胞/CellGo框架是一款使用go语言开发的框架 （轻量级的Go Framework) 版本0.23

[![Build Status](https://drone.io/github.com/mrkt/cellgo/status.png?r)](https://drone.io/github.com/mrkt/cellgo/latest)
[![Go Walker](http://gowalker.org/api/v1/badge?r)](http://gowalker.org/github.com/mrkt/cellgo)

Features
--------
* MVC模式、代码清晰,操作简单,功能齐全
* 多数据源存储,轻松切换数据源
* Dao层数据驱动,对数据进行抽象
* 异步通讯,支持异步监控/推送
* 线程路由- http/socket通信

Demo
------
 ![image](https://raw.githubusercontent.com/mrkt/cellgo/master/pic/demo.png)
 ![image](https://raw.githubusercontent.com/mrkt/cellgo/master/pic/demo2.png)

Installation
------------

安装CellGo使用"go get"命令
    
    go get github.com/mrkt/cellgo
    
基于[Go](http://golang.org/)标准库


Update


更新CellGo使用"go get -u"命令

    go get -u github.com/mrkt/cellgo


Usage
------
用法请参考[example](https://github.com/mrkt/cellgo/tree/master/example)

测试访问地址: 
* http://localhost/
* http://localhost/?c=user&a=run
* http://localhost/?c=user&a=add&username=tommy.jin&email=tommy.jin@aliyun.com

修改conf中的 [IsUri] 为true开启静态路由
* http://localhost/
* http://localhost/user/run
* http://localhost/user/add/username/tommy.jin/email/tommy.jin@aliyun.com


