-- Indexes for efficient querying
CREATE INDEX idx_sign_requests_status ON sign_requests (status);

CREATE INDEX idx_sign_requests_token ON sign_requests (token);

CREATE INDEX idx_network_interfaces_ip ON network_interfaces (ip);