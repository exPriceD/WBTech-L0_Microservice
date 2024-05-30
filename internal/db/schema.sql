CREATE TABLE orders
(
    order_uid          VARCHAR PRIMARY KEY,
    track_number       VARCHAR,
    delivery_id        INTEGER,
    entry              VARCHAR,
    locale             VARCHAR,
    internal_signature VARCHAR,
    customer_id        VARCHAR,
    delivery_service   VARCHAR,
    shardkey           VARCHAR,
    sm_id              INTEGER,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR,
    FOREIGN KEY (delivery_id) REFERENCES delivery (delivery_id),
    FOREIGN KEY (order_uid) REFERENCES payment (transaction),
    FOREIGN KEY (track_number) REFERENCES items (track_number)
);

CREATE TABLE delivery
(
    delivery_id INTEGER PRIMARY KEY,
    name      VARCHAR,
    phone     VARCHAR,
    zip       VARCHAR,
    city      VARCHAR,
    address   VARCHAR,
    region    VARCHAR,
    email     VARCHAR
);

CREATE TABLE payment
(
    transaction   VARCHAR PRIMARY KEY,
    request_id    VARCHAR,
    currency      VARCHAR,
    provider      VARCHAR,
    amount        INTEGER,
    payment_dt    INTEGER,
    bank          VARCHAR,
    delivery_cost INTEGER,
    goods_total   INTEGER,
    custom_fee    INTEGER
);

CREATE TABLE items
(
    track_number VARCHAR PRIMARY KEY,
    chrt_id      INTEGER,
    price        INTEGER,
    rid          VARCHAR,
    name         VARCHAR,
    sale         INTEGER,
    size         VARCHAR,
    total_price  INTEGER,
    nm_id        INTEGER,
    brand        VARCHAR,
    status       INTEGER
);