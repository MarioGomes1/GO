CREATE TABLE "users" (
"userid" SERIAL PRIMARY KEY,
"name" TEXT,
"age" INT,
"created_at" timestamptz NOT NULL DEFAULT (now()),
"updated_at" timestamptz NOT NULL DEFAULT (now())
);