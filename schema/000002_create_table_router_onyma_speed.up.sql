create table router_onyma_speeds
(
    id uuid default gen_random_uuid(),
    branch_service varchar(15),
    router_name varchar(15),
    interface_name varchar(50),
    interface_description varchar(500),
    in_policy_router varchar(50),
    in_speed_router varchar(50),
    out_policy_router varchar(50),
    out_speed_router varchar(50),
    ip_interface varchar(20),
    branch_contract varchar(50),
    dognum varchar(15),
    clsrv varchar(15),
    company_name varchar(300),
    in_speed_onyma varchar(50),
    out_speed_onyma varchar(50),
    insert_datetime timestamp(0)  default now()
);
alter table router_onyma_speeds alter column id set not null;

create unique index router_onyma_speeds_id_uindex
	on router_onyma_speeds (id);

alter table router_onyma_speeds
    add constraint router_onyma_speeds_pk
        primary key (id);



comment on table router_onyma_speeds is 'Данные с роутеров и с онимы';

comment on column router_onyma_speeds.branch_service is 'Филиал Общества, в котором предоставляется услуга';

comment on column router_onyma_speeds.router_name is 'Имя маршрутизатора на сети ПД, в соответствии с КТС-РИ-ДЭ-038';


comment on column router_onyma_speeds.interface_name is 'Имя интерфейса на маршрутизаторе на котором прописана услуга (Internet, VPN)';

comment on column router_onyma_speeds.interface_description is 'Описание названия клиента на интерфейса маршрутизатора на котором прописана услуга (Internet, VPN)';

comment on column router_onyma_speeds.in_policy_router is 'Название Input-политики (входящая скорость) на роутере';

comment on column router_onyma_speeds.in_speed_router is 'Входящая скорость на интерфейсе в бит/с. Значение Input Policy (gbps * 1000000000, mbps * 1000000, kbps * 1000, bps * 1)';

comment on column router_onyma_speeds.out_policy_router is 'Название Output-политики (исходящая скорость) на интерфейсе';

comment on column router_onyma_speeds.out_speed_router is 'Исходящая скорость на интерфейсе в бит/с. Значение Output Policy (gbps * 1000000000, mbps * 1000000, kbps * 1000, bps * 1)';

comment on column router_onyma_speeds.ip_interface is 'IP-адрес назначенный на интерфейсе маршрутизатора';

comment on column router_onyma_speeds.branch_contract is 'Филиал Общества, в котором заключен договор на предоставление услуги';

comment on column router_onyma_speeds.dognum is 'Номер лицевого счета клиента в АСР `Onyma`';

comment on column router_onyma_speeds.clsrv is 'Номер (ID) подключения в АСР `Onyma`';

comment on column router_onyma_speeds.company_name is 'Наименование юридического лица в АСР `Onyma` которому предоставляется услуга и от которого высталвляются счета клиенту';

comment on column router_onyma_speeds.in_speed_onyma is 'Входящая скорость услуги - данные из АСР `Onyma`';

comment on column router_onyma_speeds.out_speed_onyma is 'Исходящая скорость услуги - данные из АСР `Onyma`';


create table problem_router_onyma_speeds
(
    id uuid default gen_random_uuid(),
    router_onyma_speed_id uuid NOT NULL,
    branch_service varchar(15),
    router_name varchar(15),
    interface_name varchar(50),
    interface_description varchar(500),
    in_policy_router varchar(50),
    in_speed_router varchar(50),
    out_policy_router varchar(50),
    out_speed_router varchar(50),
    ip_interface varchar(20),
    branch_contract varchar(50),
    dognum varchar(15),
    clsrv varchar(15),
    company_name varchar(300),
    in_speed_onyma varchar(50),
    out_speed_onyma varchar(50),
    problem_status varchar (50),
    insert_datetime timestamp(0)  default now(),
    updated_at timestamp(0) default NULL::timestamp without time zone
);

alter table problem_router_onyma_speeds alter column id set not null;

create unique index problem_router_onyma_speeds_id_uindex
	on problem_router_onyma_speeds (id);

alter table problem_router_onyma_speeds
    add constraint problem_router_onyma_speeds_pk
        primary key (id);



comment on table problem_router_onyma_speeds is 'Проблемные данные с роутеров и с онимы';

comment on column problem_router_onyma_speeds.branch_service is 'Филиал Общества, в котором предоставляется услуга';

comment on column problem_router_onyma_speeds.router_name is 'Имя маршрутизатора на сети ПД, в соответствии с КТС-РИ-ДЭ-038';


comment on column problem_router_onyma_speeds.interface_name is 'Имя интерфейса на маршрутизаторе на котором прописана услуга (Internet, VPN)';

comment on column problem_router_onyma_speeds.interface_description is 'Описание названия клиента на интерфейса маршрутизатора на котором прописана услуга (Internet, VPN)';

comment on column problem_router_onyma_speeds.in_policy_router is 'Название Input-политики (входящая скорость) на роутере';

comment on column problem_router_onyma_speeds.in_speed_router is 'Входящая скорость на интерфейсе в бит/с. Значение Input Policy (gbps * 1000000000, mbps * 1000000, kbps * 1000, bps * 1)';

comment on column problem_router_onyma_speeds.out_policy_router is 'Название Output-политики (исходящая скорость) на интерфейсе';

comment on column problem_router_onyma_speeds.out_speed_router is 'Исходящая скорость на интерфейсе в бит/с. Значение Output Policy (gbps * 1000000000, mbps * 1000000, kbps * 1000, bps * 1)';

comment on column problem_router_onyma_speeds.ip_interface is 'IP-адрес назначенный на интерфейсе маршрутизатора';

comment on column problem_router_onyma_speeds.branch_contract is 'Филиал Общества, в котором заключен договор на предоставление услуги';

comment on column problem_router_onyma_speeds.dognum is 'Номер лицевого счета клиента в АСР `Onyma`';

comment on column problem_router_onyma_speeds.clsrv is 'Номер (ID) подключения в АСР `Onyma`';

comment on column problem_router_onyma_speeds.company_name is 'Наименование юридического лица в АСР `Onyma` которому предоставляется услуга и от которого высталвляются счета клиенту';

comment on column problem_router_onyma_speeds.in_speed_onyma is 'Входящая скорость услуги - данные из АСР `Onyma`';

comment on column problem_router_onyma_speeds.out_speed_onyma is 'Исходящая скорость услуги - данные из АСР `Onyma`';
comment on column problem_router_onyma_speeds.problem_status is 'Статус проблемной записи';
comment on column problem_router_onyma_speeds.insert_datetime is 'Дата записи данных в таблицу';

