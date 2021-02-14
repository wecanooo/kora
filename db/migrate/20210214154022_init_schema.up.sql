CREATE TABLE "admins" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "encrypted_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "username" varchar NOT NULL,
  "encrypted_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "owner_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "published_at" timestamptz
);

CREATE TABLE "relationships" (
  "follower_id" bigint NOT NULL,
  "followed_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "responses" (
  "id" bigserial PRIMARY KEY,
  "body" text NOT NULL,
  "post_id" bigint NOT NULL,
  "owner_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "posts" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "relationships" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "relationships" ADD FOREIGN KEY ("followed_id") REFERENCES "users" ("id");

ALTER TABLE "responses" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "responses" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

CREATE INDEX ON "admins" ("email");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("owner_id");

CREATE INDEX ON "relationships" ("follower_id");

CREATE INDEX ON "relationships" ("followed_id");

CREATE INDEX ON "relationships" ("follower_id", "followed_id");

CREATE INDEX ON "responses" ("post_id");

CREATE INDEX ON "responses" ("owner_id");

COMMENT ON COLUMN "users"."email" IS '이메일은 필수 항목입니다.';

COMMENT ON COLUMN "users"."username" IS '사용자명은 필수 항목입니다.';

COMMENT ON COLUMN "posts"."body" IS '포스트는 비어 있을 수 없습니다.';

COMMENT ON COLUMN "responses"."body" IS '답변은 반드시 내용이 있어야 합니다.';