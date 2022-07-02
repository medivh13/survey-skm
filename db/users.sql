CREATE TABLE master.users (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    email_address character varying(255) NOT NULL,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL,
    token character varying(255) DEFAULT NULL,
    password character varying(255) NOT NULL
);

CREATE SEQUENCE master.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.users_id_seq OWNED BY master.users.id;

ALTER TABLE ONLY master.users ALTER COLUMN id SET DEFAULT nextval('master.users_id_seq'::regclass);

ALTER TABLE ONLY master.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY master.users
    ADD CONSTRAINT users_email_address_key UNIQUE (email_address);

CREATE INDEX usersindex ON master.users USING btree (deleted_at, id, email_address);

-- end user--


