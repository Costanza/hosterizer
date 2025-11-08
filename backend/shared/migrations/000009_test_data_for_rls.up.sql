-- Test data for Row-Level Security validation
-- This migration creates test data to verify tenant isolation
-- Insert test users
INSERT INTO users (
        email,
        password_hash,
        first_name,
        last_name,
        role
    )
VALUES (
        'admin@hosterizer.com',
        '$2a$10$test_hash_admin',
        'Admin',
        'User',
        'administrator'
    ),
    (
        'customer1@example.com',
        '$2a$10$test_hash_customer1',
        'Customer',
        'One',
        'customer'
    ),
    (
        'customer2@example.com',
        '$2a$10$test_hash_customer2',
        'Customer',
        'Two',
        'customer'
    );
-- Insert test customers
INSERT INTO customers (name, tier, status, owner_user_id)
VALUES (
        'Customer One Inc',
        'standard',
        'active',
        (
            SELECT id
            FROM users
            WHERE email = 'customer1@example.com'
        )
    ),
    (
        'Customer Two LLC',
        'premium',
        'active',
        (
            SELECT id
            FROM users
            WHERE email = 'customer2@example.com'
        )
    );
-- Insert test sites for customer 1
INSERT INTO sites (
        customer_id,
        name,
        domain,
        cloud_provider,
        region,
        status
    )
VALUES (
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer One Inc'
        ),
        'site1-customer1',
        'site1.customer1.com',
        'aws',
        'us-east-1',
        'active'
    ),
    (
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer One Inc'
        ),
        'site2-customer1',
        'site2.customer1.com',
        'azure',
        'eastus',
        'active'
    );
-- Insert test sites for customer 2
INSERT INTO sites (
        customer_id,
        name,
        domain,
        cloud_provider,
        region,
        status
    )
VALUES (
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer Two LLC'
        ),
        'site1-customer2',
        'site1.customer2.com',
        'gcp',
        'us-central1',
        'active'
    ),
    (
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer Two LLC'
        ),
        'site2-customer2',
        'site2.customer2.com',
        'digitalocean',
        'nyc1',
        'active'
    );
-- Insert test deployments
INSERT INTO deployments (site_id, deployment_type, status)
VALUES (
        (
            SELECT id
            FROM sites
            WHERE name = 'site1-customer1'
        ),
        'create',
        'completed'
    ),
    (
        (
            SELECT id
            FROM sites
            WHERE name = 'site2-customer1'
        ),
        'create',
        'completed'
    ),
    (
        (
            SELECT id
            FROM sites
            WHERE name = 'site1-customer2'
        ),
        'create',
        'completed'
    ),
    (
        (
            SELECT id
            FROM sites
            WHERE name = 'site2-customer2'
        ),
        'create',
        'completed'
    );
-- Insert test cost records
INSERT INTO cost_records (
        site_id,
        customer_id,
        cloud_provider,
        cost_date,
        amount
    )
VALUES (
        (
            SELECT id
            FROM sites
            WHERE name = 'site1-customer1'
        ),
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer One Inc'
        ),
        'aws',
        CURRENT_DATE,
        50.00
    ),
    (
        (
            SELECT id
            FROM sites
            WHERE name = 'site2-customer1'
        ),
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer One Inc'
        ),
        'azure',
        CURRENT_DATE,
        75.00
    ),
    (
        (
            SELECT id
            FROM sites
            WHERE name = 'site1-customer2'
        ),
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer Two LLC'
        ),
        'gcp',
        CURRENT_DATE,
        100.00
    ),
    (
        (
            SELECT id
            FROM sites
            WHERE name = 'site2-customer2'
        ),
        (
            SELECT id
            FROM customers
            WHERE name = 'Customer Two LLC'
        ),
        'digitalocean',
        CURRENT_DATE,
        25.00
    );
-- Add comment
COMMENT ON TABLE users IS 'Test data added for RLS validation';