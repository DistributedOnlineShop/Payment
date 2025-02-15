ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS payments_transaction_id_fkey;

DROP TABLE IF EXISTS "payments";
DROP TABLE IF EXISTS "transactions";