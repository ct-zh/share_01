-- 联合索引长啥样？
drop table multi_col_index;
create table multi_col_index
(
    class varchar(20) not null,
    name  varchar(20) not null,
    age   tinyint not null,
    id    int not null,
    key `idx_class_name_age` (`class`, `name`, `age`),
    unique key (`id`)
) engine = innodb;

insert into multi_col_index(class, name, age, id) VALUES ("3班", "张三", 7, 101);
insert into multi_col_index(class, name, age, id) VALUES ("3班", "李四", 7, 102);
insert into multi_col_index(class, name, age, id) VALUES ("3班", "李五", 8, 103);
insert into multi_col_index(class, name, age, id) VALUES ("3班", "王五", 9, 104);
insert into multi_col_index(class, name, age, id) VALUES ("3班", "赵六", 8, 105);

insert into multi_col_index(class, name, age, id) VALUES ("4班", "田七", 7, 11);
insert into multi_col_index(class, name, age, id) VALUES ("4班", "何八", 7, 12);
insert into multi_col_index(class, name, age, id) VALUES ("4班", "刘九", 8, 14);
insert into multi_col_index(class, name, age, id) VALUES ("4班", "宋十", 9, 15);
insert into multi_col_index(class, name, age, id) VALUES ("4班", "郭德纲", 8, 18);

insert into multi_col_index(class, name, age, id) VALUES ("5班", "罗志祥", 7, 51);
insert into multi_col_index(class, name, age, id) VALUES ("5班", "宋小宝", 7, 52);
insert into multi_col_index(class, name, age, id) VALUES ("5班", "尼古拉斯", 8, 54);
insert into multi_col_index(class, name, age, id) VALUES ("5班", "斯凯奇", 9, 55);
insert into multi_col_index(class, name, age, id) VALUES ("5班", "于谦", 8, 58);

insert into multi_col_index(class, name, age, id) VALUES ("6班", "小高", 7, 31);
insert into multi_col_index(class, name, age, id) VALUES ("6班", "小李", 7, 32);
insert into multi_col_index(class, name, age, id) VALUES ("6班", "东方不败", 8, 34);
insert into multi_col_index(class, name, age, id) VALUES ("6班", "令狐冲", 9, 35);
insert into multi_col_index(class, name, age, id) VALUES ("6班", "岳不群", 8, 38);

select class, name, age, id from multi_col_index order by class asc, name asc, age asc;