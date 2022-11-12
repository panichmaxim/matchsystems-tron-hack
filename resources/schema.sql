create sequence migrations_id_seq;
create sequence migrations_lock_id_seq;
create sequence users_id_seq;
create sequence billing_id_seq;
create sequence billing_packet_id_seq;
create sequence billing_request_id_seq;
create sequence billing_api_key_id_seq;
create sequence access_requests_id_seq;
create sequence category_id_seq;
create sequence billing_calculated_risk_id_seq;

create table public.access_requests
(
    id         bigint primary key       not null default nextval('access_requests_id_seq'::regclass),
    user_id    bigint                   not null,
    created_at timestamp with time zone not null default CURRENT_TIMESTAMP
);

create table public.auth_token
(
    id                 uuid primary key         not null,
    user_id            bigint                   not null,
    access_token       character varying,
    access_expired_at  timestamp with time zone not null,
    refresh_token      character varying,
    refresh_expired_at timestamp with time zone not null,
    issued_at          timestamp with time zone not null
);
create unique index auth_token_access_token_key on auth_token using btree (access_token);
create unique index auth_token_refresh_token_key on auth_token using btree (refresh_token);

create table public.billing
(
    id         bigint primary key       not null default nextval('billing_id_seq'::regclass),
    user_id    bigint                   not null,
    requests   bigint                   not null,
    created_at timestamp with time zone not null default CURRENT_TIMESTAMP
);

create table public.billing_calculated_risk
(
    id                 bigint primary key       not null default nextval('billing_calculated_risk_id_seq'::regclass),
    billing_request_id bigint                   not null,
    risk               double precision         not null,
    directory_id       bigint                   not null,
    total              double precision         not null,
    percent            double precision         not null,
    created_at         timestamp with time zone not null default CURRENT_TIMESTAMP,
    risk_raw           double precision,
    percent_raw        double precision
);

create table public.billing_key
(
    id         bigint primary key       not null default nextval('billing_api_key_id_seq'::regclass),
    key        character varying        not null,
    user_id    bigint                   not null,
    billing_id bigint                   not null,
    created_at timestamp with time zone not null default CURRENT_TIMESTAMP
);
create unique index billing_api_key_api_key_key on billing_key using btree (key);

create table public.billing_key
(
    id         bigint primary key       not null default nextval('billing_api_key_id_seq'::regclass),
    key        character varying        not null,
    user_id    bigint                   not null,
    billing_id bigint                   not null,
    created_at timestamp with time zone not null default CURRENT_TIMESTAMP
);
create unique index billing_api_key_api_key_key on billing_key using btree (key);

create table public.billing_packet
(
    id         bigint primary key       not null default nextval('billing_packet_id_seq'::regclass),
    user_id    bigint                   not null,
    requests   bigint                   not null,
    created_at timestamp with time zone not null default CURRENT_TIMESTAMP
);

create table public.billing_request
(
    id                bigint primary key       not null default nextval('billing_request_id_seq'::regclass),
    user_id           bigint                   not null,
    query             character varying        not null,
    created_at        timestamp with time zone not null default CURRENT_TIMESTAMP,
    reported_category character varying(255),
    reported_risk     double precision,
    network           character varying        not null,
    calculated_risk   double precision,
    calculated_total  double precision
);

create table public.category
(
    id                bigint primary key not null default nextval('category_id_seq'::regclass),
    name              character varying  not null,
    risk              bigint             not null,
    description_ru    character varying  not null,
    description_en    character varying  not null,
    category_group_id bigint,
    number            integer            not null
);

create table public.migrations
(
    id          bigint primary key       not null default nextval('migrations_id_seq'::regclass),
    name        character varying,
    group_id    bigint,
    migrated_at timestamp with time zone not null default CURRENT_TIMESTAMP
);

create table public.migrations_lock
(
    id         bigint primary key not null default nextval('migrations_lock_id_seq'::regclass),
    table_name character varying
);
create unique index migrations_lock_table_name_key on migrations_lock using btree (table_name);

create table public.users
(
    id          bigint primary key       not null default nextval('users_id_seq'::regclass),
    name        character varying        not null,
    email       character varying        not null,
    is_active   boolean                  not null default false,
    password    character varying,
    token       character varying,
    permissions character varying[],
    created_at  timestamp with time zone not null default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone not null default CURRENT_TIMESTAMP
);
create unique index users_email_key on users using btree (email);
create unique index users_token_key on users using btree (token);