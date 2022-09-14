create table decimal_test(col_1 int, col2 decimal(5,1));

insert into decimal_test values(1,10.1),(2,10.11),(3,10.16);

select * from decimal_test;

-- todo 给个案例，测试decimal与bigint的数据类型的计算性能差距
-- 因此在合适的情况下，也可以考虑选用 BIGINT 的数据类型，它能同时避免浮点数计算不精确和 DECIMAL 计算代价高的问题，不过你也需要同时处理和小数点相关的问题。


