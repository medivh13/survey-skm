CREATE TABLE master.kat_opds (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL
);

CREATE SEQUENCE master.kat_opds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.kat_opds_id_seq OWNED BY master.kat_opds.id;

ALTER TABLE ONLY master.kat_opds ALTER COLUMN id SET DEFAULT nextval('master.kat_opds_id_seq'::regclass);

ALTER TABLE ONLY master.kat_opds
    ADD CONSTRAINT kat_opds_pkey PRIMARY KEY (id);


CREATE INDEX kat_opds_index ON master.kat_opds USING btree (deleted_at, id);

INSERT INTO master.kat_opds (display_name) VALUES ('test1');
INSERT INTO master.kat_opds (display_name) VALUES ('test2');