# Day 4

## 今天做的事

1. 下午完成了层次一，可算是能回复一些东西了
2. 晚上想了想如何进行层次二，应该是
   1. 要用一个数据库来存储`DDL`信息
   2. 每个时段了解一下是否有本时段需要提醒的`DDL`
   3. 对需要提醒的`DDL`，从数据库中抽取提醒内容，提醒对象
   4. 通过`POST`发消息进行提醒
3. 晚上的时候为了方便测试，又加了一个加好友的功能，可以根据验证信息进行选择加好友，但是我没有设置，现在是随便加。
4. 早点睡觉，前几天学东西太肝了，今天一天都持续耳鸣，效率低

## 学到的东西

基本语句`http`中关于报文/post请求/请求的响应等内容（从《图解HTTP上学的》

`SQL`基本语句

### 其他

[net/http中json的使用方法](https://blog.csdn.net/wangshubo1989/article/details/70245570)

[MySQL教学——廖雪峰](https://www.liaoxuefeng.com/wiki/1177760294764384)

[mysql出现ERROR1698(28000):Access denied for user root@localhost错误解决方法](https://www.cnblogs.com/cpl9412290130/p/9583868.html)

[Go如何响应http请求？](https://juejin.im/post/5ca0a2256fb9a05e6f7af992#heading-9)

[理解 Go interface 的 5 个关键点](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)

## 未解决问题

如何在`docker`中将`bot容器`和`MySQL容器`关联到一起？

如何将之前发现的`SDK`纳入`beego`环境？