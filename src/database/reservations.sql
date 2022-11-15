drop table if exists reservations;
create table reservations
(
	id serial primary key,
	user_id int not null,
	service_id int not null, 
	amount float not null
);