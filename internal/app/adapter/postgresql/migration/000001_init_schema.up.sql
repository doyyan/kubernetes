CREATE TABLE "deployments" (
                             "id" bigserial PRIMARY KEY,
                             "namespace" varchar NOT NULL,
                             "name" varchar NOT NULL,
                             "kind" varchar NOT NULL,
                             "image" varchar NOT NULL,
                             "containerport" integer NOT NULL,
                             "containername" varchar NOT NULL,
                             "labels" jsonb,
                             "replicas" integer,
                             "status" varchar,
                             "desired" int,
                             "ready" int,
                             "current" int,
                             "available" int,
                             "created_at" bigint NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
                             "updated_at" bigint NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);
