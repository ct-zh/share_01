-- 次级索引的结构、为什么低区分度的字段不适合建立索引

create table index_text1
(
    class varchar(20) not null,
    id    int         not null,
    unique key (id)
) engine = innodb;

DELIMITER //
create procedure gen_class1()
BEGIN
    declare n int;
    set n = 1;
    while n <= 500
        do
            insert into index_text1(class, id) values ("三班", n);
            set n = n + 1;
        end while;
end
//
DELIMITER ;

DELIMITER //
create procedure gen_class2()
BEGIN
    declare n int;
    set n = 501;
    while n <= 1000
        do
            insert into index_text1(class, id) values ("四班", n);
            set n = n + 1;
        end while;
end
//
DELIMITER ;

call gen_class1();
call gen_class2();

select * from index_text1 limit 1000;

