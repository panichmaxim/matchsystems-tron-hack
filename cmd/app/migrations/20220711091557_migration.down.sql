drop table if exists access_requests;
create table if not exists access_request
(
    id         uuid                                               not null
        primary key,
    user_id    bigint                                             not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);
