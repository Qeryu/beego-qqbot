# Day 3

## 今天做的事情

1. 上午经yh大佬的建议，通过启动盘试图修复昨天晚上被破坏的windows的引导文件，但是按照教程操作之后，还是没有反应，甚至连那个grub的选择系统的界面都没了，算了算了，先在`ubuntu`上弄了，就是没有`GoLand`有点难受，想下一个，奈何速度太慢。
2. 刚写完上一条就被啪啪打脸，飞速下载完成`GoLand`
3. 想清楚了层次一的大体思路，接收`POST`，分析`body`中的信息，将信息转存至`struct`，根据`usr_id`经过`API`进行回复
4. 需要做的事情就是，写一个`POST`的函数，设计一个存放信息的`struct`，写一个回复调用`api`的函数
5. 写了，但是上线之后没有结果，明天再改吧
6. 东西放到了另外一个仓库里，因为要放到不同的文件夹，现在也不知道一个仓库怎么操作，就先这样弄着了
7. https://github.com/Uniqsy/beego-qqbot.git

## 今天学到的东西

1. `dockerfile`的使用
2. `linux`中环境变量的配置
3. `golang`中`struct/json`的相互转换（`net/http`
4. `json`格式的具体要求
5. 本地测试工具`postman`的使用

### 今天学习使用`beego`框架进行操作

从快速指南到开发文档

感觉这种框架使用完全是从零开始学期，看到的每一个词几乎都是看不懂的东西，都需要去查

今天了解了`Controller`/`Router`等等的东西，看到的例子都试了试，现在可以算是初步了解GET和POST的处理方法了

### 其他东西

[linux系统中修改环境变量以及令修改生效的方法](https://www.cnblogs.com/franson-2016/p/6063226.html)

[golang中struct、json、map互相转化](https://blog.csdn.net/xiaoquantouer/article/details/80233177)

[Beego控制器发送GET/POST请求并获取返回信息](https://blog.csdn.net/qq_38280150/article/details/102841773?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task)

[docker 中文手册](https://m.php.cn/manual/view/35252.html)

[dockerfile的使用](https://www.cnblogs.com/edisonchou/p/dockerfile_inside_introduction.html)



## 还未能解决的问题

如何能在本地方便地调试代码？除了在本地布置一个一模一样的环境之外？

是否能用一个仓库布置两处的东西，分别取出一个仓库中不同的东西放到不同的本地仓库中？

为什么`source`命令可以应用配置的`path`，还能怎么用？

~~容器中的环境变量如果退出容器后就失效了，怎么处理？~~

什么是典型的`MVC框架`？

`linux`软连接是什么，具体参数是什么？

今天晚上写的这些到底哪里出错了啊啊啊啊啊，要崩溃了