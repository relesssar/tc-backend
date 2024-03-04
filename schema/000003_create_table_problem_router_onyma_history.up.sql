create table problem_router_onyma_history
(
    id uuid default gen_random_uuid() not null,
    problem_router_onyma_speed_id uuid not null,
    user_id uuid,
    old_val varchar(300),
    new_val varchar(300),
    msg varchar(300),
    created_at timestamp(0) default now()

);

comment on table problem_router_onyma_history is 'сообщения пользователей, история изменений по проблемным записям';

comment on column problem_router_onyma_history.created_at is 'Дата внесения записи';

comment on column problem_router_onyma_history.old_val is 'старое значение';

comment on column problem_router_onyma_history.new_val is 'новое значение';

comment on column problem_router_onyma_history.msg is 'описание, сообщение от пользователя';

create unique index problem_router_onyma_history_id_uindex
	on problem_router_onyma_history (id);

alter table problem_router_onyma_history
    add constraint problem_router_onyma_history_pk
        primary key (id);


create index problem_router_onyma_history_created_at_index
	on problem_router_onyma_history (created_at);

create index problem_router_onyma_history_user_id_index
	on problem_router_onyma_history (user_id);

create index problem_router_onyma_speed_id_index
	on problem_router_onyma_history (problem_router_onyma_speed_id);

alter table problem_router_onyma_speeds
    add client_status_onyma int default 0;

comment on column problem_router_onyma_speeds.client_status_onyma is 'Статус клиента в Ониме';

