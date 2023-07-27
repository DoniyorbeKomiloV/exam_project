CREATE TABLE "staff"(
    "id" uuid PRIMARY KEY,
    "branch_id" uuid REFERENCES branch("id"),
    "tarif_id" uuid REFERENCES staff_tarif("id"),
    "type" VARCHAR(16) NOT NULL,
    "name" VARCHAR(64) NOT NULL,
    "balance" NUMERIC DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    "deleted" boolean
);