CREATE TABLE orders
(
    order_uid          VARCHAR PRIMARY KEY,
    track_number       VARCHAR,
    entry              VARCHAR,
    locale             VARCHAR,
    internal_signature VARCHAR,
    customer_id        VARCHAR,
    delivery_service   VARCHAR,
    shardkey           VARCHAR,
    sm_id              INTEGER,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR
);

CREATE TABLE delivery
(
    order_uid VARCHAR PRIMARY KEY,
    name      VARCHAR,
    phone     VARCHAR,
    zip       VARCHAR,
    city      VARCHAR,
    address   VARCHAR,
    region    VARCHAR,
    email     VARCHAR,
    FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);

CREATE TABLE payment
(
    order_uid     VARCHAR PRIMARY KEY,
    transaction   VARCHAR,
    request_id    VARCHAR,
    currency      VARCHAR,
    provider      VARCHAR,
    amount        INTEGER,
    payment_dt    INTEGER,
    bank          VARCHAR,
    delivery_cost INTEGER,
    goods_total   INTEGER,
    custom_fee    INTEGER,
    FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);

CREATE TABLE items
(
    item_id      SERIAL PRIMARY KEY,
    order_uid    VARCHAR,
    chrt_id      INTEGER,
    track_number VARCHAR,
    price        INTEGER,
    rid          VARCHAR,
    name         VARCHAR,
    sale         INTEGER,
    size         VARCHAR,
    total_price  INTEGER,
    nm_id        INTEGER,
    brand        VARCHAR,
    status       INTEGER,
    FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);