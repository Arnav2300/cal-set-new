CREATE TABLE "users" (
  "id" integer PRIMARY KEY NOT NULL,
  "email" varchar UNIQUE,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar,
  "role" varchar NOT NULL,
  "updated_at" timestamp,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "provider_tokens" (
  "id" integer PRIMARY KEY NOT NULL,
  "user_id" integer NOT NULL,
  "provider" varchar NOT NULL,
  "provider_id" varchar NOT NULL,
  "access_token" varchar,
  "refresh_token" varchar,
  "token_expiry" timestamp,
  "connected_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "users" ("created_at");

CREATE INDEX ON "provider_tokens" ("user_id");

CREATE INDEX ON "provider_tokens" ("provider");

CREATE INDEX ON "provider_tokens" ("provider_id");

CREATE INDEX ON "provider_tokens" ("connected_at");

ALTER TABLE "provider_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
