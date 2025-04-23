CREATE TABLE ebook (
    ebook_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT NOT NULL REFERENCES category(id),
    price INT NOT NULL
);