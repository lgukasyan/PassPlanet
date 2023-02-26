/* User Table */
create table users (
  user_id serial primary key,
  name varchar(255) not null,
  lastname varchar(255) not null,
  email varchar(255) not null,
  password char(60)
);



