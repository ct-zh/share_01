-- 手机号

drop table phone_test;
create table phone_test(
    id int auto_increment,
    name varchar(20) not null,
    phone varchar(20) not null ,
    primary key (id),
   key(phone, name)
)ENGINE = innodb charset = utf8mb4;

insert into phone_test(name, phone) values ("张三", 13838381438);
insert into phone_test(name, phone) values ("李四", 14444444444);
insert into phone_test(name, phone) values ("王五", 15555555555);

select * from phone_test;
-- 猜猜这里会不会走索引
explain select name from phone_test where phone=14444444444;





-- phone 2
drop table phone_test_2;
create table phone_test_2(
   id int auto_increment,
   sale_name varchar(20) not null,
   phone varchar(20) not null ,
   primary key (id),
   key(phone, sale_name)
)ENGINE = innodb charset =utf8;

insert into phone_test_2(sale_name, phone) values ("张三的销售", 13838381438);
insert into phone_test_2(sale_name, phone) values ("李四的销售", 14444444444);
insert into phone_test_2(sale_name, phone) values ("王五的销售", 15555555555);

select * from phone_test_2;

explain select sale_name from phone_test_2 where phone=14444444444;

explain select name, sale_name from phone_test, phone_test_2 where phone_test.phone=phone_test_2.phone;



