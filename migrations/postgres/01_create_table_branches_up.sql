CREATE TABLE "branch"(
    "id" uuid PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "address" VARCHAR(128),
    "phone_number" VARCHAR(32),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    "deleted" boolean
);