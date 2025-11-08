# PostgreSQL Coding Standards

## General Principles

- Leverage PostgreSQL-specific features for performance and functionality
- Use appropriate data types including PostgreSQL extensions
- Follow SQL standards while utilizing PostgreSQL enhancements
- Optimize for MVCC (Multi-Version Concurrency Control)
- Use schemas for logical organization

## Naming Conventions

- Tables: `plural_snake_case` (e.g., `users`, `order_items`)
- Columns: `snake_case` (e.g., `first_name`, `created_at`)
- Schemas: `snake_case` (e.g., `public`, `analytics`, `audit`)
- Primary keys: `id` or `{table_name}_id`
- Foreign keys: `{referenced_table}_id`
- Indexes: `idx_{table}_{columns}` or `{table}_{columns}_idx`
- Constraints: `{table}_{columns}_{type}` (e.g., `users_email_key`, `orders_user_id_fkey`)

## PostgreSQL-Specific Data Types

- Use `SERIAL` or `BIGSERIAL` for auto-incrementing IDs
- Use `UUID` for distributed systems or when exposing IDs externally
- Use `JSONB` for semi-structured data (not JSON)
- Use `ARRAY` types for lists of values
- Use `ENUM` types for fixed sets of values
- Use `TIMESTAMP WITH TIME ZONE` (timestamptz) for timestamps
- Use `TEXT` instead of VARCHAR without length limit

```sql
CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped', 'delivered', 'cancelled');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    email TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    preferences JSONB DEFAULT '{}'::jsonb,
    tags TEXT[] DEFAULT ARRAY[]::TEXT[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    status order_status NOT NULL DEFAULT 'pending',
    order_date DATE NOT NULL DEFAULT CURRENT_DATE,
    total_amount NUMERIC(10, 2) NOT NULL,
    metadata JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT orders_total_amount_check CHECK (total_amount >= 0)
);
```

## Schemas

- Use schemas to organize related tables
- Default schema is `public`
- Create separate schemas for different domains or purposes

```sql
-- Create schemas
CREATE SCHEMA analytics;
CREATE SCHEMA audit;

-- Create table in specific schema
CREATE TABLE analytics.user_metrics (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    metric_date DATE NOT NULL,
    page_views INTEGER NOT NULL DEFAULT 0
);

-- Query with schema prefix
SELECT * FROM analytics.user_metrics;

-- Set search path
SET search_path TO analytics, public;
```

## Indexing

- Use B-tree indexes (default) for equality and range queries
- Use GIN indexes for JSONB, arrays, and full-text search
- Use GiST indexes for geometric data and full-text search
- Use partial indexes for filtered queries
- Use expression indexes for computed values

```sql
-- Standard B-tree index
CREATE INDEX idx_users_email ON users(email);

-- Composite index
CREATE INDEX idx_orders_user_date ON orders(user_id, order_date DESC);

-- GIN index for JSONB
CREATE INDEX idx_users_preferences ON users USING GIN (preferences);

-- GIN index for arrays
CREATE INDEX idx_users_tags ON users USING GIN (tags);

-- Partial index (filtered)
CREATE INDEX idx_orders_pending ON orders(user_id) 
WHERE status = 'pending';

-- Expression index
CREATE INDEX idx_users_lower_email ON users(LOWER(email));

-- Full-text search index
CREATE INDEX idx_products_search ON products USING GIN (to_tsvector('english', name || ' ' || description));
```

## JSONB Operations

- Use JSONB for flexible, semi-structured data
- Index JSONB columns with GIN for performance
- Use operators: `->`, `->>`, `@>`, `?`, `?|`, `?&`

```sql
-- Insert JSONB data
INSERT INTO users (email, first_name, last_name, preferences)
VALUES ('john@example.com', 'John', 'Doe', '{"theme": "dark", "notifications": true}'::jsonb);

-- Query JSONB (-> returns JSONB, ->> returns text)
SELECT 
    email,
    preferences->>'theme' as theme,
    preferences->'notifications' as notifications
FROM users
WHERE preferences->>'theme' = 'dark';

-- Check if JSONB contains key
SELECT * FROM users WHERE preferences ? 'theme';

-- Check if JSONB contains object
SELECT * FROM users WHERE preferences @> '{"theme": "dark"}'::jsonb;

-- Update JSONB field
UPDATE users 
SET preferences = preferences || '{"language": "en"}'::jsonb
WHERE id = 1;

-- Remove JSONB key
UPDATE users 
SET preferences = preferences - 'theme'
WHERE id = 1;
```

## Array Operations

- Use arrays for ordered lists of values
- Index arrays with GIN for containment queries

