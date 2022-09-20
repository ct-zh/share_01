-- 测试 select class, name from t_student where class > '4班' and name = '赵赵';
create table t_student
(
    id    int auto_increment,
    name  varchar(20) not null,
    class varchar(12) not null,
    age   tinyint,
    primary key (id),
    key (class, name)
) engine = innodb;

-- 问题1. 如何证明name='赵赵'做到了索引下沉？
-- 答：explain 的extra 字段 有 Using where condition ；
set OPTIMIZER_SWITCH = 'index_condition_pushdown=off';
set OPTIMIZER_SWITCH = 'index_condition_pushdown=on';

-- 问题2. 如果select增加非索引字段；还会有索引下沉吗？
-- 问题3. 如果where用到了非索引字段，还会有索引下沉吗？
explain select class,name from t_student where class > '19班' and name = '赵赵';



