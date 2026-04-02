
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    category TEXT NOT NULL,
    in_stock BOOLEAN NOT NULL DEFAULT true
);

INSERT INTO products (name, price, category, in_stock) VALUES
('Perceuse visseuse', 89.99, 'outillage', true),
('Aspirateur robot', 249.99, 'électroménager', false),
('Casque audio sans fil', 129.99, 'high-tech', true);