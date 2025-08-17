-- Пользователи системы
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       phone VARCHAR(20),
                       password_hash VARCHAR(255) NOT NULL,
                       role VARCHAR(50) DEFAULT 'customer' CHECK(role IN ('customer', 'restaurant_owner', 'courier', 'admin')),
                       email_verified BOOLEAN DEFAULT FALSE,
                       phone_verified BOOLEAN DEFAULT FALSE,
                       is_active BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

-- Адреса доставки пользователей
CREATE TABLE addresses (
                           id SERIAL PRIMARY KEY,
                           user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                           title VARCHAR(100),
                           street VARCHAR(255) NOT NULL,
                           house VARCHAR(50) NOT NULL,
                           apartment VARCHAR(50),
                           entrance VARCHAR(50),
                           floor VARCHAR(50),
                           comment TEXT,
                           latitude DECIMAL(10,8),
                           longitude DECIMAL(11,8),
                           is_default BOOLEAN DEFAULT FALSE,
                           created_at TIMESTAMP DEFAULT NOW()
);

-- Рестораны
CREATE TABLE restaurants (
                             id SERIAL PRIMARY KEY,
                             owner_id INTEGER REFERENCES users(id),
                             name VARCHAR(255) NOT NULL,
                             description TEXT,
                             phone VARCHAR(20),
                             email VARCHAR(255),
                             address VARCHAR(500) NOT NULL,
                             latitude DECIMAL(10,8),
                             longitude DECIMAL(11,8),
                             image_url VARCHAR(500),
                             min_order_amount DECIMAL(10,2) DEFAULT 0,
                             delivery_fee DECIMAL(10,2) DEFAULT 0,
                             delivery_time_min INTEGER DEFAULT 30,
                             delivery_time_max INTEGER DEFAULT 60,
                             rating DECIMAL(3,2) DEFAULT 0,
                             reviews_count INTEGER DEFAULT 0,
                             is_active BOOLEAN DEFAULT TRUE,
                             is_featured BOOLEAN DEFAULT FALSE,
                             created_at TIMESTAMP DEFAULT NOW(),
                             updated_at TIMESTAMP DEFAULT NOW()
);

-- Часы работы ресторанов
CREATE TABLE restaurant_hours (
                                  id SERIAL PRIMARY KEY,
                                  restaurant_id INTEGER REFERENCES restaurants(id) ON DELETE CASCADE,
                                  day_of_week INTEGER NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
                                  open_time TIME,
                                  close_time TIME,
                                  is_closed BOOLEAN DEFAULT FALSE
);

-- Зоны доставки ресторанов
CREATE TABLE delivery_zones (
                                id SERIAL PRIMARY KEY,
                                restaurant_id INTEGER REFERENCES restaurants(id) ON DELETE CASCADE,
                                zone_name VARCHAR(255),
                                coordinates JSON,
                                delivery_fee DECIMAL(10, 2) DEFAULT 0,
                                min_order_amount DECIMAL(10, 2) DEFAULT 0
);

-- Категории продуктов
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            restaurant_id INTEGER REFERENCES restaurants(id) ON DELETE CASCADE,
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            image_url VARCHAR(500),
                            sort_order INTEGER DEFAULT 0,
                            is_active BOOLEAN DEFAULT TRUE,
                            created_at TIMESTAMP DEFAULT NOW()
);

-- Продукты
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          restaurant_id INTEGER REFERENCES restaurants(id) ON DELETE CASCADE,
                          category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price DECIMAL(10,2) NOT NULL,
                          old_price DECIMAL(10, 2),
                          image_url VARCHAR(500),
                          weight VARCHAR(50),
                          calories INTEGER,
                          proteins DECIMAL(5, 2),
                          fats DECIMAL(5, 2),
                          carbs DECIMAL(5, 2),
                          allergens TEXT[],
                          is_vegetarian BOOLEAN DEFAULT FALSE,
                          is_vegan BOOLEAN DEFAULT FALSE,
                          is_popular BOOLEAN DEFAULT FALSE,
                          is_available BOOLEAN DEFAULT TRUE,
                          sort_order INTEGER DEFAULT 0,
                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW()
);

-- Опции продуктов (размеры, добавки)
CREATE TABLE product_options (
                                 id SERIAL PRIMARY KEY,
                                 product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                                 name VARCHAR(255) NOT NULL,
                                 type VARCHAR(50) DEFAULT 'single' CHECK (type IN ('single','multiple')),
                                 is_required BOOLEAN DEFAULT FALSE,
                                 sort_order INTEGER DEFAULT 0
);

-- Варианты опций
CREATE TABLE product_option_values (
                                       id SERIAL PRIMARY KEY,
                                       option_id INTEGER REFERENCES product_options(id) ON DELETE CASCADE,
                                       name VARCHAR(255) NOT NULL,
                                       price_modifier DECIMAL(10, 2) DEFAULT 0,
                                       is_default BOOLEAN DEFAULT FALSE,
                                       sort_order INTEGER DEFAULT 0
);

