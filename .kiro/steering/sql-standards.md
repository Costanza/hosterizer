# SQL Coding Standards

## General Principles

- Write readable, maintainable SQL
- Use consistent formatting and naming conventions
- Optimize for performance and scalability
- Follow database-specific best practices
- Document complex queries and business logic

## Naming Conventions

- Tables: `plural_snake_case` (e.g., `users`, `order_items`)
- Columns: `snake_case` (e.g., `first_name`, `created_at`)
- Primary keys: `id` or `{table_name}_id`
- Foreign keys: `{referenced_table}_id` (e.g., `user_id`, `order_id`)
- Indexes: `idx_{table}_{columns}` (e.g., `idx_users_email`)
- Constraints: `{type}_{table}_{columns}` (e.g., `fk_orders_user_id`, `uk_users_email`)

## Formatting

- Use uppercase for SQL keywords: `SELECT`, `FROM`, `WHERE`, `JOIN`
- Use lowercase for table and column names
- One clause per line for complex queries
- Indent subqueries and nested logic
- Use meaningful aliases

```sql
-- Good: readable formatting
SELECT 
    u.id,
    u.first_name,
    u.last_name,
    o.order_date,
    o.total_amount
FROM users u
INNER JOIN orders o ON u.id = o.user_id
WHERE o.order_date >= '2024-01-01'
    AND o.status = 'completed'
ORDER BY o.order_date DESC;

-- Bad: hard to read
select u.id,u.first_name,u.last_name,o.order_date,o.total_amount from users u inner join orders o on u.id=o.user_id where o.order_date>='2024-01-01' and o.status='completed' order by o.order_date desc;
```

## Table Design

- Use appropriate data types (avoid VARCHAR(MAX) or TEXT unless necessary)
- Always define primary keys
- Use foreign keys to enforce referential integrity
- Add NOT NULL constraints where appropriate
- Use CHECK constraints for data validation
- Include audit columns: `created_at`, `updated_at`, `created_by`, `updated_by`

```sql
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT chk_status CHECK (status IN ('active', 'inactive', 'suspended'))
);

CREATE TABLE orders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    order_date DATE NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_orders_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT chk_total_amount CHECK (total_amount >= 0)
);
```

## Indexing

- Index foreign keys
- Index columns used in WHERE, JOIN, and ORDER BY clauses
- Use composite indexes for multi-column queries
- Avoid over-indexing (impacts INSERT/UPDATE performance)
- Monitor and remove unused indexes

```sql
-- Single column index
CREATE INDEX idx_users_email ON users(email);

-- Composite index
CREATE INDEX idx_orders_user_date ON orders(user_id, order_date);

-- Unique index
CREATE UNIQUE INDEX uk_users_email ON users(email);
```

## Query Optimization

- Use `EXPLAIN` or `EXPLAIN ANALYZE` to understand query execution
- Select only needed columns (avoid `SELECT *`)
- Use appropriate JOIN types
- Filter early with WHERE clauses
- Use EXISTS instead of IN for subqueries when checking existence
- Avoid functions on indexed columns in WHERE clauses
- Use LIMIT for pagination

```sql
-- Good: specific columns
SELECT id, first_name, last_name 
FROM users 
WHERE status = 'active';

-- Bad: unnecessary columns
SELECT * FROM users WHERE status = 'active';

-- Good: EXISTS for existence check
SELECT u.id, u.first_name
FROM users u
WHERE EXISTS (
    SELECT 1 FROM orders o WHERE o.user_id = u.id
);

-- Less efficient: IN with subquery
SELECT u.id, u.first_name
FROM users u
WHERE u.id IN (SELECT user_id FROM orders);
```

## Joins

- Use explicit JOIN syntax (not implicit joins in WHERE)
- Specify JOIN type: INNER, LEFT, RIGHT, FULL OUTER
- Use table aliases for readability
- Join on indexed columns when possible

```sql
-- Good: explicit INNER JOIN
SELECT 
    u.first_name,
    o.order_date,
    o.total_amount
FROM users u
INNER JOIN orders o ON u.id = o.user_id
WHERE u.status = 'active';

-- Bad: implicit join
SELECT 
    u.first_name,
    o.order_date,
    o.total_amount
FROM users u, orders o
WHERE u.id = o.user_id
    AND u.status = 'active';
```

