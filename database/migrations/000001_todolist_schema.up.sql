CREATE TABLE "Tasks" (

    "task_id" SERIAL PRIMARY KEY,
    "title" VARCHAR (100) NOT NULL,
    "descryption" VARCHAR (256),
    "createdAt" timestamp,
    "updateAt" timestamp,
    "isFinished" boolean,

    "parents_id" integer,
    "priority_id" integer
);

CREATE TABLE "Priorities" (
    "id" SERIAL PRIMARY KEY,
    "priority" VARCHAR (50) NOT NULL
);

CREATE TABLE "Files" (
    "id" SERIAL PRIMARY KEY,
    "extension" VARCHAR (5) NOT NULL,
    "filename" VARCHAR (100) NOT NULL,
    "task_id" integer
);

ALTER TABLE "Tasks" ADD FOREIGN KEY ("priority_id") REFERENCES "Priorities" ("id");
ALTER TABLE "Tasks" ADD FOREIGN KEY ("parents_id") REFERENCES "Tasks" ("id");
ALTER TABLE "Files" ADD FOREIGN KEY ("task_id") REFERENCES "Tasks" ("task_id");