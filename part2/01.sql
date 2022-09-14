-- 不建议使用 VARCHAR 数据类型存储货币值。
-- 因为使用max等数学函数时会出现意料之外的情况
drop table varchar_test;
create table varchar_test (col_1 varchar(10));

insert into varchar_test values('0'),('1'),('2'),('9'),('10');

select max(col_1) from varchar_test;

