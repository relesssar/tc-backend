create table ad_log
(
    id uuid default gen_random_uuid(),
    ad_login varchar(300) not null,
    date_time timestamp(0),
    eventid int not null,
    ip varchar(20),
    str varchar(500) not null
);

comment on table ad_log is 'Логи acrive directory';

comment on column ad_log.date_time is 'дата и время события';

comment on column ad_log.eventid is 'windows Айди события';

comment on column ad_log.ip is 'айпи отдельно, если есть';

comment on column ad_log.str is 'текст лога';

create unique index ad_log_id_uindex
	on ad_log (id);

alter table ad_log
    add constraint ad_log_pk
        primary key (id);

