create table users
(
    id uuid default gen_random_uuid(),
    name varchar(255),
    email varchar(255),
    phone varchar(250),
    password_hash varchar(255)
);

create unique index users_id_uindex
	on users (id);

alter table users
    add constraint users_pk
        primary key (id);



create table user_module
(
    id uuid default gen_random_uuid(),
    user_id uuid not null,
    module_name varchar(50) not null,
    module_desc varchar(500)
);

create unique index user_module_id_uindex
	on user_module (id);

alter table user_module
    add constraint user_module_pk
        primary key (id);



