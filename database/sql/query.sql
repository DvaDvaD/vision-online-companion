-- Query to get all sign requests
-- name: GetAllSignRequests :many
SELECT * FROM sign_requests;

-- Query to get a sign request by ID
-- name: GetSignRequestByID :one
SELECT * FROM sign_requests WHERE request_id = $1;

-- Query to create a new sign request
-- name: CreateSignRequest :one
INSERT INTO
    sign_requests (
        request_user,
        holder_id,
        holder_name,
        notes,
        status,
        token
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- Query to update a sign request status
-- name: UpdateSignRequestStatus :exec
UPDATE sign_requests
SET
    status = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE
    request_id = $2;

-- Query to delete a sign request
-- name: DeleteSignRequest :exec
DELETE FROM sign_requests WHERE request_id = $1;

-- Query to get all issues for a sign request
-- name: GetIssuesBySignRequestID :many
SELECT * FROM issues WHERE sign_request_id = $1;

-- Query to create a new issue
-- name: CreateIssue :one
INSERT INTO
    issues (
        sign_request_id,
        number,
        copy,
        description
    )
VALUES ($1, $2, $3, $4) RETURNING *;

-- Query to delete an issue
-- name: DeleteIssue :exec
DELETE FROM issues WHERE id = $1;

-- Query to get all network interfaces for a sign request
-- name: GetNetworkInterfacesBySignRequestID :many
SELECT * FROM network_interfaces WHERE sign_request_id = $1;

-- Query to create a new network interface
-- name: CreateNetworkInterface :one
INSERT INTO
    network_interfaces (
        sign_request_id,
        name,
        ip,
        dsn
    )
VALUES ($1, $2, $3, $4) RETURNING *;

-- Query to delete a network interface
-- name: DeleteNetworkInterface :exec
DELETE FROM network_interfaces WHERE id = $1;

-- Query to get all sign submissions for a sign request
-- name: GetSignSubmissionsBySignRequestID :many
SELECT * FROM sign_submissions WHERE sign_request_id = $1;

-- Query to create a new sign submission
-- name: CreateSignSubmission :one
INSERT INTO
    sign_submissions (
        sign_request_id,
        sign,
        location_latitude,
        location_longitude
    )
VALUES ($1, $2, $3, $4) RETURNING *;

-- Query to delete a sign submission
-- name: DeleteSignSubmission :exec
DELETE FROM sign_submissions WHERE id = $1;