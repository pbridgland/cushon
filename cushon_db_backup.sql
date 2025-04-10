--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

-- Started on 2025-04-10 23:10:39

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4927 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 223 (class 1255 OID 16436)
-- Name: add_fund(text); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.add_fund(IN name text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO funds (name) VALUES (name);
END;
$$;


ALTER PROCEDURE public.add_fund(IN name text) OWNER TO postgres;

--
-- TOC entry 227 (class 1255 OID 16472)
-- Name: add_investment(integer, integer, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.add_investment(IN user_id integer, IN fund_id integer, IN amount integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO investments (user_id, fund_id, amount, created_at) VALUES (user_id, fund_id, amount, now());
END;
$$;


ALTER PROCEDURE public.add_investment(IN user_id integer, IN fund_id integer, IN amount integer) OWNER TO postgres;

--
-- TOC entry 225 (class 1255 OID 16448)
-- Name: add_user(text, text); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.add_user(IN username text, IN password text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO users (username, password) VALUES (username, password);
END;
$$;


ALTER PROCEDURE public.add_user(IN username text, IN password text) OWNER TO postgres;

--
-- TOC entry 224 (class 1255 OID 16437)
-- Name: get_all_funds(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_all_funds() RETURNS TABLE(id integer, name text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT funds.id, funds.name FROM funds;
END;
$$;


ALTER FUNCTION public.get_all_funds() OWNER TO postgres;

--
-- TOC entry 228 (class 1255 OID 16474)
-- Name: get_all_investments(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_all_investments() RETURNS TABLE(username text, fundname text, amount integer, created_at timestamp with time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT users.username, funds.name, investments.amount, investments.created_at FROM 
	investments inner join users on users.id = investments.user_id
	inner join funds on funds.id = investments.fund_id;
END;
$$;


ALTER FUNCTION public.get_all_investments() OWNER TO postgres;

--
-- TOC entry 226 (class 1255 OID 16454)
-- Name: get_user_by_username(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_user_by_username(p_username text) RETURNS TABLE(id integer, username text, password text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT users.id, users.username, users.password FROM users where users.username=p_username;
END;
$$;


ALTER FUNCTION public.get_user_by_username(p_username text) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 16428)
-- Name: funds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.funds (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.funds OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16427)
-- Name: funds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.funds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.funds_id_seq OWNER TO postgres;

--
-- TOC entry 4928 (class 0 OID 0)
-- Dependencies: 217
-- Name: funds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.funds_id_seq OWNED BY public.funds.id;


--
-- TOC entry 221 (class 1259 OID 16455)
-- Name: investments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.investments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.investments_id_seq OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16456)
-- Name: investments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.investments (
    id integer DEFAULT nextval('public.investments_id_seq'::regclass) NOT NULL,
    user_id integer NOT NULL,
    fund_id integer NOT NULL,
    amount integer NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.investments OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16438)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16445)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4929 (class 0 OID 0)
-- Dependencies: 220
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4758 (class 2604 OID 16431)
-- Name: funds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.funds ALTER COLUMN id SET DEFAULT nextval('public.funds_id_seq'::regclass);


--
-- TOC entry 4759 (class 2604 OID 16446)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4917 (class 0 OID 16428)
-- Dependencies: 218
-- Data for Name: funds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.funds (id, name) FROM stdin;
2	Cushon Sustainable Global Equity
3	Cushon Sustainable UK Equity
4	Cushon Sustainable Europe (ex UK) Equity
5	Cushon Sustainable Japanese Equity
6	Cushon Fixed Interes Gilts
7	Cushon Cash
\.


--
-- TOC entry 4921 (class 0 OID 16456)
-- Dependencies: 222
-- Data for Name: investments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.investments (id, user_id, fund_id, amount, created_at) FROM stdin;
1	2	3	500	2025-04-10 20:17:59.379805+01
2	2	5	2500000	2025-04-10 20:22:05.484996+01
3	2	7	50000	2025-04-10 22:54:57.106791+01
8	2	7	12312	2025-04-10 23:07:00.982332+01
\.


--
-- TOC entry 4918 (class 0 OID 16438)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password) FROM stdin;
2	admin	$2a$12$tji5COlH41QGgui3k6baKenkCHKUmuZUuzjWdzSGVt/of5gVJiLq2
\.


--
-- TOC entry 4930 (class 0 OID 0)
-- Dependencies: 217
-- Name: funds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.funds_id_seq', 7, true);


--
-- TOC entry 4931 (class 0 OID 0)
-- Dependencies: 221
-- Name: investments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.investments_id_seq', 8, true);


--
-- TOC entry 4932 (class 0 OID 0)
-- Dependencies: 220
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 4762 (class 2606 OID 16435)
-- Name: funds funds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.funds
    ADD CONSTRAINT funds_pkey PRIMARY KEY (id);


--
-- TOC entry 4768 (class 2606 OID 16476)
-- Name: investments investments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.investments
    ADD CONSTRAINT investments_pkey PRIMARY KEY (id);


--
-- TOC entry 4764 (class 2606 OID 16444)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4766 (class 2606 OID 16452)
-- Name: users users_username_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_unique UNIQUE (username);


--
-- TOC entry 4769 (class 2606 OID 16467)
-- Name: investments fund_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.investments
    ADD CONSTRAINT fund_id_fkey FOREIGN KEY (fund_id) REFERENCES public.funds(id) NOT VALID;


--
-- TOC entry 4770 (class 2606 OID 16462)
-- Name: investments user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.investments
    ADD CONSTRAINT user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) NOT VALID;


-- Completed on 2025-04-10 23:10:39

--
-- PostgreSQL database dump complete
--

