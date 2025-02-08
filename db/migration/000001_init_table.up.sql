CREATE TABLE "payments" (
  "payment_id" UUID PRIMARY KEY NOT NULL UNIQUE,
  "order_id" UUID NOT NULL,
  "user_id" UUID NOT NULL,
  "amount" DECIMAL(10,2),
  "method" VARCHAR NOT NULL,
  "status" VARCHAR NOT NULL,
  "transaction_id" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "transactions" (
  "transaction_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "vendor_id" UUID NOT NULL,
  "order_id" UUID NOT NULL,
  "type" VARCHAR NOT NULL,
  "amount" DECIMAL(10,2),
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);