CREATE TABLE products (
  id uuid PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  category VARCHAR(20) NOT NULL CHECK (category IN ('SANDWICH', 'SIDEDISHES', 'DRINKS', 'DESSERTS')),
  price NUMERIC(10, 2) NOT NULL,
  description TEXT,
  image TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_category ON products(category);
