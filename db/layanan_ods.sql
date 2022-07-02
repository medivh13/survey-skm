CREATE TABLE master.layanan_opds (
    id bigint NOT NULL,
    opd_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL,
    unsur_pelayanan character varying(255) DEFAULT ''::character varying NOT NULL,
    CONSTRAINT fk_opds
      FOREIGN KEY(opd_id) 
	  REFERENCES master.opds(id)
);

CREATE SEQUENCE master.layanan_opds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.layanan_opds_id_seq OWNED BY master.layanan_opds.id;

ALTER TABLE ONLY master.layanan_opds ALTER COLUMN id SET DEFAULT nextval('master.layanan_opds_id_seq'::regclass);

ALTER TABLE ONLY master.layanan_opds
    ADD CONSTRAINT layanan_opds_pkey PRIMARY KEY (id);

CREATE INDEX layanan_opds_index ON master.layanan_opds USING btree (deleted_at, id, opd_id);


INSERT INTO master.layanan_opds (display_name, unsur_pelayanan, opd_id) VALUES ('test1', 'test1', 1)
INSERT INTO master.layanan_opds (display_name, unsur_pelayanan, opd_id) VALUES ('test2', 'test2', 2)