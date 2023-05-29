create table public.go_user
(
    name varchar,
    "Id" bigint generated always as identity
        constraint go_user_pk
            primary key
);

alter table public.go_user
    owner to postgres;

