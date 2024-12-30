CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "email" text UNIQUE NOT NULL,
  "username" text UNIQUE NOT NULL,
  "password" text NOT NULL,
  "role" text NOT NULL,
  "updated_at" timestamp,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "provider_tokens" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "provider" text NOT NULL,
  "provider_id" text NOT NULL,
  "access_token" text,
  "refresh_token" text,
  "token_expiry" timestamp,
  "connected_at" timestamp DEFAULT (now())
);

CREATE TABLE "password_reset_tokens" (
  "user_id" uuid PRIMARY KEY NOT NULL,
  "token" text,
  "expires_at" timestamp,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "users" ("created_at");

CREATE INDEX ON "provider_tokens" ("user_id");

CREATE INDEX ON "provider_tokens" ("provider");

CREATE INDEX ON "provider_tokens" ("provider_id");

CREATE INDEX ON "provider_tokens" ("connected_at");

CREATE INDEX ON "password_reset_tokens" ("user_id");

ALTER TABLE "provider_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "password_reset_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
