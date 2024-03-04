create table access_groups
(
    id uuid default gen_random_uuid() not null,
    router_name varchar(15),
    iface_host varchar(15),
    ip varchar(15),
    iface_name varchar(200),
    iface_desc varchar(500),
    client_status smallint,
    in_policy varchar(50),
    out_policy varchar(50),
    access_group varchar(250),
    dognum char(10),
    clsrv char(10),
        created_at timestamp(0) default now()
);

comment on table access_groups is 'список интерфейсов не имеющих access-group';

comment on column access_groups.router_name is 'имя маршрутизатора';

comment on column access_groups.iface_host is 'Сеть';

comment on column access_groups.ip is 'Айпи интерфейса';

comment on column access_groups.iface_name is 'название интерфейса';

comment on column access_groups.iface_desc is 'описание интерфейса';

comment on column access_groups.client_status is 'статус клиента, 0-вкл., 1-выкл.';

comment on column access_groups.in_policy is 'Политика входящей скорости';

comment on column access_groups.out_policy is 'Политика исходящей скорости';

comment on column access_groups.dognum is 'Лицевой счёт клиента в Ониме';

comment on column access_groups.clsrv is 'айди подключения в Ониме';

comment on column access_groups.created_at is 'дата внесения записи';

create unique index access_groups_id_uindex
	on access_groups (id);

alter table access_groups
    add constraint access_groups_pk
        primary key (id);

