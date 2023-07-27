CREATE TABLE "staff_transaction"(
    "id" uuid PRIMARY KEY,
    "sales_id" uuid NOT NULL REFERENCES sales("id"),
    "type" VARCHAR NOT NULL,
    "source_type" VARCHAR NOT NULL,
    "text" TEXT,
    "amount" NUMERIC NOT NULL,
    "staff_id" uuid NOT NULL REFERENCES staff("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN,
    "deleted_at" TIMESTAMP
);