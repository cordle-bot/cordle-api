create database cordle;
\c cordle;

create table public.users
(
    id     varchar(50)         not null
        constraint id
            primary key,
    wins   integer default 0   not null,
    losses integer default 0   not null,
    draws  integer default 0   not null,
    elo    integer default 500 not null
);

alter table public.users
    owner to "user";
