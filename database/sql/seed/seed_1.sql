-- Seeding data for sign_requests table
INSERT INTO sign_requests (request_id, request_user, holder_id, holder_name, notes, status, token, created_at, updated_at, signed_at)
VALUES
    (gen_random_uuid(), 'user1@example.com', 'holder123', 'John Doe', 'Initial request', 'pending', 'unique-token-1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (gen_random_uuid(), 'user2@example.com', 'holder456', 'Jane Smith', 'Second request', 'expired', 'unique-token-2', CURRENT_TIMESTAMP - INTERVAL '2 days', CURRENT_TIMESTAMP - INTERVAL '2 days', NULL),
    (gen_random_uuid(), 'user3@example.com', 'holder789', 'Alice Brown', 'Third request with success status', 'success', 'unique-token-3', CURRENT_TIMESTAMP - INTERVAL '1 day', CURRENT_TIMESTAMP - INTERVAL '1 day', CURRENT_TIMESTAMP - INTERVAL '1 day');