BEGIN;

CREATE TABLE category (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE ebook (
    ebook_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT NOT NULL REFERENCES category(id),
    price INT NOT NULL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE "user" (
    user_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE "order" (
    order_id INT PRIMARY KEY,
    user_id INT NOT NULL REFERENCES "user"(user_id),
    payment_status VARCHAR(50) NOT NULL,
    total_payment INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE order_details (
    order_details_id INT PRIMARY KEY,
    ebook_id INT NOT NULL REFERENCES ebook(ebook_id),
    order_id INT NOT NULL REFERENCES "order"(order_id),
    price INT NOT NULL,
    quantity INT NOT NULL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);


CREATE UNIQUE INDEX ebook_name_per_category ON ebook(name, category_id); -- Nama ebook gaboleh sama di dalam satu kategori, tapi boleh sama di kategori lain. Contoh "Science" ada ebook "A", di kategori "Fiction" juga ada ebook "A"
CREATE UNIQUE INDEX orderitem_per_order ON order_details(order_id, ebook_id); -- Satu order gaboleh ada ebook yang sama, kalo mau 2 pcs / quantity.

CREATE INDEX order_user ON "order"(user_id); -- buat mempercepat pencarian user di tabel order
CREATE UNIQUE INDEX email_user ON "user"(email); -- buat mempercepat pencarian user di tabel user

COMMIT;