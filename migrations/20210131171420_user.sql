-- migrate:up

CREATE TABLE public.user (
    "id" SERIAL PRIMARY KEY,
    "fullname" TEXT,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "salt" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT now() NOT NULL,
    "deleted_at" TIMESTAMPTZ,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX unique_email_idx ON public.user ("email");


-- migrate:down

DROP TABLE public.user;
