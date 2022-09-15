-- not null 的坑

drop table user;
create table user (
    id int,
    name varchar(20),
index(id)
)engine=innodb;

insert into user values(1,'aaa');
insert into user values(2,'bbb');
insert into user values(3,'ccc');

select * from user where id!=1;

insert into user(name) values('wangwu');

explain select * from user where id!=1;

explain select * from user where id!=1 or id is null;

EXPLAIN
select * from user where id is null
UNION
select * from user where id=1;






