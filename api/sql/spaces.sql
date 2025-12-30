-- Spaces table
CREATE TABLE IF NOT EXISTS spaces (
  id bigserial primary key,
  user_id bigint not null references users(id) on delete cascade,
  name varchar(45),
  description text,
  address jsonb,
  price_per_hour numeric(11,2) default 0,
  price_per_day numeric(11,2) default 0,
  price_per_month numeric(11,2) default 0,
  size numeric(11,2) default 0,
  capacity bigint default 0,
  availability varchar(1) not null default 'A',
  active boolean not null default true,
  status varchar(1) not null default 'O',
  flag varchar(45),
  uuid uuid not null default gen_random_uuid() unique,
  created_at timestamptz not null default now(),
  created_by bigint default 0,
  updated_at timestamptz not null default now(),
  updated_by bigint default 0,
  deleted_at timestamptz,
  deleted_by bigint default 0
);