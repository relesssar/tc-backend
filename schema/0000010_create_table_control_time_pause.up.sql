create table control_time_pause
(
    id uuid default gen_random_uuid() not null,
    router_onyma_speed_id uuid not null,
    control_status smallint default 1 not null,
    created_at timestamp(0) default now()
);

comment on table control_time_pause is 'Контроль временного отключения.Сюда попадают записи если Роутер-пауза а Онима акивная. Или Роутер активный а Онима пауза.';

comment on column control_time_pause.router_onyma_speed_id is 'Данные с роутера и онимы';
comment on column control_time_pause.control_status is 'Статус инцидента.1-проблема.0-нет проблем';
comment on column control_time_pause.created_at is 'Дата внесения записи';

create unique index control_time_pause_id_uindex
	on control_time_pause (id);

alter table control_time_pause
    add constraint control_time_pause_pk
        primary key (id);





create table control_time_pause_history
(
    id uuid default gen_random_uuid() not null,
    control_time_pause_id uuid not null,
    user_id uuid not null,
    msg text not null,
    created_at timestamp(0)  default now()
);

comment on table control_time_pause_history is 'история сообщений по инцеденту';

comment on column control_time_pause_history.control_time_pause_id is 'айди инцедента ';

comment on column control_time_pause_history.user_id is 'Айди пользователя который оставил комментарий';

comment on column control_time_pause_history.msg is 'текст сообщения';

comment on column control_time_pause_history.created_at is 'время создания записи';

create unique index control_time_pause_history_id_uindex
	on control_time_pause_history (id);

alter table control_time_pause_history
    add constraint control_time_pause_history_pk
        primary key (id);

