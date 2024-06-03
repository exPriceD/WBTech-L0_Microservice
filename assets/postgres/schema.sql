CREATE TABLE IF NOT EXISTS orders
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

CREATE TABLE IF NOT EXISTS delivery
(
    delivery_id SERIAL UNIQUE PRIMARY KEY,
    order_uid   VARCHAR,
    name        VARCHAR,
    phone       VARCHAR,
    zip         VARCHAR,
    city        VARCHAR,
    address     VARCHAR,
    region      VARCHAR,
    email       VARCHAR,
    FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);

CREATE TABLE IF NOT EXISTS payment
(
    transaction   VARCHAR UNIQUE PRIMARY KEY,
    request_id    VARCHAR,
    currency      VARCHAR,
    provider      VARCHAR,
    amount        INTEGER,
    payment_dt    INTEGER,
    bank          VARCHAR,
    delivery_cost INTEGER,
    goods_total   INTEGER,
    custom_fee    INTEGER,
    FOREIGN KEY (transaction) REFERENCES orders (order_uid)
);

CREATE TABLE IF NOT EXISTS items
(
    chrt_id      INTEGER PRIMARY KEY,
    order_uid    VARCHAR,
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
