-- Seeding data for issues table
INSERT INTO issues (sign_request_id, number, copy, description)
VALUES
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-1'), 'ISSUE-001', 1, 'First issue for request 1'),
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-2'), 'ISSUE-002', 2, 'Second issue for request 2'),
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-3'), 'ISSUE-003', 1, 'First issue for request 3');

-- Seeding data for network_interfaces table
INSERT INTO network_interfaces (sign_request_id, name, ip, dsn)
VALUES
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-1'), 'eth0', '192.168.0.1', 'dsn1'),
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-2'), 'eth1', '192.168.0.2', 'dsn2'),
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-3'), 'wlan0', '192.168.1.1', 'dsn3');

-- Seeding data for sign_submissions table
INSERT INTO sign_submissions (sign_request_id, sign, location_latitude, location_longitude, submitted_at)
VALUES
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-1'), decode('aGVsbG93b3JsZA==', 'base64'), 37.774929, -122.419418, CURRENT_TIMESTAMP),
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-2'), decode('dGVzdHN1Ym1pc3Npb24=', 'base64'), 34.052235, -118.243683, CURRENT_TIMESTAMP - INTERVAL '1 day'),
    ((SELECT request_id FROM sign_requests WHERE token = 'unique-token-3'), decode('c3VjY2Vzc3NpZ25hdHVyZQ==', 'base64'), 40.712776, -74.005974, CURRENT_TIMESTAMP - INTERVAL '2 days');