```sql
-- Insert array data
INSERT INTO users (email, first_name, last_name, tags)
VALUES ('jane@example.com', 'Jane', 'Smith', ARRAY['premium', 'verified']);

-- Query arrays
SELECT * FROM users WHERE 'premium' = ANY(tags);

-- Array contains
SELECT * FROM users WHERE tags @> ARRAY['premium'];

-- Array overlap
SELECT * FROM users WHERE tags && ARRAY['premium', 'verified'];

-- Append to array
UPDATE users 
SET tags = array_append(tags, 'new_tag')
WHERE id = 1;

-- Remove from array
UPDATE users 
SET tags = array_remove(tags, 'old_tag')
WHERE id = 1;
```

## Full-Text Search

- Use `tsvector` and `tsquery` for full-text search
- Create GIN indexes on tsvector columns
- Use text search configurations for language-specific search

```sql
-- Add tsvector column
ALTER TABLE products ADD COLUMN search_vector tsvector;

-- Update tsvector column
UPDATE products 
SET search_vector = to_tsvector('english', name || ' ' || description);

-- Create index
CREATE INDEX idx_products_search ON products USING GIN (search_vector);

-- Full-text search query
SELECT 
    name,
    description,
    ts_rank(search_vector, query) as rank
FROM products, to_tsquery('english', 'laptop & wireless') as query
WHERE search_vector @@ query
ORDER BY rank DESC;

-- Trigger to auto-update tsvector
CREATE TRIGGER products_search_update 
BEFORE INSERT OR UPDATE ON products
FOR EACH ROW EXECUTE FUNCTION
tsvector_update_trigger(search_vector, 'pg_catalog.english', name, description);
```

## Window Functions

- Use window functions for analytics and ranking
- Avoid self-joins when window functions can be used

```sql
-- Ranking
SELECT 
    user_id,
    order_date,
    total_amount,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY order_date DESC) as order_rank,
    RANK() OVER (ORDER BY total_amount DESC) as amount_rank
FROM orders;

-- Running totals
SELECT 
    order_date,
    total_amount,
    SUM(total_amount) OVER (ORDER BY order_date) as running_total
FROM orders;

-- Moving average
SELECT 
    order_date,
    total_amount,
    AVG(total_amount) OVER (
        ORDER BY order_date 
        ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
    ) as moving_avg_7day
FROM orders;
```

## Common Table Expressions (CTEs)

- Use CTEs for complex queries and readability
- Use recursive CTEs for hierarchical data

```sql
-- Standard CTE
WITH active_users AS (
    SELECT id, email, first_name, last_name
    FROM users
    WHERE created_at >= NOW() - INTERVAL '30 days'
),
user_orders AS (
    SELECT 
        user_id,
        COUNT(*) as order_count,
        SUM(total_amount) as total_spent
    FROM orders
    GROUP BY user_id
)
SELECT 
    au.email,
    au.first_name,
    COALESCE(uo.order_count, 0) as orders,
    COALESCE(uo.total_spent, 0) as spent
FROM active_users au
LEFT JOIN user_orders uo ON au.id = uo.user_id;

-- Recursive CTE (hierarchical data)
WITH RECURSIVE category_tree AS (
    -- Base case
    SELECT id, name, parent_id, 1 as level
    FROM categories
    WHERE parent_id IS NULL
    
    UNION ALL
    
    -- Recursive case
    SELECT c.id, c.name, c.parent_id, ct.level + 1
    FROM categories c
    INNER JOIN category_tree ct ON c.parent_id = ct.id
)
SELECT * FROM category_tree ORDER BY level, name;
```

## Transactions and Isolation

- Use appropriate isolation levels
- Default is READ COMMITTED
- Use SERIALIZABLE for strict consistency requirements

```sql
-- Standard transaction
BEGIN;
    INSERT INTO orders (user_id, total_amount) VALUES (1, 99.99);
    UPDATE users SET last_order_date = NOW() WHERE id = 1;
COMMIT;

-- Transaction with isolation level
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
    -- Your queries here
COMMIT;

-- Savepoints for partial rollback
BEGIN;
    INSERT INTO orders (user_id, total_amount) VALUES (1, 99.99);
    SAVEPOINT order_created;
    
    INSERT INTO order_items (order_id, product_id) VALUES (1, 100);
    -- Error occurs
    ROLLBACK TO SAVEPOINT order_created;
    
COMMIT;
```

## Performance Optimization

- Use `EXPLAIN ANALYZE` to understand query performance
- Use `VACUUM` and `ANALYZE` regularly (usually automatic)
- Monitor and tune `work_mem`, `shared_buffers`, `effective_cache_size`
- Use connection pooling (PgBouncer, pgpool)
- Partition large tables

