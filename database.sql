-- gorm can create tables: 	db.Table("time_track").CreateTable(&TimeTrack{})

CREATE TABLE public.time_track
(
  id integer NOT NULL DEFAULT nextval('time_track_id_seq'::regclass),
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  method_name text,
  app_instance text,
  "time" bigint,
  date timestamp with time zone,
  result boolean,
  reason_code text,
  CONSTRAINT time_track_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.time_track
  OWNER TO postgres;

-- Index: public.idx_time_track_deleted_at

-- DROP INDEX public.idx_time_track_deleted_at;

CREATE INDEX idx_time_track_deleted_at
  ON public.time_track
  USING btree
  (deleted_at);
