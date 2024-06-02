CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE listings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    price INT NOT NULL CHECK (price > 0),
    listing_type VARCHAR(10) NOT NULL CHECK (listing_type IN ('rent', 'sale')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
