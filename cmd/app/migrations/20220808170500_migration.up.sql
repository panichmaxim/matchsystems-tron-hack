drop table if exists category_group;
drop table if exists category_groups;

alter table category
    add number integer default null;

update category set number = id where number is null;

alter table category
    alter column number set not null;
