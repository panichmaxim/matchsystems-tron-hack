--- add final risk field

alter table billing_request
    add risk double precision;

--- fill risk field with values

update billing_request
set risk = calculated_risk
where calculated_risk > 0;
update billing_request
set risk = reported_risk
where reported_risk > 0;
update billing_request
set risk = 0
where risk is null;

--- mark risk field as not null

alter table billing_request
    alter column risk set not null;

--- add new risks and category for wallet

alter table billing_request
    add wallet_risk double precision;

alter table billing_request
    add wallet_category varchar;

--- mark requests as duplicates

alter table billing_request
    add last bool default true;

with subquery AS (select dups.id
                  from (select id,
                               row_number() over (partition by query, network, user_id order by created_at desc) as row
                        from billing_request) dups
                  where dups.row > 1)
update billing_request b
set last = false
from subquery
where b.id = subquery.id;
