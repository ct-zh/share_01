
drop table messages;
create table messages(
    `id` bigint unsigned auto_increment,
    `uid` int unsigned,
    `status` tinyint default 0,
    `text` varchar(256),
     primary key (`id`),
     key(status)
)engine =innodb;

-- 构造一些数据
drop procedure gen_message_succ;
DELIMITER //
CREATE PROCEDURE gen_message_succ()
BEGIN
    DECLARE i INT;
    SET i=1;
    WHILE i<=10000000 DO
            INSERT INTO messages(uid, status, text) VALUES(i, 1, "hello world");
            SET i=i+1;
        END WHILE;
END
//
DELIMITER ;

drop procedure gen_message_fail;
DELIMITER //
CREATE PROCEDURE gen_message_fail()
BEGIN
    DECLARE i INT;
    SET i=1;
    WHILE i<=30 DO
            INSERT INTO messages(uid, status, text) VALUES(i, 2, "hello world fail");
            SET i=i+1;
        END WHILE;
END
//
DELIMITER ;

call gen_message_succ();
call gen_message_fail();

select status, count(*) from messages group by status;

-- 目标sql
select * from messages where status=2;