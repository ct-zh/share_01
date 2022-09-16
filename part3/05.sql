-- MySQL大数据量查询
-- 当存在主键 asc排序和 limit N N的值很小时，可能出现的bug

drop table `order_info`;

CREATE TABLE `order_info`(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `uid`          int(11) unsigned,
    `order_status` tinyint(3) DEFAULT NULL,
    `desc`         varchar(1024),
    PRIMARY KEY (`id`),
    KEY `idx_uid_stat` (`uid`, `order_status`),
     key `idx_desc` (`desc`(30))
) ENGINE = InnoDB;

-- 构造一些数据
drop procedure add_order_info;
DELIMITER //
CREATE PROCEDURE add_order_info()
BEGIN
    DECLARE i INT;
    SET i=1000000;
    WHILE i<=3000000 DO
            INSERT INTO order_info(uid, order_status) VALUES(i, mod(i, 3));
            SET i=i+1;
        END WHILE;
END
//
DELIMITER ;

call add_order_info();


explain select * from order_info where uid = 2999971 order by id asc limit 10;


SET optimizer_trace="enabled=on";        -- 打开 optimizer_trace
SELECT * FROM order_info where uid = 2154280 order by id asc limit 1;
SELECT * FROM information_schema.OPTIMIZER_TRACE;    -- 查看执行计划表
SET optimizer_trace="enabled=off"; -- 关闭 optimizer_trace
