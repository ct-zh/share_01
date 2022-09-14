-- 创建表格
drop table incr_row_test
(
    id   int not null auto_increment,
    name varchar(10) UNIQUE,
    primary key (id)
)ENGINE=INNODB;

-- step1
INSERT INTO incr_row_test values(1, "aaa");
INSERT INTO incr_row_test(name) values ("bbb"), ("ccc"), ("ddd");
delete from incr_row_test;

-- step2
insert into incr_row_test values (0, "eee");
insert into incr_row_test values (1, "fff");
INSERT INTO incr_row_test(name) values ("ggg");

-- result
select * from incr_row_test;
