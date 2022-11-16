drop table if exists accounting;
create table accounting(
    id serial primary key,
	service_id int not null,
	completion_date date not null,
	amount float not null default 0
);