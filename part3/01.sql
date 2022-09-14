
create table student_test(
    id int auto_increment,
    name varchar(20),
    class varchar(8),
    age tinyint,
primary key (id),
key `idx_class_name_age` (`class`, `name`, `age`)
)engine=innodb;

