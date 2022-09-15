-- 最左匹配原则
drop table multi_col_index_2;
create table multi_col_index_2
(
    class varchar(20) not null,
    name  varchar(20) not null,
    id    int not null,
    unique key (`id`)
) engine = innodb;

insert into multi_col_index_2(class, name,  id) VALUES ("1班", "刘德华", 201);
insert into multi_col_index_2(class, name,  id) VALUES ("1班", "张学友", 202);
insert into multi_col_index_2(class, name,  id) VALUES ("1班", "黎明", 203);
insert into multi_col_index_2(class, name,  id) VALUES ("1班", "郭富城", 204);

insert into multi_col_index_2(class, name,  id) VALUES ("3班", "张三", 101);
insert into multi_col_index_2(class, name,  id) VALUES ("3班", "李四", 102);
insert into multi_col_index_2(class, name,  id) VALUES ("3班", "李五", 103);
insert into multi_col_index_2(class, name,  id) VALUES ("3班", "王五", 104);
insert into multi_col_index_2(class, name,  id) VALUES ("3班", "赵六", 105);

insert into multi_col_index_2(class, name,  id) VALUES ("4班", "田七", 11);
insert into multi_col_index_2(class, name,  id) VALUES ("4班", "何八", 12);
insert into multi_col_index_2(class, name,  id) VALUES ("4班", "刘九", 14);
insert into multi_col_index_2(class, name,  id) VALUES ("4班", "宋十", 15);
insert into multi_col_index_2(class, name,  id) VALUES ("4班", "郭德纲", 18);

insert into multi_col_index_2(class, name,  id) VALUES ("5班", "罗志祥", 51);
insert into multi_col_index_2(class, name,  id) VALUES ("5班", "宋小宝", 52);
insert into multi_col_index_2(class, name,  id) VALUES ("5班", "尼古拉斯", 54);
insert into multi_col_index_2(class, name,  id) VALUES ("5班", "斯凯奇", 55);
insert into multi_col_index_2(class, name,  id) VALUES ("5班", "于谦", 58);

insert into multi_col_index_2(class, name,  id) VALUES ("6班", "小李", 31);
insert into multi_col_index_2(class, name,  id) VALUES ("6班", "小高", 32);
insert into multi_col_index_2(class, name,  id) VALUES ("6班", "东方不败", 34);
insert into multi_col_index_2(class, name,  id) VALUES ("6班", "令狐冲", 35);
insert into multi_col_index_2(class, name,  id) VALUES ("6班", "岳不群", 38);

select class, name,  id from multi_col_index_2 order by class asc, name asc;