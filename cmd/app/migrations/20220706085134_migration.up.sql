create table billing
(
    id         bigserial
        primary key,
    user_id    bigint                                             not null,
    requests   bigint                                             not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);

create table billing_packet
(
    id         bigserial
        primary key,
    user_id    bigint                                             not null,
    requests   bigint                                             not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);

create table billing_request
(
    id         bigserial
        primary key,
    user_id    bigint                                             not null,
    query      varchar                                            not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);
