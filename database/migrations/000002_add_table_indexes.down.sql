-- Drop indexes for sign_requests
DROP INDEX IF EXISTS idx_sign_requests_status;

DROP INDEX IF EXISTS idx_sign_requests_token;

-- Drop indexes for network_interfaces
DROP INDEX IF EXISTS idx_network_interfaces_ip;