drop table phone_test;
create table phone_test(
    id int auto_increment,
    name varchar(20) not null,
    phone varchar(20) not null ,
    primary key (id),
   key(phone, name)
)ENGINE = innodb;

insert into phone_test(name, phone) values ("张三", 13838381438);
insert into phone_test(name, phone) values ("李四", 14444444444);
insert into phone_test(name, phone) values ("王五", 15555555555);

select * from phone_test;

explain select name from phone_test where phone=14444444444;