```sql
-- Explain query
EXPLAIN ANALYZE
SELECT * FROM orders WHERE user_id = 123;

-- Vacuum and analyze
VACUUM ANALYZE orders;

-- Table partitioning (range partitioning)
CREATE TABLE orders_partitioned (
    id BIGSERIAL,
    user_id BIGINT NOT NULL,
    order_date DATE NOT NULL,
    total_amount NUMERIC(10, 2)
) PARTITION BY RANGE (order_date);

CREATE TABLE orders_2024_q1 PARTITION OF orders_partitioned
    FOR VALUES FROM ('2024-01-01') TO ('2024-04-01');

CREATE TABLE orders_2024_q2 PARTITION OF orders_partitioned
    FOR VALUES FROM ('2024-04-01') TO ('2024-07-01');
```

## Constraints and Validation

- Use CHECK constraints for data validation
- Use EXCLUDE constraints for complex uniqueness rules
- Use foreign keys with appropriate ON DELETE/ON UPDATE actions

```sql
-- CHECK constraint
ALTER TABLE orders 
ADD CONSTRAINT orders_total_positive CHECK (total_amount > 0);

-- EXCLUDE constraint (requires btree_gist extension)
CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TABLE bookings (
    room_id INTEGER,
    during TSTZRANGE,
    EXCLUDE USING GIST (room_id WITH =, during WITH &&)
);

-- Foreign key with cascade
ALTER TABLE orders
ADD CONSTRAINT fk_orders_user
FOREIGN KEY (user_id) REFERENCES users(id)
ON DELETE CASCADE ON UPDATE CASCADE;
```

## Functions and Procedures

- Use functions for reusable logic
- Use procedures for transactions (PostgreSQL 11+)
- Use PL/pgSQL for complex logic

```sql
-- Function
CREATE OR REPLACE FUNCTION get_user_order_count(p_user_id BIGINT)
RETURNS INTEGER AS $$
DECLARE
    v_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO v_count
    FROM orders
    WHERE user_id = p_user_id;
    
    RETURN v_count;
END;
$$ LANGUAGE plpgsql;

-- Usage
SELECT get_user_order_count(123);

-- Procedure (can commit transactions)
CREATE OR REPLACE PROCEDURE process_order(p_user_id BIGINT, p_amount NUMERIC)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO orders (user_id, total_amount)
    VALUES (p_user_id, p_amount);
    
    UPDATE users 
    SET last_order_date = NOW()
    WHERE id = p_user_id;
    
    COMMIT;
END;
$$;

-- Usage
CALL process_order(123, 99.99);
```

## Triggers

- Use triggers for automatic actions
- Keep trigger logic simple
- Consider using BEFORE vs AFTER triggers appropriately

```sql
-- Trigger function
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger
CREATE TRIGGER users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
```

## Security

- Use row-level security (RLS) for multi-tenant applications
- Grant minimum necessary privileges
- Use roles for permission management
- Never store passwords in plain text (use pgcrypto)

```sql
-- Enable row-level security
ALTER TABLE orders ENABLE ROW LEVEL SECURITY;

-- Create policy
CREATE POLICY user_orders ON orders
FOR ALL
TO app_user
USING (user_id = current_setting('app.current_user_id')::BIGINT);

-- Create role
CREATE ROLE app_user;
GRANT SELECT, INSERT, UPDATE ON orders TO app_user;

-- Password hashing
CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users (email, password_hash)
VALUES ('user@example.com', crypt('password123', gen_salt('bf')));

-- Verify password
SELECT * FROM users 
WHERE email = 'user@example.com' 
AND password_hash = crypt('password123', password_hash);
```

## Best Practices

- Use `TEXT` instead of `VARCHAR` without specific length requirements
- Use `TIMESTAMPTZ` for all timestamps
- Use `NUMERIC` for money (not FLOAT or REAL)
- Use `gen_random_uuid()` for UUIDs (requires pgcrypto extension)
- Use `RETURNING` clause to get inserted/updated data
- Use `ON CONFLICT` for upserts
- Use `COPY` for bulk data loading
- Monitor with pg_stat_statements extension
- Use logical replication for read replicas
- Implement audit logging with triggers or extensions

```sql
-- RETURNING clause
INSERT INTO users (email, first_name, last_name)
VALUES ('new@example.com', 'New', 'User')
RETURNING id, created_at;

-- UPSERT (ON CONFLICT)
INSERT INTO users (email, first_name, last_name)
VALUES ('existing@example.com', 'Updated', 'Name')
ON CONFLICT (email) 
DO UPDATE SET 
    first_name = EXCLUDED.first_name,
    last_name = EXCLUDED.last_name,
    updated_at = NOW();

-- Bulk insert with COPY
COPY users (email, first_name, last_name) 
FROM '/path/to/data.csv' 
WITH (FORMAT csv, HEADER true);
```
