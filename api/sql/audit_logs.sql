-- Audit Logs table
CREATE TABLE IF NOT EXISTS audit_logs (
  id bigserial primary key,
  user_id bigint not null references users(id),
  token text,
  path varchar(250),
  action varchar(150),
  response_status int8 default 0,
  module_id bigint default 0,
  module varchar(150),
  before_data_change jsonb,
  after_data_change jsonb,
  description text,
  ip_address inet,
  user_agent text,
  created_at timestamptz not null default now()
);