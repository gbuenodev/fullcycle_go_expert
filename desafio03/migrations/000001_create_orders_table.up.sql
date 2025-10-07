-- filepath: /Users/admin/FullCycle/fullcytcle_go_expert/desafio03/migrations/000001_create_orders_table.down.sql
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    tax DECIMAL(10,2) NOT NULL,
    final_price DECIMAL(10,2) NOT NULL,
    PRIMARY KEY (id)
);