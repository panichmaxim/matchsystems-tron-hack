alter table billing_request
    add category varchar(255);

alter table billing_request
    add risk integer;

alter table billing_key
    rename column api_key to key;

drop index billing_api_key_api_key_key;

alter table billing_key
    drop constraint billing_api_key_api_key_key;

alter table billing_key
    add constraint billing_api_key_api_key_key
        unique (api_key);
