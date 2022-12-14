# 9.15分享
- 标题： 工作中遇到的MySQL问题 探讨与分享
- 声明：MySQL版本 5.7.2；引擎只使用Innodb；不考虑分库分表；不考虑分布式事务；仅讨论单机单表的情况；
- 查看DB文件位置：`show global variables like "%datadir%"`

## part 1 开胃菜
### int(11) 中的11代表什么？
- 扩展1：自增int超出了范围会怎样？
- 扩展2：主键int按照顺序递增，与不按照顺序递增的区别；联系一下雪花算法；

### 自增id的坑，你踩过吗？


## part 2
### 该用什么类型存货币
- 补充：unsigned int 导致的bug

### 该用什么类型存手机号
- 补充：手机号查询导致P0事故

### 加字段的时候，到底要不要NOT NULL?
- 系统历史包袱，很多为null的字段，应该如何避免踩坑？

### 待定：IP地址的存储

## part 3 索引
### 为什么低区分度的字段不适合做索引？
- 关键字：主键索引结构（聚簇索引）；二级索引结构
- MySQL建立索引的流程；索引的具体底层结构；
- 建立索引的几大原则
- 最左匹配
- todo 补充figma动画，索引的建立流程；
- todo 补充 golang实现的b+树，声明b+树的时间复杂度；

### 下面哪些查询能命中索引？
- 关键字：主键索引查询流程；页分裂；二级索引查询流程；回表查询；
- MySQL查询使用索引的具体流程；
- todo 补充figma动画，查询的流程；
- todo 补充golang实现的b+树，页分裂的代码；

### 案例
- 案例1：不合适的索引导致扫表
- 案例2：回表性能问题
- 案例3：几个会导致扫表的查询（连接查询，连接字符串的字符集不同导致扫表）
- 案例4：低区分度的字段真的不适合建立索引吗？
- 案例5：大数据表的分页查询
- 案例6：优化查询，不一定要盯着索引来做

## part 4 锁
### InnoDB行锁（数据库的隔离级别）
### 遇到死锁应该如何处理
