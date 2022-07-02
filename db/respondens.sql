CREATE TABLE survey.respondens (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    display_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email_address character varying(255) NOT NULL,
    umur bigint NOT NULL,
    pekerjaan_id bigint NOT NULL,
    pendidikan_id bigint NOT NULL,
    layanan_id bigint NOT NULL,
    CONSTRAINT fk_pekerjaan
      FOREIGN KEY(pekerjaan_id) 
	  REFERENCES master.pekerjaans(id),
    CONSTRAINT fk_pendidikan
      FOREIGN KEY(pendidikan_id) 
	  REFERENCES master.pendidikans(id),
    CONSTRAINT fk_layanan
      FOREIGN KEY(layanan_id) 
	  REFERENCES master.layanan_opds(id)
);

CREATE SEQUENCE survey.respondens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE survey.respondens_id_seq OWNED BY survey.respondens.id;

ALTER TABLE ONLY survey.respondens ALTER COLUMN id SET DEFAULT nextval('survey.respondens_id_seq'::regclass);

ALTER TABLE ONLY survey.respondens
    ADD CONSTRAINT respondens_pkey PRIMARY KEY (id);

ALTER TABLE ONLY survey.respondens
    ADD CONSTRAINT respondens_email_address_key UNIQUE (email_address);


CREATE INDEX respondens_index ON survey.respondens USING btree (deleted_at, id, umur, pekerjaan_id, pendidikan_id);