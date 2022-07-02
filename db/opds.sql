CREATE TABLE master.opds (
    id bigint NOT NULL,
    kat_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    code character varying(255) DEFAULT ''::character varying NOT NULL,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL,
    short_name character varying(255) DEFAULT ''::character varying NOT NULL,
    CONSTRAINT fk_kat_opds
      FOREIGN KEY(kat_id) 
	  REFERENCES master.kat_opds(id)
);

CREATE SEQUENCE master.opds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.opds_id_seq OWNED BY master.opds.id;

ALTER TABLE ONLY master.opds ALTER COLUMN id SET DEFAULT nextval('master.opds_id_seq'::regclass);

ALTER TABLE ONLY master.opds
    ADD CONSTRAINT opds_pkey PRIMARY KEY (id);


CREATE INDEX opds_index ON master.opds USING btree (deleted_at, id, code, kat_id);

INSERT INTO master.opds (display_name, short_name, kat_id) VALUES ('test1', 'test1', 1)
INSERT INTO master.opds (display_name, short_name, kat_id) VALUES ('test2', 'test2', 1)
INSERT INTO master.opds (display_name, short_name, kat_id) VALUES ('test3', 'test3', 2)