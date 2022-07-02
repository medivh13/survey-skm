CREATE TABLE survey.responden_details (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    responden_id bigint NOT NULL,
    pertanyaan_id bigint NOT NULL,
    jawaban_id bigint NOT NULL,
    nilai bigint NOT NULL,
    CONSTRAINT fk_responden
      FOREIGN KEY(responden_id) 
	  REFERENCES survey.respondens(id),
    CONSTRAINT fk_pertanyaan
      FOREIGN KEY(pertanyaan_id) 
	  REFERENCES master.pertanyaans(id),
    CONSTRAINT fk_jawaban
      FOREIGN KEY(jawaban_id) 
	  REFERENCES master.jawabans(id)
);

CREATE SEQUENCE survey.responden_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE survey.responden_details_id_seq OWNED BY survey.responden_details.id;

ALTER TABLE ONLY survey.responden_details ALTER COLUMN id SET DEFAULT nextval('survey.responden_details_id_seq'::regclass);

ALTER TABLE ONLY survey.responden_details
    ADD CONSTRAINT responden_details_pkey PRIMARY KEY (id);


CREATE INDEX responden_details_index ON survey.responden_details USING btree (deleted_at, id, responden_id);