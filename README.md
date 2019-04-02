jkgo
====
这是一个综合学习记录的代码存储区域，不作为任务独立项目

> 由于一开始的学习，整个结构还有些混乱，正逐步调整中

目录结构
```
etc - 相关的配置文件
html - 用于 http 服务器的 web 端
src
  - bveth 老公司的一个实验功能，已无用，保留用于记录学习
  - demo 简单功能的调试，将逐步取消
  - helper 一些通用小功能集，将逐步整体到 jkbase 中
  - jkbase 基础功能集
  - jkdbs 基础数据库操作功能库
  - jk 带有业务功能的接口集，其中部分业务的将逐步移入 jkbase
  - simplerserver 简洁服务器
  - jkencoderbackground 图片上添加文字
  - jkencoderimg 图片压缩功能
  - jkhttpclient http 客户端功能
  - jkhttpserver http 文件服务器功能
  - jkmisc 杂项小功能
  - jkservice 后台服务程序集合，可能会在未来提取出去
  - jkssh ssh 连接工具
  - jkstock 分析股市，实验阶段，未使用
  - jktime 时间功具
  - jktransfer pc 间传输文件工具
  - morserver monitor 服务器功能，监控，探测周边设备，当然要设备实现对应协议
tools - 一些下载，编译相关的脚本工具
  - build_simpleserver.bat simpleserver  的编译特殊，用了脚本实现
  - copy_simpleserver.bat 复制，打包的操作
  - get-goproj.bat 获取相关软件包的过程
  - get-goproj.sh 获取相关软件包的 linux 脚本，将移入 makefile 中
  - install-go-proj.bat go install 所有的包，其实不需要，将移除
Makefile - linux 下的编译 make help
```

#### jk/jkservice/jkstatus

```
http://xxx:12307/st/cmd=getstatus
```

## golang mqtt
Test on this `https://github.com/jeffallen/mqtt.git`

