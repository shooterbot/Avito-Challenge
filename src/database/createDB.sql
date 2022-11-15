create role avito with password 'challenge';
grant all privileges on database avito_challenge to avito;
grant pg_read_all_data to avito;
grant pg_write_all_data to avito;
alter role avito with login;