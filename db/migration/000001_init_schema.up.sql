CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "nick_name" varchar DEFAULT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" int NOT NULL DEFAULT 0,
  "verified" boolean NOT NULL DEFAULT false,
  "image" varchar DEFAULT NULL,
  "phone" varchar DEFAULT NULL,
  "address" varchar DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "services" (
  "id" bigserial PRIMARY KEY,
  "owner_id" int NOT NULL,
  "name" varchar NOT NULL,
  "price" int NOT NULL,
  "image" varchar NOT NULL,
  "type_id" int NOT NULL,
  "location" int NOT NULL,
  "capacity" int NOT NULL DEFAULT 1,
  "descriptin" TEXT NOT NULL,
  "rating" INT DEFAULT NULL,
  "comments" INT DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "type_services" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "locations" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "image" varchar DEFAULT NULL,
  "coordinate " varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "userId" INT NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE "services"
ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "services"
ADD FOREIGN KEY ("type_id") REFERENCES "type_services" ("id") ON DELETE CASCADE;

ALTER TABLE "services"
ADD FOREIGN KEY ("location") REFERENCES "locations" ("id") ON DELETE CASCADE;

ALTER TABLE "services"
ADD FOREIGN KEY ("comments") REFERENCES "comments" ("id") ON DELETE CASCADE;

ALTER TABLE "comments"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE;

CREATE INDEX ON "users" ("id");
CREATE INDEX ON "services" ("id");
CREATE INDEX ON "type_services" ("id");
CREATE INDEX ON "locations" ("id");
CREATE INDEX ON "comments" ("id");
CREATE INDEX ON "services" ("owner_id", "type_id", "location", "comments" );
CREATE INDEX ON "comments" ("userId");
