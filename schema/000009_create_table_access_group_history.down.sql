drop table access_group_history;
alter table access_groups drop column access_status;
alter table access_groups drop column updated_at;
alter table router_onyma_speeds drop column iface_shutdown_router;
alter table router_onyma_speeds drop column client_status_onyma;