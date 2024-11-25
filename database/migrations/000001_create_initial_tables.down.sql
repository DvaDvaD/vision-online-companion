-- Drop table for sign_submissions (no other tables depend on this table)
DROP TABLE IF EXISTS sign_submissions;

-- Drop table for network_interfaces (depends on sign_requests)
DROP TABLE IF EXISTS network_interfaces;

-- Drop table for issues (depends on sign_requests)
DROP TABLE IF EXISTS issues;

-- Drop table for sign_requests (the main table, independent)
DROP TABLE IF EXISTS sign_requests;