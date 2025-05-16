CREATE TABLE sites (
    shipper_id STRING(63) NOT NULL,
    site_id STRING(63) NOT NULL,
    create_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    update_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    delete_time TIMESTAMP OPTIONS (allow_commit_timestamp=true),
    display_name STRING(63) NOT NULL,
    latitude FLOAT64,
    longitude FLOAT64,
    CONSTRAINT fk_sites_parent
        FOREIGN KEY (shipper_id)
        REFERENCES shippers (shipper_id),
) PRIMARY KEY(shipper_id, site_id);
