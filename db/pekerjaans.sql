CREATE TABLE master.pekerjaans (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE SEQUENCE master.pekerjaans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.pekerjaans_id_seq OWNED BY master.pekerjaans.id;

ALTER TABLE ONLY master.pekerjaans ALTER COLUMN id SET DEFAULT nextval('master.pekerjaans_id_seq'::regclass);

ALTER TABLE ONLY master.pekerjaans
    ADD CONSTRAINT pekerjaans_pkey PRIMARY KEY (id);


CREATE INDEX pekerjaans_index ON master.pekerjaans USING btree (deleted_at, id);