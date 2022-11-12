alter table billing_request
    add is_wallet bool default false not null;
alter table billing_request
    add is_reported bool default false not null;
alter table billing_request
    add is_calculated bool default false not null;

-- column reordering is not supported billing_request.is_reported

alter table billing_request
    drop column reported_category;

alter table billing_request
    drop column reported_risk;

alter table billing_request
    drop column calculated_risk;

alter table billing_request
    drop column calculated_total;

alter table billing_request
    drop column wallet_risk;

alter table billing_request
    drop column wallet_category;

create table billing_risk
(
    id                 bigserial
        primary key,
    billing_request_id bigint                                             not null,
    is_reported        boolean                                            not null,
    is_wallet          boolean                                            not null,
    is_calculated      boolean                                            not null,
    risk               double precision                                   not null,
    risk_raw           double precision                                   not null,
    directory_id       bigint                                             not null,
    category_id        bigint,
    total              double precision                                   not null,
    created_at         timestamp with time zone default CURRENT_TIMESTAMP not null
);

drop table billing_calculated_risk;