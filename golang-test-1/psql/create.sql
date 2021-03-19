create database test;

\c test;

create table if not exists categories(
    id serial primary key,
    description varchar(100) not null
);

create table if not exists products(
    id bigserial primary key,
    name varchar(100) not null,
    price real not null,
    quantity integer default 0,
    amount real default 0.0,
    category bigint not null,
    constraint products_categories_pk foreign key(category)
    references categories(id)
);

create table if not exists users(
    id bigserial primary key,
    firstname varchar(15) not null,
    lastname varchar(20) not null,
    email varchar(40) not null unique,
    password varchar(100) not null,
    status char(1) default '0'
);