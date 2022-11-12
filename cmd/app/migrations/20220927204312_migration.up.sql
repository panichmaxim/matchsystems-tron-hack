update billing_request
set reported_category = null
where reported_category = '';

UPDATE billing_request AS v
SET reported_category = s.id
FROM category AS s
WHERE lower(v.reported_category) = lower(s.name);

alter table billing_request
    alter column reported_category type int using reported_category::int;