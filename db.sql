-- Active: 1677423790558@@127.0.0.1@5432@passplanet
/* User Table */
create table users (
  user_id serial primary key,
  name varchar(255) not null,
  lastname varchar(255) not null,
  email varchar(255) not null,
  password char(60)
);

create table passwords (
  password_id serial primary key,
  user_id serial references users(user_id) not null,
  title varchar(50) not null,
  url varchar(255),
  icon_base64data text,
  description varchar(100),
  password char(60) not null
);