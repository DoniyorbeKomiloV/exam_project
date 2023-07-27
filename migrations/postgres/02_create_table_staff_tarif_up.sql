CREATE TABLE "staff_tarif"(
    "id" uuid PRIMARY KEY,
    "name" VARCHAR(36) NOT NULL,
    "type" VARCHAR(36) NOT NULL DEFAULT 'fixed',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    "deleted" boolean
);