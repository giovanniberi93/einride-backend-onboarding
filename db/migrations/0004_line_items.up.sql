CREATE TABLE line_items (
    shipper_id STRING(63) NOT NULL,
    shipment_id STRING(63) NOT NULL,
    line_number INT64 NOT NULL,
    title STRING(63) NOT NULL,
    quantity FLOAT64 NOT NULL,
    weight_kg FLOAT64 NOT NULL,
    volume_m3 FLOAT64 NOT NULL,
) PRIMARY KEY(shipper_id, shipment_id, line_number ASC),
  INTERLEAVE IN PARENT shipments ON DELETE CASCADE;
