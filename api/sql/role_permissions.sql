-- Role Permissions table
CREATE TABLE IF NOT EXISTS role_permissions (
  id bigserial primary key,
  role_id bigint not null references roles(id) on delete cascade,
  permission_id bigint not null references permissions(id) on delete cascade,
  active boolean not null default true,
  status varchar(1) not null default 'O',
  flag varchar(45),
  uuid uuid not null default gen_random_uuid() unique,
  created_at timestamptz not null default now(),
  created_by bigint default 0,
  updated_at timestamptz not null default now(),
  updated_by bigint default 0,
  deleted_at timestamptz,
  deleted_by bigint default 0,
  CONSTRAINT unique_role_permission UNIQUE (role_id, permission_id, deleted_at)
);

CREATE UNIQUE INDEX unique_role_permission_active ON role_permissions(role_id, permission_id)
WHERE deleted_at IS NULL;