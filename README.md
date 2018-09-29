## StandardMS电子规范管理系统


StandardMS电子规范管理系统是一款方便中小团队和工程师个人管理规范、计算书、图集的web应用.
纸质的规范不可能人人都拥有齐全，最多还是需要电子扫描的规范来查阅一些不经常用到的关键数据。一个团队中也不可能人人去拥有这么齐全的规范，所以，团队中找一台电脑运行这个应用即可，大家都向上传电子规范，人人都可以查阅了。
至于规范的有效性，采用另外一个数据库来管理，查阅过程中，自动对照有效版本数据库。
有效版本数据库的维护，一般定期更新一下，我也可以提供。QQ504284


##Quick Start
######Download and run

    Windows系统：解压压缩包到硬盘上，运行StandardMS-win64.exe或StandardMS-win32.exe，在浏览器中打开IP即可访问。本地IP即127.0.0.1。
    
    Linuxx系统：git clone https://github.com/3xxx/EngineerCMS.git
    
    go build main.go编译即可。

## 特性
* 采用golang语言，sqlite数据库，无需部署，直接运行
* 设计规范、计算书、图集等查阅
* 快捷批量上传
* 对报告中的规范名称进行批量核对


## Documentation

* [中文文档]——请查阅document文件夹

## 免费开源和问题反馈

* 开源地址[https://github.com/3xxx/standardms/](https://github.com/3xxx/standardms/)
* 问题反馈: [https://github.com/3xxx/standardms/issues](https://github.com/3xxx/standardms/issues)

## LICENSE

HydroCMS source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
