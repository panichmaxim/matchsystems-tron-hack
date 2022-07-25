drop table if exists access_request;
create table if not exists access_requests
(
    id         bigserial
        primary key,
    user_id    bigint                                             not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);
