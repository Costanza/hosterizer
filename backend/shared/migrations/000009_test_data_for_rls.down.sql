-- Remove test data for RLS validation
-- Delete test cost records
DELETE FROM cost_records
WHERE customer_id IN (
        SELECT id
        FROM customers
        WHERE name IN ('Customer One Inc', 'Customer Two LLC')
    );
-- Delete test deployments
DELETE FROM deployments
WHERE site_id IN (
        SELECT id
        FROM sites
        WHERE customer_id IN (
                SELECT id
                FROM customers
                WHERE name IN ('Customer One Inc', 'Customer Two LLC')
            )
    );
-- Delete test sites
DELETE FROM sites
WHERE customer_id IN (
        SELECT id
        FROM customers
        WHERE name IN ('Customer One Inc', 'Customer Two LLC')
    );
-- Delete test customers
DELETE FROM customers
WHERE name IN ('Customer One Inc', 'Customer Two LLC');
-- Delete test users
DELETE FROM users
WHERE email IN (
        'admin@hosterizer.com',
        'customer1@example.com',
        'customer2@example.com'
    );