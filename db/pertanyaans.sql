CREATE TABLE master.pertanyaans (
    id bigint NOT NULL,
    kat_id bigint NOT NULL,
    opd_id bigint NOT NULL,
    layanan_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    soal character varying(255) DEFAULT ''::character varying NOT NULL,
    CONSTRAINT fk_opds
      FOREIGN KEY(opd_id) 
	  REFERENCES master.opds(id),
    CONSTRAINT fk_kat_opds
      FOREIGN KEY(kat_id) 
	  REFERENCES master.kat_opds(id),
    CONSTRAINT fk_layanan_opds
      FOREIGN KEY(layanan_id) 
	  REFERENCES master.layanan_opds(id)
);

CREATE SEQUENCE master.pertanyaans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.pertanyaans_id_seq OWNED BY master.pertanyaans.id;

ALTER TABLE ONLY master.pertanyaans ALTER COLUMN id SET DEFAULT nextval('master.pertanyaans_id_seq'::regclass);

ALTER TABLE ONLY master.pertanyaans
    ADD CONSTRAINT pertanyaans_pkey PRIMARY KEY (id);

CREATE INDEX pertanyaans_index ON master.pertanyaans USING btree (deleted_at, id, opd_id, layanan_id, kat_id);


INSERT INTO master.pertanyaans (kat_id, opd_id, layanan_id, soal)
VALUES (1,1,1, 'soal 1')
INSERT INTO master.pertanyaans (kat_id, opd_id, layanan_id, soal)
VALUES (1,1,1, 'soal 2');
INSERT INTO master.pertanyaans (kat_id, opd_id, layanan_id, soal)
VALUES (1,1,1, 'soal 3');
INSERT INTO master.pertanyaans (kat_id, opd_id, layanan_id, soal)
VALUES (1,1,1, 'soal 4');