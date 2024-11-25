-- Table to store sign requests
CREATE TABLE sign_requests (
    request_id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid (),
    request_user VARCHAR(255) NOT NULL,
    holder_id VARCHAR(255) NOT NULL,
    holder_name VARCHAR(255) NOT NULL,
    notes TEXT,
    status VARCHAR(20) CHECK (
        status IN (
            'pending',
            'expired',
            'failed',
            'success'
        )
    ) DEFAULT 'pending',
    token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        signed_at TIMESTAMP
    WITH
        TIME ZONE
);

-- Table to store issues related to a sign request
CREATE TABLE issues (
    id SERIAL PRIMARY KEY,
    sign_request_id VARCHAR(36) NOT NULL REFERENCES sign_requests (request_id) ON DELETE CASCADE,
    number VARCHAR(255) NOT NULL,
    copy INTEGER NOT NULL,
    description TEXT
);

-- Table to store network interfaces associated with sign requests
CREATE TABLE network_interfaces (
    id SERIAL PRIMARY KEY,
    sign_request_id VARCHAR(36) NOT NULL REFERENCES sign_requests (request_id) ON DELETE CASCADE,
    name VARCHAR(50),
    ip VARCHAR(50),
    dsn VARCHAR(255)
);

-- Table to store sign submissions (includes sign data and location)
CREATE TABLE sign_submissions (
    id SERIAL PRIMARY KEY,
    sign_request_id VARCHAR(36) NOT NULL REFERENCES sign_requests (request_id) ON DELETE CASCADE,
    sign BYTEA NOT NULL, -- storing base64-encoded data
    location_latitude DOUBLE PRECISION,
    location_longitude DOUBLE PRECISION,
    submitted_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);