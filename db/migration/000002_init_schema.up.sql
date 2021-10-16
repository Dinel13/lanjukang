CREATE TABLE "bookings" (
  "id" bigserial PRIMARY KEY,
  "user_id" INT NOT NULL,
  "service_id" INT NOT NULL,
  "owner_id" INT NOT NULL,
  "transaction_id" INT DEFAULT NULL,
  "amount" INT NOT NULL DEFAULT 1,
  "start_at" timestamptz NOT NULL DEFAULT (now()),
  "end_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "bookings"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "bookings"
ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id") ON DELETE CASCADE;

ALTER TABLE "bookings"
ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON DELETE CASCADE;


ALTER TABLE "bookings"
ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id") ON DELETE CASCADE;


CREATE INDEX ON "bookings" ("id");
CREATE INDEX ON "transactions" ("id");
CREATE INDEX ON "bookings" ("owner_id", "user_id",  "service_id", "transaction_id" );

