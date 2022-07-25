create table users
(
    id              bigserial
        primary key,
    name            varchar                                            not null,
    email           varchar
        unique,
    is_active       boolean                  default false             not null,
    password        varchar,
    token           varchar
        unique,
    permissions     character varying[],
    social_email    varchar,
    social_identity varchar,
    social_network  varchar,
    created_at      timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at      timestamp with time zone default CURRENT_TIMESTAMP not null
);

create table auth_token
(
    id                 uuid                     not null
        primary key,
    user_id            bigint                   not null,
    access_token       varchar
        unique,
    access_expired_at  timestamp with time zone not null,
    refresh_token      varchar
        unique,
    refresh_expired_at timestamp with time zone not null,
    issued_at          timestamp with time zone not null
);
