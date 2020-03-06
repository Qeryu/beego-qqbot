# Day5

## 今天做的事

1. 成功将昨天一股脑塞到一起的代码拆分，自己捋清楚了

2. 在服务器上布置了mysql，处理了一系列过程中遇到的问题

   1. ```sql
      ALTER TABLE users CHANGE COLUMN use_ddl_reminder use_ddl_reminder CHAR DEFAULT 'N';
      ```

   2. | Field            | Type             | Null | Key  | Default | Extra          |
      | ---------------- | ---------------- | ---- | ---- | ------- | -------------- |
      | id               | int(10) unsigned | NO   | PRI  | NULL    | auto_increment |
      | user_id          | bigint(20)       | YES  |      | NULL    |                |
      | nickname         | varchar(40)      | YES  |      | NULL    |                |
      | use_ddl_reminder | char(1)          | YES  |      | N       |                |

1. 熟练了昨天晚上速成的sql语言基础，综合廖雪峰/菜鸟教程又学了一些新的东西
2. 学习了beego框架下的orm内容，其中模型定义以及命令模式这些都卡了很久才看懂
3. 重新看了一边coolq-http-api和qq-bot-api的文档，感觉自己第一遍看的时候啥都不懂，有很多东西其实都能拿来用的，尤其是在层次一中，没用qq-bot-api，全部都是自己在beego框架下手写的内容，当时比较麻烦，现在一想，如果使用了，可以很快完成
4. 在服务器端配置好了qqbotapi
5. 将代码中某些东西比如

## 今天学到的东西

[go get golang.org/x 包失败解决方法](https://github.com/AlexWoo/doc/blob/master/GOLang/go get golang.org:x 包失败解决方法.md)

[docker笔记（五、docker安装mysql数据库](https://blog.csdn.net/liu540885284/article/details/101065407)

[docker笔记（六、docker将beego程序和mysql关联起来](https://blog.csdn.net/liu540885284/article/details/101072019)

[Git冲突：commit your changes or stash them before you can merge.](https://blog.csdn.net/lincyang/article/details/21519333)

[mysql设置主键从1开始自增](https://blog.csdn.net/qq_40576301/article/details/100522559?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task)

[package strconv](https://godoc.org/strconv)

[beego_orm](https://beego.me/docs/mvc/model/overview.md)

## 今天发现的问题

1. 在qqbotapi中是否存在接受事件上报的方法，如果有，怎么用
2. golang中的定时器是否可以拿来配合ddlReminder使用，如果可以，怎么用
3. 层次一所写内容还可以用别人写好的sdk来快速完成，可以增加代码可读性，但是貌似没时间翻工了
4. beego的rom