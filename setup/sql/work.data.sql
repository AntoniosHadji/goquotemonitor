--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Debian 15.1-1.pgdg110+1)
-- Dumped by pg_dump version 15.1 (Debian 15.1-1.pgdg110+1)

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
-- Data for Name: work; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.work (lp, ticker, size) FROM stdin;
Enigma	BTC	1.000
Enigma	BTC	0.010
Enigma	ETH	1.000
DV	ETH	1.000
DV	BTC	1.000
DV	BTC	0.010
DV	LTC	1.000
DV	AVAX	10.000
DV	SOL	10.000
DV	USDC	100.000
Enigma	USDC	100.000
DV	USDT	100.000
Enigma	USDT	100.000
Coinbase	BTC	1.000
Coinbase	ETH	1.000
\.


--
-- PostgreSQL database dump complete
--
-- added USDC USDT Enigma manually 05/24/23
