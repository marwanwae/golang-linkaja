-- Membuat tabel tbl_customer
CREATE TABLE tbl_customer (
    customer_number INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Membuat tabel tbl_account
CREATE TABLE tbl_account (
    account_number INT PRIMARY KEY,
    customer_number INT,
    balance INT,
    FOREIGN KEY (customer_number) REFERENCES tbl_customer(customer_number)
);