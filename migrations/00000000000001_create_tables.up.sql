CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "created_at" timestamp default now()
);

CREATE TABLE "links" (
  "id" bigserial PRIMARY KEY,
  "url" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" text,
  "user_id" bigint NOT NULL,
  "tag_id" bigint,
  "is_read" bool default false,
  "created_at" timestamp default now(),
  "updated_at" timestamp default now(),
  "deleted_at" bigint default 0
);

CREATE TABLE "tags" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "created_at" timestamp default now(),
  "updated_at" timestamp default now(),
  "deleted_at" bigint default 0
);

ALTER TABLE "links" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "links" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