## Subqueries and CTEs

- Use Common Table Expressions (CTEs) for complex queries
- CTEs improve readability over nested subqueries
- Use subqueries in FROM clause sparingly (can impact performance)

```sql
-- Good: CTE for readability
WITH active_users AS (
    SELECT id, first_name, last_name
    FROM users
    WHERE status = 'active'
),
recent_orders AS (
    SELECT user_id, COUNT(*) as order_count
    FROM orders
    WHERE order_date >= DATE_SUB(CURRENT_DATE, INTERVAL 30 DAY)
    GROUP BY user_id
)
SELECT 
    au.first_name,
    au.last_name,
    COALESCE(ro.order_count, 0) as recent_orders
FROM active_users au
LEFT JOIN recent_orders ro ON au.id = ro.user_id;
```

## Transactions

- Use transactions for multi-statement operations
- Keep transactions short to avoid locking issues
- Handle errors and rollback appropriately
- Use appropriate isolation levels

```sql
START TRANSACTION;

INSERT INTO orders (user_id, order_date, total_amount, status)
VALUES (123, CURRENT_DATE, 99.99, 'pending');

SET @order_id = LAST_INSERT_ID();

INSERT INTO order_items (order_id, product_id, quantity, price)
VALUES (@order_id, 456, 2, 49.99);

COMMIT;
```

## Aggregations

- Use GROUP BY with aggregate functions
- Use HAVING for filtering aggregated results
- Consider performance impact of aggregations on large datasets

```sql
SELECT 
    u.id,
    u.first_name,
    COUNT(o.id) as order_count,
    SUM(o.total_amount) as total_spent
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
GROUP BY u.id, u.first_name
HAVING COUNT(o.id) > 5
ORDER BY total_spent DESC;
```

## NULL Handling

- Use IS NULL and IS NOT NULL (not = NULL)
- Use COALESCE for default values
- Be aware of NULL behavior in comparisons and aggregations

```sql
-- Good: proper NULL check
SELECT * FROM users WHERE middle_name IS NULL;

-- Bad: incorrect NULL check
SELECT * FROM users WHERE middle_name = NULL;

-- Use COALESCE for defaults
SELECT 
    first_name,
    COALESCE(middle_name, '') as middle_name,
    last_name
FROM users;
```

## Security

- Use parameterized queries to prevent SQL injection
- Never concatenate user input into SQL strings
- Grant minimum necessary privileges
- Encrypt sensitive data at rest and in transit
- Use stored procedures for complex operations
- Audit and log database access

```sql
-- Application code should use parameterized queries
-- Good (pseudocode):
-- query = "SELECT * FROM users WHERE email = ?"
-- execute(query, [user_email])

-- Bad (vulnerable to SQL injection):
-- query = "SELECT * FROM users WHERE email = '" + user_email + "'"
```

## Best Practices

- Use meaningful table and column names
- Document complex queries with comments
- Avoid SELECT * in production code
- Use appropriate data types (don't use VARCHAR for numbers)
- Normalize data to reduce redundancy (usually 3NF)
- Denormalize strategically for performance when needed
- Use views for frequently used complex queries
- Implement soft deletes with a `deleted_at` column instead of hard deletes
- Use database migrations for schema changes
- Test queries with realistic data volumes

```sql
-- Soft delete pattern
UPDATE users 
SET deleted_at = CURRENT_TIMESTAMP 
WHERE id = 123;

-- Query excluding soft-deleted records
SELECT * FROM users WHERE deleted_at IS NULL;
```

## Date and Time

- Use appropriate date/time types: DATE, TIME, DATETIME, TIMESTAMP
- Store timestamps in UTC
- Use database functions for date arithmetic
- Be consistent with time zones

```sql
-- Date arithmetic
SELECT * FROM orders 
WHERE order_date >= DATE_SUB(CURRENT_DATE, INTERVAL 30 DAY);

-- Extract date parts
SELECT 
    YEAR(order_date) as year,
    MONTH(order_date) as month,
    COUNT(*) as order_count
FROM orders
GROUP BY YEAR(order_date), MONTH(order_date);
```
