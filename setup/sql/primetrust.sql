--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Debian 15.1-1.pgdg110+1)
-- Dumped by pg_dump version 15.1 (Debian 15.1-1.pgdg110+1)

--
-- Name: primetrust; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE primetrust WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE primetrust OWNER TO postgres;

\connect primetrust

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: spreads; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.spreads (
    ts timestamp(3) with time zone,
    bid numeric(9,4),
    ask numeric(9,4),
    size numeric(6,3),
    width_bps numeric(7,3),
    ticker character varying(4),
    lp character varying(10)
);


ALTER TABLE public.spreads OWNER TO postgres;

--
-- PostgreSQL database dump complete
--
