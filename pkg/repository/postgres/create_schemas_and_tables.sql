-------------------------------------------------
-- Create schema "chatty"
-- todo: (dont create schema - use public)
-- mirgations? goose
-------------------------------------------------
create schema if not exists chatty;

-------------------------------------------------
-- Create table chatty.users
-------------------------------------------------

//use varchar(27) for id
//created  timestamp without timezone
create table if not exists chatty.users
(
    id       bytea     not null primary key,
    name     varchar   not null unique,
    password varchar   not null,
    created  timestamp not null default now()
);

create unique index if not exists chatty_user_name_idx on chatty.users (name);
