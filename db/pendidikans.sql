CREATE TABLE master.pendidikans (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE SEQUENCE master.pendidikans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.pendidikans_id_seq OWNED BY master.pendidikans.id;

ALTER TABLE ONLY master.pendidikans ALTER COLUMN id SET DEFAULT nextval('master.pendidikans_id_seq'::regclass);

ALTER TABLE ONLY master.pendidikans
    ADD CONSTRAINT pendidikans_pkey PRIMARY KEY (id);


CREATE INDEX pendidikans_index ON master.pendidikans USING btree (deleted_at, id);