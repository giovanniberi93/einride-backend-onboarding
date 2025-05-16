CREATE TABLE shippers (
    shipper_id STRING(63) NOT NULL,
    create_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    update_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    delete_time TIMESTAMP OPTIONS (allow_commit_timestamp=true),
    display_name STRING(63) NOT NULL,
) PRIMARY KEY(shipper_id);
