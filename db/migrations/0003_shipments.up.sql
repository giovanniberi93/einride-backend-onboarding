CREATE TABLE shipments (
    shipper_id STRING(63) NOT NULL,
    shipment_id STRING(63) NOT NULL,
    create_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    update_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    delete_time TIMESTAMP OPTIONS (allow_commit_timestamp=true),
    origin_site_id STRING(63) NOT NULL,
    destination_site_id STRING(63) NOT NULL,
    pickup_earliest_time TIMESTAMP NOT NULL,
    pickup_latest_time TIMESTAMP NOT NULL,
    delivery_earliest_time TIMESTAMP NOT NULL,
    delivery_latest_time TIMESTAMP NOT NULL,
    annotations ARRAY<STRING(MAX)> NOT NULL,
    CONSTRAINT fk_shipments_parent
        FOREIGN KEY (shipper_id)
        REFERENCES shippers (shipper_id),
    CONSTRAINT fk_shipments_origin_site
        FOREIGN KEY (shipper_id, origin_site_id)
        REFERENCES sites (shipper_id, site_id),
    CONSTRAINT fk_shipments_destination_site
        FOREIGN KEY (shipper_id, destination_site_id)
        REFERENCES sites (shipper_id, site_id),
) PRIMARY KEY(shipper_id, shipment_id);
