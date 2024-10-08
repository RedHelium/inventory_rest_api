PGDMP      /                |            task    16.2    16.0     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    24355    task    DATABASE     x   CREATE DATABASE task WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE task;
                postgres    false            M           1247    24364    inv_operation_status    TYPE     k   CREATE TYPE public.inv_operation_status AS ENUM (
    'accept',
    'cancel',
    'reject',
    'await'
);
 '   DROP TYPE public.inv_operation_status;
       public          postgres    false            �            1259    24409 	   inventory    TABLE     �   CREATE TABLE public.inventory (
    name character varying(50) NOT NULL,
    id integer NOT NULL,
    id_executor integer NOT NULL
);
    DROP TABLE public.inventory;
       public         heap    postgres    false            �            1259    24415    inventory_id_seq    SEQUENCE     �   CREATE SEQUENCE public.inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.inventory_id_seq;
       public          postgres    false    219            �           0    0    inventory_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.inventory_id_seq OWNED BY public.inventory.id;
          public          postgres    false    220            �            1259    24374    inventory_operations    TABLE     V  CREATE TABLE public.inventory_operations (
    id integer NOT NULL,
    src_executor integer NOT NULL,
    dst_executor integer NOT NULL,
    request_time timestamp with time zone DEFAULT now() NOT NULL,
    status_time timestamp with time zone,
    status public.inv_operation_status DEFAULT 'await'::public.inv_operation_status NOT NULL
);
 (   DROP TABLE public.inventory_operations;
       public         heap    postgres    false    845    845            �            1259    24390    inventory_operations_detail    TABLE     �   CREATE TABLE public.inventory_operations_detail (
    id integer NOT NULL,
    id_inventory integer NOT NULL,
    id_inventory_operation integer NOT NULL
);
 /   DROP TABLE public.inventory_operations_detail;
       public         heap    postgres    false            �            1259    24389 "   inventory_operations_detail_id_seq    SEQUENCE     �   CREATE SEQUENCE public.inventory_operations_detail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 9   DROP SEQUENCE public.inventory_operations_detail_id_seq;
       public          postgres    false    218            �           0    0 "   inventory_operations_detail_id_seq    SEQUENCE OWNED BY     i   ALTER SEQUENCE public.inventory_operations_detail_id_seq OWNED BY public.inventory_operations_detail.id;
          public          postgres    false    217            �            1259    24373    inventory_operations_id_seq    SEQUENCE     �   CREATE SEQUENCE public.inventory_operations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 2   DROP SEQUENCE public.inventory_operations_id_seq;
       public          postgres    false    216            �           0    0    inventory_operations_id_seq    SEQUENCE OWNED BY     [   ALTER SEQUENCE public.inventory_operations_id_seq OWNED BY public.inventory_operations.id;
          public          postgres    false    215            +           2604    24416    inventory id    DEFAULT     l   ALTER TABLE ONLY public.inventory ALTER COLUMN id SET DEFAULT nextval('public.inventory_id_seq'::regclass);
 ;   ALTER TABLE public.inventory ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    219            '           2604    24377    inventory_operations id    DEFAULT     �   ALTER TABLE ONLY public.inventory_operations ALTER COLUMN id SET DEFAULT nextval('public.inventory_operations_id_seq'::regclass);
 F   ALTER TABLE public.inventory_operations ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            *           2604    24393    inventory_operations_detail id    DEFAULT     �   ALTER TABLE ONLY public.inventory_operations_detail ALTER COLUMN id SET DEFAULT nextval('public.inventory_operations_detail_id_seq'::regclass);
 M   ALTER TABLE public.inventory_operations_detail ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            �          0    24409 	   inventory 
   TABLE DATA           :   COPY public.inventory (name, id, id_executor) FROM stdin;
    public          postgres    false    219   �!       �          0    24374    inventory_operations 
   TABLE DATA           q   COPY public.inventory_operations (id, src_executor, dst_executor, request_time, status_time, status) FROM stdin;
    public          postgres    false    216   T"       �          0    24390    inventory_operations_detail 
   TABLE DATA           _   COPY public.inventory_operations_detail (id, id_inventory, id_inventory_operation) FROM stdin;
    public          postgres    false    218   �"       �           0    0    inventory_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.inventory_id_seq', 20, true);
          public          postgres    false    220            �           0    0 "   inventory_operations_detail_id_seq    SEQUENCE SET     P   SELECT pg_catalog.setval('public.inventory_operations_detail_id_seq', 6, true);
          public          postgres    false    217            �           0    0    inventory_operations_id_seq    SEQUENCE SET     J   SELECT pg_catalog.setval('public.inventory_operations_id_seq', 12, true);
          public          postgres    false    215            /           2606    24395 <   inventory_operations_detail inventory_operations_detail_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY public.inventory_operations_detail
    ADD CONSTRAINT inventory_operations_detail_pkey PRIMARY KEY (id_inventory_operation);
 f   ALTER TABLE ONLY public.inventory_operations_detail DROP CONSTRAINT inventory_operations_detail_pkey;
       public            postgres    false    218            -           2606    24381 .   inventory_operations inventory_operations_pkey 
   CONSTRAINT     l   ALTER TABLE ONLY public.inventory_operations
    ADD CONSTRAINT inventory_operations_pkey PRIMARY KEY (id);
 X   ALTER TABLE ONLY public.inventory_operations DROP CONSTRAINT inventory_operations_pkey;
       public            postgres    false    216            1           2606    24421    inventory inventory_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.inventory DROP CONSTRAINT inventory_pkey;
       public            postgres    false    219            2           2606    24422 .   inventory_operations_detail fk_id_to_inventory    FK CONSTRAINT     �   ALTER TABLE ONLY public.inventory_operations_detail
    ADD CONSTRAINT fk_id_to_inventory FOREIGN KEY (id_inventory) REFERENCES public.inventory(id) NOT VALID;
 X   ALTER TABLE ONLY public.inventory_operations_detail DROP CONSTRAINT fk_id_to_inventory;
       public          postgres    false    219    218    4657            3           2606    24396 8   inventory_operations_detail fk_id_to_inventory_operation    FK CONSTRAINT     �   ALTER TABLE ONLY public.inventory_operations_detail
    ADD CONSTRAINT fk_id_to_inventory_operation FOREIGN KEY (id_inventory_operation) REFERENCES public.inventory_operations(id);
 b   ALTER TABLE ONLY public.inventory_operations_detail DROP CONSTRAINT fk_id_to_inventory_operation;
       public          postgres    false    218    4653    216            �   o   x��H�ϋW04�4�44�4�r�/JUȴ�54�40��44�4Bs�44*��i5i5�4�*3*3k5�j5C6� b�9�q���&��5Z�"��42�4����� �,5      �   �   x�}�M
1F��)����j�Cx�ٔ҅""2����h��B�x�fp� $v ?0
�Ct:��t���#�f<,h�������"Ѳ��s�3zЭA�k�G�Ր�5����6F����)�r��Nv[�vf������7̭�8����%Hf      �   3   x�ȱ� �x��A�^�5�`Bi�H��d�}Hgz��<���
�     