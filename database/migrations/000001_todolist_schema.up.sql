CREATE TABLE "tasks" (

    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR (100) NOT NULL,
    "descryption" VARCHAR (256),
    "created_at" timestamp,
    "updated_at" timestamp,
    "is_finished" boolean,

    "parents_id" integer,
    "priority_id" integer
);

CREATE TABLE "priorities" (
    "id" SERIAL PRIMARY KEY,
    "priority" VARCHAR (50) NOT NULL
);

CREATE TABLE "files" (
    "id" SERIAL PRIMARY KEY,
    "extension" VARCHAR (5) NOT NULL,
    "filename" VARCHAR (100) NOT NULL,
    "task_id" integer
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("priority_id") REFERENCES "priorities" ("id");
ALTER TABLE "files" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id");