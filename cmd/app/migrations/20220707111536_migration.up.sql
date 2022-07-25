create table billing_api_key
(
    id         bigserial
        primary key,
    api_key    varchar                                            not null
        unique,
    user_id    bigint                                             not null,
    billing_id bigint                                             not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);
