alter table billing_request
    add network varchar default 'btc';

alter table billing_request
    alter column network set not null;

alter table billing_request
    alter column network drop default;
