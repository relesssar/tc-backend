create table dppp_adress_base
(
    id uuid default gen_random_uuid(),
    real_ip varchar(20) not null,
    dppp_name varchar(500) not null
);

comment on table dppp_adress_base is 'Данные с базы ДППП';

create unique index dppp_adress_base_id_uindex
	on dppp_adress_base (id);

alter table dppp_adress_base
    add constraint dppp_adress_base_pk
        primary key (id);

alter table dppp_adress_base
    add created_at timestamp(0) default now();
