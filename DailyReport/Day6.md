# Day 6

## 今天做的事

1. 早上起来又测试了一下昨天晚上写的，发现了很大的问题
2. 上午又修修补补一上午，重新看了一下`beego.orm`中的原生`SQL`查询返回结果给出的方法
3. 给数据库新增一列`ALTER TABLE <qq+qqnum>  ADD COLUMN start_time  VARCHAR(100) NOT NULL;`
4. 成功通过timer和ticker的组合完成了按时按要求提醒，美中不足的是缺少一些可靠性校验
5. 通过一番调试了解到了time库关于duration类型的一些内容，时区/毫秒什么的真难搞
6. 成功上线了层次二以及进阶内容，算是完美完成了前两个层次
7. 临睡觉之前又看了看关于CQ码的东西，感觉现在还有一点模糊不清的就是如何将上传的图片打包下载，明天进行一些接收图片的实践吧，力求完成部分层次三的内容

## 学到的东西

[CQ码](https://richardchien.gitee.io/coolq-http-api/docs/4.12/#/CQCode)

[如何接收图片](https://cqp.cc/forum.php?mod=viewthread&tid=25079&page=1#pid1131927)

[GCTT - Go 中文翻译组](https://studygolang.com/subject/1)

[Go 系列教程（Golang tutorial series）](https://studygolang.com/subject/2)

[golang的时区和神奇的time.Parse](https://www.jianshu.com/p/f809b06144f7)

[golang time.Duration() 问题](https://blog.csdn.net/yangxiaodong88/article/details/96299831)

[golang包time用法详解](https://blog.csdn.net/wschq/article/details/80114036)



## 还没想明白的事情

1. 怎么能在`beego.orm`中让一个结构体对应到多个表上呢
   1. 这个到现在还没解决，不过可以曲线救国，通过其中的原生SQL语句进行操作
2. 自己使用`orm.Raw`的时候，为什么问好位置总填不对呢，还得重新组合`string`输入，好麻烦
3. 如何加入完备的可靠性校验？这个问题可以等到实习结束之后再学学看看了，现在是没时间处理了