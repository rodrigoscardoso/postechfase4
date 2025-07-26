CREATE TYPE "orders_status_enum" AS ENUM ('payment_pending', 'received', 'preparing', 'ready', 'completed');
CREATE SEQUENCE order_number_seq START 1;


CREATE TABLE IF NOT EXISTS "orders" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"number" INT NOT NULL DEFAULT nextval('order_number_seq'),
    "status" orders_status_enum DEFAULT 'payment_pending',
	"client_id" uuid NULL,
	"payment_id" uuid NULL,
	"total" numeric(10, 2) NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);

ALTER TABLE "orders"
      ADD CONSTRAINT fk_orders_client FOREIGN KEY (client_id) 
          REFERENCES "clients" (id);