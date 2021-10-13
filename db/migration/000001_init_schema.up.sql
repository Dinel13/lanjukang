CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "nick_name" varchar DEFAULT NULL,
  "email" varchar NOT NULL,
  "role" int NOT NULL DEFAULT 0,
  "image" varchar DEFAULT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bookings" (
  "id" bigserial PRIMARY KEY,
  "booker" bigint NOT NULL,
  "toko" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "kapals" (
  "id" bigserial PRIMARY KEY,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("booker") REFERENCES "users" ("id");
ALTER TABLE "bookings" ADD FOREIGN KEY ("toko") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "bookings" ("id");

CREATE INDEX ON "kapals" ("id");

CREATE INDEX ON "bookings" ("booker");

CREATE INDEX ON "bookings" ("toko");

CREATE INDEX ON "bookings" ("booker", "toko");