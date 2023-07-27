CREATE TABLE "sales"(
    "id" uuid PRIMARY KEY,
    "branch_id" uuid NOT NULL REFERENCES branch("id"),
    "shop_assistant_id" uuid REFERENCES staff("id"),
    "cashier_id" uuid NOT NULL REFERENCES staff("id") ,
    "price" NUMERIC NOT NULL,
    "payment_type" VARCHAR NOT NULL,
    "status" VARCHAR NOT NULL,
    "client_name" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted" BOOLEAN,
    "deleted_at" TIMESTAMP
);