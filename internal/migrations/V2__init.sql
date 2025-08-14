-- Добавить недостающие поля в существующие таблицы
ALTER TABLE users ADD COLUMN phone TEXT;
ALTER TABLE users ADD COLUMN role TEXT DEFAULT 'client';

ALTER TABLE orders ADD COLUMN delivery_type TEXT DEFAULT 'delivery';
ALTER TABLE orders ADD COLUMN payment_method TEXT DEFAULT 'cash';
ALTER TABLE orders ADD COLUMN delivery_address TEXT;
ALTER TABLE orders ADD COLUMN comment TEXT;

-- Новые таблицы для корзины (когда понадобится)
CREATE TABLE cart_items (
                            id SERIAL PRIMARY KEY,
                            user_id INT REFERENCES users(id),
                            product_id INT REFERENCES products(id),
                            quantity INT NOT NULL DEFAULT 1,
                            created_at TIMESTAMP DEFAULT now()
);

-- Адреса доставки (когда понадобится)
CREATE TABLE addresses (
                           id SERIAL PRIMARY KEY,
                           user_id INT REFERENCES users(id),
                           title TEXT,
                           street TEXT NOT NULL,
                           house TEXT NOT NULL,
                           apartment TEXT,
                           comment TEXT,
                           is_default BOOLEAN DEFAULT false
);