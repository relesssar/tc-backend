create table access_group_history
(
    id uuid default gen_random_uuid() not null,
    access_group_id uuid not null,
    user_id uuid,
    old_val varchar(300),
    new_val varchar(300),
    msg varchar(300),
    created_at timestamp(0) default now()

);

comment on table access_group_history is 'сообщения пользователей, история изменений по записям';

comment on column access_group_history.created_at is 'Дата внесения записи';

comment on column access_group_history.old_val is 'старое значение';

comment on column access_group_history.new_val is 'новое значение';

comment on column access_group_history.msg is 'описание, сообщение от пользователя';

create unique index access_group_history_id_uindex
	on access_group_history (id);

alter table access_group_history
    add constraint access_group_history_pk
        primary key (id);



alter table access_groups
    add access_status varchar(25);

comment on column access_groups.access_status is 'Статус строки, в обратотке, закрыто';

alter table access_groups
    add updated_at timestamp(0) default NULL::timestamp without time zone;


alter table router_onyma_speeds
    add iface_shutdown_router int default 0 not null;

comment on column router_onyma_speeds.iface_shutdown_router is 'На роутере интерфейс 1-shutdown,0-включен';


alter table router_onyma_speeds
    add client_status_onyma int default 0;

comment on column router_onyma_speeds.client_status_onyma is 'Статус клиента в Ониме';


