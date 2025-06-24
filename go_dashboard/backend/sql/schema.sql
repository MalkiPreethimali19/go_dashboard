-- 1. Create the database
DROP DATABASE IF EXISTS dashboard_db;
CREATE DATABASE dashboard_db;
USE dashboard_db;

-- 2. Create the Regions table
CREATE TABLE regions (
    region_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL
);

-- 3. Create the Customers table
CREATE TABLE customers (
    customer_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    region_id INT,
    FOREIGN KEY (region_id) REFERENCES regions(region_id)
        ON DELETE SET NULL ON UPDATE CASCADE
);

-- 4. Create the Products table
CREATE TABLE products (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100),
    stock_quantity INT DEFAULT 0,
    unit_price DECIMAL(10, 2) NOT NULL
);

-- 5. Create the Transactions table
CREATE TABLE transactions (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    customer_id INT,
    transaction_date DATE NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(12, 2) GENERATED ALWAYS AS 
        (quantity * (SELECT unit_price FROM products WHERE products.product_id = transactions.product_id)) STORED,
    
    FOREIGN KEY (product_id) REFERENCES products(product_id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
        ON DELETE SET NULL ON UPDATE CASCADE
);

-- 6. Indexes for performance
CREATE INDEX idx_transaction_date ON transactions(transaction_date);
CREATE INDEX idx_country ON regions(country);
