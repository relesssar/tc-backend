alter table problem_router_onyma_speeds
    add problem_status_old varchar(50) default null;

comment on column problem_router_onyma_speeds.problem_status_old is 'предыдущий статус проблемной записи';

