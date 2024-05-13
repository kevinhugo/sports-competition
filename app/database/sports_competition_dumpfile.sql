--
-- PostgreSQL database dump
--

-- Dumped from database version 15.0
-- Dumped by pg_dump version 15.0

-- Started on 2024-05-13 14:39:00

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
-- TOC entry 844 (class 1247 OID 17498)
-- Name: user_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.user_status AS ENUM (
    'ACTIVE',
    'INACTIVE'
);


ALTER TYPE public.user_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 17504)
-- Name: inbound_traffict_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inbound_traffict_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tracking_id character varying(50) NOT NULL,
    request_name character varying(50),
    request_host text,
    request_url text,
    request_header text NOT NULL,
    request_body text,
    request_time timestamp with time zone,
    response_status_code smallint,
    response_header text,
    response_body text,
    response_time timestamp with time zone
);


ALTER TABLE public.inbound_traffict_logs OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 17503)
-- Name: inbound_traffict_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.inbound_traffict_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.inbound_traffict_logs_id_seq OWNER TO postgres;

--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 214
-- Name: inbound_traffict_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.inbound_traffict_logs_id_seq OWNED BY public.inbound_traffict_logs.id;


--
-- TOC entry 217 (class 1259 OID 17518)
-- Name: outbound_traffict_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.outbound_traffict_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tracking_id character varying(50) NOT NULL,
    request_name character varying(100),
    request_url text,
    request_header text NOT NULL,
    request_body text,
    request_time timestamp with time zone,
    response_status_code smallint,
    response_header text,
    response_body text,
    response_time timestamp with time zone
);


ALTER TABLE public.outbound_traffict_logs OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 17517)
-- Name: outbound_traffict_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.outbound_traffict_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.outbound_traffict_logs_id_seq OWNER TO postgres;

--
-- TOC entry 3374 (class 0 OID 0)
-- Dependencies: 216
-- Name: outbound_traffict_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.outbound_traffict_logs_id_seq OWNED BY public.outbound_traffict_logs.id;


--
-- TOC entry 221 (class 1259 OID 17566)
-- Name: sports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sports (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    user_id bigint NOT NULL,
    proficiency integer NOT NULL
);


ALTER TABLE public.sports OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 17565)
-- Name: sports_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sports_id_seq OWNER TO postgres;

--
-- TOC entry 3375 (class 0 OID 0)
-- Dependencies: 220
-- Name: sports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sports_id_seq OWNED BY public.sports.id;


--
-- TOC entry 219 (class 1259 OID 17531)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    user_id character varying(50) NOT NULL,
    username character varying(50) NOT NULL,
    password text NOT NULL,
    status public.user_status DEFAULT 'ACTIVE'::public.user_status NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 17530)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3376 (class 0 OID 0)
-- Dependencies: 218
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3191 (class 2604 OID 17507)
-- Name: inbound_traffict_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inbound_traffict_logs ALTER COLUMN id SET DEFAULT nextval('public.inbound_traffict_logs_id_seq'::regclass);


--
-- TOC entry 3193 (class 2604 OID 17521)
-- Name: outbound_traffict_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.outbound_traffict_logs ALTER COLUMN id SET DEFAULT nextval('public.outbound_traffict_logs_id_seq'::regclass);


--
-- TOC entry 3199 (class 2604 OID 17569)
-- Name: sports id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sports ALTER COLUMN id SET DEFAULT nextval('public.sports_id_seq'::regclass);


--
-- TOC entry 3195 (class 2604 OID 17534)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3204 (class 2606 OID 17512)
-- Name: inbound_traffict_logs inbound_traffict_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inbound_traffict_logs
    ADD CONSTRAINT inbound_traffict_logs_pkey PRIMARY KEY (id);


--
-- TOC entry 3210 (class 2606 OID 17526)
-- Name: outbound_traffict_logs outbound_traffict_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.outbound_traffict_logs
    ADD CONSTRAINT outbound_traffict_logs_pkey PRIMARY KEY (id);


--
-- TOC entry 3222 (class 2606 OID 17573)
-- Name: sports sports_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sports
    ADD CONSTRAINT sports_pkey PRIMARY KEY (id);


--
-- TOC entry 3207 (class 2606 OID 17514)
-- Name: inbound_traffict_logs uni_inbound_traffict_logs_tracking_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inbound_traffict_logs
    ADD CONSTRAINT uni_inbound_traffict_logs_tracking_id UNIQUE (tracking_id);


--
-- TOC entry 3212 (class 2606 OID 17528)
-- Name: outbound_traffict_logs uni_outbound_traffict_logs_tracking_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.outbound_traffict_logs
    ADD CONSTRAINT uni_outbound_traffict_logs_tracking_id UNIQUE (tracking_id);


--
-- TOC entry 3224 (class 2606 OID 17575)
-- Name: sports uni_sports_user_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sports
    ADD CONSTRAINT uni_sports_user_id UNIQUE (user_id);


--
-- TOC entry 3215 (class 2606 OID 17543)
-- Name: users uni_users_user_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_user_id UNIQUE (user_id);


--
-- TOC entry 3218 (class 2606 OID 17541)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3202 (class 1259 OID 17516)
-- Name: idx_inbound_traffict_logs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inbound_traffict_logs_deleted_at ON public.inbound_traffict_logs USING btree (deleted_at);


--
-- TOC entry 3208 (class 1259 OID 17529)
-- Name: idx_outbound_traffict_logs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_outbound_traffict_logs_deleted_at ON public.outbound_traffict_logs USING btree (deleted_at);


--
-- TOC entry 3219 (class 1259 OID 17582)
-- Name: idx_sports_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sports_deleted_at ON public.sports USING btree (deleted_at);


--
-- TOC entry 3213 (class 1259 OID 17545)
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- TOC entry 3220 (class 1259 OID 17581)
-- Name: sport_user_id_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX sport_user_id_unique ON public.sports USING btree (user_id);


--
-- TOC entry 3205 (class 1259 OID 17515)
-- Name: tracking_id_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tracking_id_unique ON public.inbound_traffict_logs USING btree (tracking_id);


--
-- TOC entry 3216 (class 1259 OID 17544)
-- Name: user_id_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_id_unique ON public.users USING btree (user_id);


--
-- TOC entry 3225 (class 2606 OID 17576)
-- Name: sports fk_sports_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sports
    ADD CONSTRAINT fk_sports_user FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2024-05-13 14:39:00

--
-- PostgreSQL database dump complete
--

