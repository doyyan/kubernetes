CREATE TABLE "deployments" (
                             "id" bigserial PRIMARY KEY,
                             "namespace" varchar NOT NULL,
                             "name" varchar NOT NULL,
                             "kind" varchar NOT NULL,
                             "labels" text[],
                             "replicas" integer,
                             "status" varchar,
                             "desired" int,
                             "ready" int,
                             "current" int,
                             "available" int,
                             "created_at" bigint NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                             "updated_at" bigint NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);
