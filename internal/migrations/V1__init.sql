CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name TEXT NOT NULL
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          price NUMERIC(10,2) NOT NULL,
                          image_url TEXT,
                          category_id INT REFERENCES categories(id)
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES users(id),
                        status TEXT NOT NULL,
                        total_price NUMERIC(10,2) NOT NULL,
                        created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INT REFERENCES orders(id),
                             product_id INT REFERENCES products(id),
                             quantity INT NOT NULL,
                             price NUMERIC(10,2) NOT NULL
);