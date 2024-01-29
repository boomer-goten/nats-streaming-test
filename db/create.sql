DROP TABLE IF EXISTS Order_info;
DROP TABLE IF EXISTS Delivery;
DROP TABLE IF EXISTS Payment;
DROP TABLE IF EXISTS Items;
CREATE TABLE IF NOT EXISTS Order_info(
    order_uid varchar PRIMARY KEY not null,
    track_number varchar not null,
    entry varchar not null,
    locale varchar,
    internal_signature varchar,
    customer_id varchar not null,
    delivery_service varchar not null,
    shardKey int not null,
    sm_id int not null,
    date_created timestamp not null,
    off_shard int not null
);
CREATE TABLE IF NOT EXISTS Delivery(
    order_uid varchar PRIMARY KEY not null,
    name varchar not null,
    phone varchar,
    zip int not null,
    city varchar not null,
    address varchar not null,
    region varchar not null,
    email varchar,
    CONSTRAINT fk_delivery_order_uid FOREIGN KEY (order_uid) REFERENCES ORder_info(order_uid)
);
CREATE TABLE IF NOT EXISTS Payment(
    transaction varchar PRIMARY KEY not null,
    request_id varchar,
    currency varchar,
    provider varchar,
    amount real not null,
    payment_dt int not null,
    bank varchar not null,
    delivery_cost real,
    goods_total int,
    custom_fee int,
    order_uid varchar,
    CONSTRAINT fk_payment_order_uid FOREIGN KEY (order_uid) REFERENCES ORder_info(order_uid)
);
CREATE TABLE IF NOT EXISTS items (
    chrt_id int PRIMARY KEY not null,
    transaction varchar not null,
    price real not null,
    rid varchar,
    name varchar not null,
    sale real,
    total_price real not null,
    nm_id int not null,
    brand varchar,
    status int not null,
    CONSTRAINT fk_items_transaction FOREIGN KEY (TRANSACTION) REFERENCES Payment(TRANSACTION)
);