alter table ad_log
    add full_name varchar(500);

comment on column ad_log.full_name is 'ФИО';

alter table ad_log
    add department varchar(500);

comment on column ad_log.department is 'Департамент';

