create table billing_calculated_risk
(
    id                 bigserial
        primary key,
    billing_request_id bigint                                             not null,
    risk               double precision                                   not null,
    directory_id       bigint                                             not null,
    total              double precision                                   not null,
    percent            double precision                                   not null,
    created_at         timestamp with time zone default CURRENT_TIMESTAMP not null
);

alter table billing_request
    add calculated_risk double precision;

alter table billing_request
    rename column category to reported_category;

alter table billing_request
    rename column risk to reported_risk;

alter table billing_request
    alter column reported_risk type double precision using reported_risk::double precision;

