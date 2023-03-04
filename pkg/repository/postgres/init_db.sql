-------------------------------------------------
-- Create table users
-------------------------------------------------

create table if not exists public.user
(
    id           varchar                       primary key,
    login        varchar                     not null unique,
    password     varchar                     not null,
    email        varchar                     not null unique,
    phone_number varchar                     not null unique,
    created      timestamp without time zone not null default now()
);