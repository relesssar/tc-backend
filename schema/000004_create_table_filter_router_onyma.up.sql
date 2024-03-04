create table filter_router_onyma
(
    id uuid default gen_random_uuid() not null
        constraint filter_router_onyma_pk
            primary key,
    filter_type varchar(32) not null,
    filter_val varchar(250) not null,
    filter_desc varchar(500),
    user_id uuid,
    created_at timestamp(0) default now()
);

comment on table filter_router_onyma is 'Фильтр исключения, всё это исключается из списка';

comment on column filter_router_onyma.filter_type is 'тип фильтра,(interface_name)';

alter table filter_router_onyma owner to tl_user_db_2021;


create unique index if not exists filter_router_onyma_id_uindex
	on filter_router_onyma (id);


alter table filter_router_onyma
    add router_name varchar(15);

comment on column filter_router_onyma.router_name is 'имя роутера';


create unique index filter_val_user_id_uindex
    on filter_router_onyma (filter_val, user_id);