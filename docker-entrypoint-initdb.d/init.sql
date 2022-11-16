grant pg_read_all_data to avito;
grant pg_write_all_data to avito;

drop table if exists balances;
create table balances
(
    id serial primary key,
    user_id int unique not null,
    amount float not null default 0
);

drop table if exists reservations;
create table reservations
(
    id serial primary key,
    user_id int not null,
    service_id int not null,
    order_id int not null unique,
    amount float not null
);

drop table if exists accounting;
create table accounting(
    id serial primary key,
	service_id int not null,
	completion_date date not null,
	amount float not null default 0
);