-- 超大分页查询
drop table `limit_optimize_tbl`;
CREATE TABLE `limit_optimize_tbl` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `account` varchar(50) NOT NULL,
      `order_id` varchar(100) NOT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create index idx_account on limit_optimize_tbl(account);
create index idx_order_id on limit_optimize_tbl(order_id);

-- 构造一些数据
DELIMITER //
CREATE PROCEDURE limit_optimize_test()
BEGIN
    DECLARE i INT;
    SET i=1000000;
    WHILE i<=3000000 DO
            INSERT INTO limit_optimize_tbl(account,order_id) VALUES('test_123',concat('order', i));
            SET i=i+1;
        END WHILE;
END
//
DELIMITER ;

call limit_optimize_test();

explain select * from limit_optimize_tbl order by order_id limit 0,10;

explain select * from limit_optimize_tbl order by order_id limit 1000000,10;

explain select * from limit_optimize_tbl order by order_id limit 5660,10;

explain select * from limit_optimize_tbl order by order_id limit 5661,10;


