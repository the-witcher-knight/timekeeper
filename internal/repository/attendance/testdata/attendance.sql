INSERT INTO "users" (id, name, email, password, role, created_at, updated_at, deleted_at)
VALUES (1, 'user1', 'user1@mail.com', 'pwd', '{EMPLOYEE}', '2024-06-20T01:00:00+00:00', '2024-06-20T01:00:00+00:00', NULL);

INSERT INTO "attendances" (id, employer_id, check_in_time, notes, created_at, updated_at, deleted_at)
VALUES
    (1, 1, '2024-06-20T01:00:00+00:00', 'Check in', '2024-06-20T01:00:00+00:00', '2024-06-20T01:00:00+00:00', NULL),
    (2, 1, '2024-06-20T01:00:00+00:00', 'Check in', '2024-06-20T01:00:00+00:00', '2024-06-20T01:00:00+00:00', NULL),
    (3, 1, '2024-06-20T01:00:00+00:00', 'Check in', '2024-06-20T01:00:00+00:00', '2024-06-20T01:00:00+00:00', '2024-06-20T01:00:00+00:00');