-- Промокоды
CREATE TABLE promo_codes (
                             id SERIAL PRIMARY KEY,
                             code VARCHAR(50) UNIQUE NOT NULL,
                             name VARCHAR(255),
                             description TEXT,
                             type VARCHAR(50) DEFAULT 'percentage' CHECK (type IN ('percentage', 'fixed', 'free_delivery')),
                             value DECIMAL(10, 2) NOT NULL,
                             min_order_amount DECIMAL(10, 2) DEFAULT 0,
                             max_discount DECIMAL(10,2),
                             usage_limit INTEGER,
                             usage_limit_per_user INTEGER DEFAULT 1,
                             usage_count INTEGER DEFAULT 0,
                             valid_from TIMESTAMP DEFAULT NOW(),
                             valid_until TIMESTAMP,
                             is_active BOOLEAN DEFAULT TRUE,
                             created_at TIMESTAMP DEFAULT NOW()
);

-- Товары в корзине
CREATE TABLE cart_items (
                            id SERIAL PRIMARY KEY,
                            user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                            product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                            quantity INTEGER NOT NULL DEFAULT 1,
                            selected_options JSON,
                            created_at TIMESTAMP DEFAULT NOW(),
                            updated_at TIMESTAMP DEFAULT NOW()
);

-- Заказы
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER REFERENCES users(id),
                        restaurant_id INTEGER REFERENCES restaurants(id),
                        courier_id INTEGER REFERENCES users(id),

                        status VARCHAR(50) DEFAULT 'pending' CHECK (status IN(
                                                                              'pending', 'confirmed', 'preparing', 'ready',
                                                                              'picked_up', 'delivering', 'delivered',
                                                                              'cancelled', 'refunded'
                            )),

    -- Финансы
                        subtotal DECIMAL(10, 2) NOT NULL,
                        delivery_fee DECIMAL(10, 2) DEFAULT 0,
                        service_fee DECIMAL(10, 2) DEFAULT 0,
                        discount DECIMAL(10, 2) DEFAULT 0,
                        total_amount DECIMAL(10, 2) NOT NULL,

    -- Доставка
                        delivery_type VARCHAR(50) DEFAULT 'delivery' CHECK (delivery_type IN ('delivery', 'pickup')),
                        delivery_address JSON,
                        delivery_time_requested TIMESTAMP,
                        delivery_time_estimated TIMESTAMP,
                        delivery_time_actual TIMESTAMP,

    -- Оплата
                        payment_method VARCHAR(50) DEFAULT 'cash' CHECK (payment_method IN ('cash', 'card', 'online')),
                        payment_status VARCHAR(50) DEFAULT 'pending' CHECK (payment_status IN ('pending', 'paid', 'failed','refunded')),
                        payment_id VARCHAR(255),

    -- Промокод
                        promo_code_id INTEGER REFERENCES promo_codes(id),

    -- Дополнительно
                        comment TEXT,
                        rating INTEGER CHECK (rating BETWEEN 1 AND 5),
                        review TEXT,

                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
);

-- Позиции заказа
CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
                             product_id INTEGER REFERENCES products(id),
                             name VARCHAR(255) NOT NULL,
                             price DECIMAL(10, 2) NOT NULL,
                             quantity INTEGER NOT NULL,
                             selected_options JSON,
                             subtotal DECIMAL(10, 2) NOT NULL
);

-- История статусов заказа
CREATE TABLE order_status_history (
                                      id SERIAL PRIMARY KEY,
                                      order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
                                      status VARCHAR(50) NOT NULL,
                                      comment TEXT,
                                      created_by INTEGER REFERENCES users(id),
                                      created_at TIMESTAMP DEFAULT NOW()
);

-- Использования промокодов
CREATE TABLE promo_code_usages (
                                   id SERIAL PRIMARY KEY,
                                   promo_code_id INTEGER REFERENCES promo_codes(id),
                                   user_id INTEGER REFERENCES users(id),
                                   order_id INTEGER REFERENCES orders(id),
                                   discount_amount DECIMAL(10, 2),
                                   used_at TIMESTAMP DEFAULT NOW()
);

-- Отзывы
CREATE TABLE reviews (
                         id SERIAL PRIMARY KEY,
                         user_id INTEGER REFERENCES users(id),
                         restaurant_id INTEGER REFERENCES restaurants(id),
                         order_id INTEGER REFERENCES orders(id),
                         rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
                         comment TEXT,
                         food_rating INTEGER CHECK (food_rating BETWEEN 1 AND 5),
                         delivery_rating INTEGER CHECK (delivery_rating BETWEEN 1 AND 5),
                         is_visible BOOLEAN DEFAULT TRUE,
                         created_at TIMESTAMP DEFAULT NOW(),

                         UNIQUE (user_id, order_id)
);

-- Уведомления
CREATE TABLE notifications(
                              id SERIAL PRIMARY KEY,
                              user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                              type VARCHAR(50) NOT NULL,
                              title VARCHAR(255) NOT NULL,
                              message TEXT NOT NULL,
                              data JSON,
                              is_read BOOLEAN DEFAULT FALSE,
                              created_at TIMESTAMP DEFAULT NOW()
);

-- Пользователи
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);

-- Рестораны
CREATE INDEX idx_restaurants_location ON restaurants(latitude,longitude);
CREATE INDEX idx_restaurants_active ON restaurants(is_active);

-- Продукты
CREATE INDEX idx_products_restaurant ON products(restaurant_id);
CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_products_available ON products(is_available);

-- Заказы
CREATE INDEX idx_orders_user ON orders(user_id);
CREATE INDEX idx_orders_restaurant ON orders(restaurant_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_created ON orders(created_at);

-- Корзина
CREATE INDEX idx_cart_user ON cart_items(user_id);