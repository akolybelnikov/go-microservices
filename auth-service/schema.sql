--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE if not exists public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq
    OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users
(
    id          integer                     DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    email       character varying(255)                                                      NOT NULL,
    first_name  character varying(255)                                                      NOT NULL,
    last_name   character varying(255)                                                      NOT NULL,
    password    character varying(60)                                                       NOT NULL,
    user_active integer                     DEFAULT 0                                       NOT NULL,
    created_at  timestamp without time zone DEFAULT now()                                   NOT NULL,
    updated_at  timestamp without time zone DEFAULT now()                                   NOT NULL
);


ALTER TABLE public.users
    OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, true);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);