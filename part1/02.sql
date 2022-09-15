-- 自增id的坑 2

drop table t1;
create table t1
(
    id    int not null auto_increment,
    name  varchar(10) unique,
    count int default 0,
    primary key (id),
    index (name)
) engine = innodb;

insert into t1(id, name) values (1, "aaa");

insert into t1(id, name)
values (111, "111"),
       (NULL, "abc"),
       (222, "222"),
       (NULL, "xyz");

select * from t1;

insert into t1(name)values("aaa"),("bbb"),("ccc") on duplicate key update count=100;
