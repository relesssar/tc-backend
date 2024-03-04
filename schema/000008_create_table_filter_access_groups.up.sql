create table filter_access_groups
(
    id uuid default gen_random_uuid() not null
        constraint filter_access_groups_pk
            primary key,
    filter_type varchar(32) not null,
    filter_val varchar(250) not null,
    filter_desc varchar(500),
    user_id uuid,
    created_at timestamp(0) default now(),
    router_name varchar(15) not null

);

comment on table filter_access_groups is 'Фильтр исключения, всё это исключается из списка';

comment on column filter_access_groups.filter_type is 'тип фильтра,(interface_name, ip)';
comment on column filter_access_groups.router_name is 'имя роутера';

alter table filter_access_groups owner to tl_user_db_2021;


create unique index if not exists filter_access_groups_id_uindex
on filter_access_groups (id);

create unique index filter_access_groups_filter_val_user_id_uindex
    on filter_access_groups (filter_val, user_id);


