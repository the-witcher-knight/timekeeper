-- Set up users table
CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGINT PRIMARY KEY,
    "name"       VARCHAR      NOT NULL DEFAULT '',
    "email"      VARCHAR      NOT NULL DEFAULT '',
    "password"   VARCHAR      NOT NULL DEFAULT '',
    "role"       VARCHAR(8)[] NOT NULL DEFAULT '{EMPLOYEE}', -- EMPLOYEE or ADMIN
    "created_at" TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMPTZ
);
CREATE UNIQUE INDEX IF NOT EXISTS "email_on_users" ON "users"("email");

-- Setup attendance table
CREATE TABLE IF NOT EXISTS "attendances" (
    "id" BIGINT PRIMARY KEY,
    "employer_id"   BIGINT      NOT NULL,
    "check_in_time" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "notes"         VARCHAR     NOT NULL DEFAULT '',
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "deleted_at"    TIMESTAMPTZ,
    FOREIGN KEY ("employer_id") REFERENCES "users" ("id")
);
