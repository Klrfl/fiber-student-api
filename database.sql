--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: students; Type: TABLE; Schema: public; Owner: efrayanglain
--

CREATE TABLE public.students (
    id uuid NOT NULL,
    name text NOT NULL,
    major text,
    grade integer
);


ALTER TABLE public.students OWNER TO efrayanglain;

--
-- Data for Name: students; Type: TABLE DATA; Schema: public; Owner: efrayanglain
--

COPY public.students (id, name, major, grade) FROM stdin;
952cd967-a239-4c3d-8dcf-ba6faeb7fc84	Efraim Munthe	PPLG	10
2e44e075-d6ee-4586-b65d-eeabef955b0d	Langit Aulia Umbara	DKV 1	11
e7b9541d-16d3-4f80-813d-189367978093	ATha	DKV 1	11
\.


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: public; Owner: efrayanglain
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: students unique_id; Type: CONSTRAINT; Schema: public; Owner: efrayanglain
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT unique_id UNIQUE (id);


--
-- PostgreSQL database dump complete
--

