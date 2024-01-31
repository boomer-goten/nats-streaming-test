DROP TABLE IF EXISTS Order_info CASCADE;
DROP TABLE IF EXISTS Delivery CASCADE;
DROP TABLE IF EXISTS Payment CASCADE;
DROP TABLE IF EXISTS Items CASCADE;
CREATE TABLE IF NOT EXISTS Order_info(
    order_uid varchar PRIMARY KEY not null,
    track_number varchar not null,
    entry varchar not null,
    locale varchar,
    internal_signature varchar,
    customer_id varchar not null,
    delivery_service varchar not null,
    shardKey varchar not null,
    sm_id int not null,
    date_created timestamp not null,
    off_shard varchar not null
);
CREATE TABLE IF NOT EXISTS Delivery(
    order_uid varchar PRIMARY KEY not null,
    name varchar not null,
    phone varchar,
    zip varchar not null,
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
    amount int not null,
    payment_dt int not null,
    bank varchar not null,
    delivery_cost int,
    goods_total int,
    custom_fee int,
    order_uid varchar,
    CONSTRAINT fk_payment_order_uid FOREIGN KEY (order_uid) REFERENCES ORder_info(order_uid)
);
CREATE TABLE IF NOT EXISTS items (
    chrt_id int PRIMARY KEY not null,
    transaction varchar not null,
    price int not null,
    rid varchar,
    name varchar not null,
    sale int,
    size varchar,
    total_price int not null,
    nm_id int not null,
    brand varchar,
    status int not null,
    CONSTRAINT fk_items_transaction FOREIGN KEY (TRANSACTION) REFERENCES Payment(TRANSACTION)
);
-- SELECT *
-- FROM Order_info
--     JOIN Delivery on Order_info.order_uid = Delivery.order_uid
--     JOIN Payment on Order_info.order_uid = Payment.order_uid
-- SELECT *
-- FROM items
-- WHERE items.transaction = '
SELECT chrt_id,
    price,
    rid,
    name,
    sale,
    size,
    total_price,
    nm_id,
    brand,
    status
FROM items