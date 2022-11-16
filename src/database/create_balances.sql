drop table if exists balances;
create table balances
(
    id serial primary key,
    user_id int unique not null,
    amount float not null default 0
);