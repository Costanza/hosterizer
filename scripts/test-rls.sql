-- Test script to validate Row-Level Security (RLS) for multi-tenancy
-- Run this script after migrations to verify tenant isolation
\ echo '=== Testing Row-Level Security ===' \ echo '' -- Get customer IDs for testing
\ echo 'Getting customer IDs...'
SELECT id,
    name
FROM customers
ORDER BY id;
\ echo '' -- Test 1: Query sites as Customer 1
\ echo 'Test 1: Query sites as Customer 1' \ echo 'Expected: Only sites for Customer One Inc'
SET app.current_customer_id = (
        SELECT id
        FROM customers
        WHERE name = 'Customer One Inc'
    );
SET app.current_user_role = 'customer';
SELECT id,
    name,
    customer_id
FROM sites;
\ echo '' -- Test 2: Query sites as Customer 2
\ echo 'Test 2: Query sites as Customer 2' \ echo 'Expected: Only sites for Customer Two LLC'
SET app.current_customer_id = (
        SELECT id
        FROM customers
        WHERE name = 'Customer Two LLC'
    );
SET app.current_user_role = 'customer';
SELECT id,
    name,
    customer_id
FROM sites;
\ echo '' -- Test 3: Query sites as Administrator
\ echo 'Test 3: Query sites as Administrator' \ echo 'Expected: All sites' RESET app.current_customer_id;
SET app.current_user_role = 'administrator';
SELECT id,
    name,
    customer_id
FROM sites;
\ echo '' -- Test 4: Query cost records as Customer 1
\ echo 'Test 4: Query cost records as Customer 1' \ echo 'Expected: Only cost records for Customer One Inc'
SET app.current_customer_id = (
        SELECT id
        FROM customers
        WHERE name = 'Customer One Inc'
    );
SET app.current_user_role = 'customer';
SELECT id,
    customer_id,
    cloud_provider,
    amount
FROM cost_records;
\ echo '' -- Test 5: Query cost records as Customer 2
\ echo 'Test 5: Query cost records as Customer 2' \ echo 'Expected: Only cost records for Customer Two LLC'
SET app.current_customer_id = (
        SELECT id
        FROM customers
        WHERE name = 'Customer Two LLC'
    );
SET app.current_user_role = 'customer';
SELECT id,
    customer_id,
    cloud_provider,
    amount
FROM cost_records;
\ echo '' -- Test 6: Query cost records as Administrator
\ echo 'Test 6: Query cost records as Administrator' \ echo 'Expected: All cost records' RESET app.current_customer_id;
SET app.current_user_role = 'administrator';
SELECT id,
    customer_id,
    cloud_provider,
    amount
FROM cost_records;
\ echo '' -- Test 7: Query deployments as Customer 1
\ echo 'Test 7: Query deployments as Customer 1' \ echo 'Expected: Only deployments for Customer One Inc sites'
SET app.current_customer_id = (
        SELECT id
        FROM customers
        WHERE name = 'Customer One Inc'
    );
SET app.current_user_role = 'customer';
SELECT d.id,
    d.site_id,
    s.name as site_name,
    d.deployment_type,
    d.status
FROM deployments d
    JOIN sites s ON d.site_id = s.id;
\ echo '' -- Test 8: Query deployments as Administrator
\ echo 'Test 8: Query deployments as Administrator' \ echo 'Expected: All deployments' RESET app.current_customer_id;
SET app.current_user_role = 'administrator';
SELECT d.id,
    d.site_id,
    s.name as site_name,
    d.deployment_type,
    d.status
FROM deployments d
    JOIN sites s ON d.site_id = s.id;
\ echo '' -- Test 9: Query customers as Customer 1
\ echo 'Test 9: Query customers as Customer 1' \ echo 'Expected: Only Customer One Inc record'
SET app.current_customer_id = (
        SELECT id
        FROM customers
        WHERE name = 'Customer One Inc'
    );
SET app.current_user_role = 'customer';
SELECT id,
    name,
    tier
FROM customers;
\ echo '' -- Test 10: Query customers as Administrator
\ echo 'Test 10: Query customers as Administrator' \ echo 'Expected: All customers' RESET app.current_customer_id;
SET app.current_user_role = 'administrator';
SELECT id,
    name,
    tier
FROM customers;
\ echo '' -- Reset session variables
RESET app.current_customer_id;
RESET app.current_user_role;
\ echo '=== RLS Testing Complete ==='