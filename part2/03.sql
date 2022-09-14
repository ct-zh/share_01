create table int_test_2 (col_1 int unsigned, col_2 int unsigned);

insert into int_test_2 values(1,2);

select col_1 - col_2 from int_test_2;

SET sql_mode = 'NO_UNSIGNED_SUBTRACTION';