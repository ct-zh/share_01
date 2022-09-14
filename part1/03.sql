drop table int_test;
create table int_test
(
    id    int auto_increment,
    col_1 int(1),
    col_2 int(20),
    primary key (id)
);

insert into int_test values(1,1,1),(10,10,10);

select * from int_test;

-- -------

show create table int_test;

alter table int_test modify col_2 int(20) zerofill;

select * from int_test;

