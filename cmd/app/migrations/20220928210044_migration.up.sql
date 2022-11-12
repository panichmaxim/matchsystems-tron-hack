alter table billing_request
    alter column wallet_category type int using wallet_category::int;
