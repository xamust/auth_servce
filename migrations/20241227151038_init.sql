-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS system_roles (
                       uuid uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                       description varchar(128) NOT NULL,
                       organization_uuid uuid NOT NULL,
                       metadata jsonb,
                       created_at timestamptz DEFAULT now(),
                       updated_at timestamptz DEFAULT now(),
                       deleted_at timestamptz
);
CREATE INDEX IF NOT EXISTS system_roles_id_idx ON system_roles (uuid);
CREATE UNIQUE INDEX IF NOT EXISTS system_roles_organization_description ON system_roles (description, organization_uuid);
INSERT INTO system_roles (uuid,description, organization_uuid, metadata)
VALUES
    ('c901adc3-f360-4db8-992d-bf630c2386fa','admin', '5549f99b-b05a-4604-834f-2147e1687cc4', '{"role_title":"Администратор системы"}'),
    ('e3d62285-fbaa-46f4-8077-83dfb5b17227','organization_owner', '5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"role_title":"Владелец организации"}'),
    ('253beae8-7be4-4d60-a6ca-2d3ed92a2f8b','manager', '5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"role_title":"Менеджер организации"}'),
    ('0516441f-0477-4ffb-a2a9-2477f831e60d','user', '5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"role_title":"Пользователь системы"}');


CREATE TABLE IF NOT EXISTS system_roles_permissions (
                       uuid uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                       description varchar(128) NOT NULL,
                       organization_uuid uuid NOT NULL,
                       metadata jsonb,
                       created_at timestamptz DEFAULT now(),
                       updated_at timestamptz DEFAULT now(),
                       deleted_at timestamptz
);
CREATE INDEX IF NOT EXISTS system_roles_permissions_id_idx ON system_roles_permissions (uuid);
CREATE UNIQUE INDEX IF NOT EXISTS system_roles_permissions_organization_description ON system_roles_permissions (description, organization_uuid);
INSERT INTO system_roles_permissions (uuid,description, organization_uuid, metadata)
VALUES
    ('deaadff5-7983-4ac0-bcc3-00de18da0ae1','manage_all_users','5549f99b-b05a-4604-834f-2147e1687cc4', '{"permission_title":"Управление всеми пользователями"}'),
    ('41b56cd0-6b65-4873-8235-b117bff11f19','manage_all_roles','5549f99b-b05a-4604-834f-2147e1687cc4', '{"permission_title":"Управление всеми ролями"}'),
    ('5a92be63-9463-4c24-a436-04233f89ef90','manage_all_permissions','5549f99b-b05a-4604-834f-2147e1687cc4', '{"permission_title":"Управление всеми правами"}'),
    ('18e971a5-4cd7-4dd5-b7d8-261ee63a8ccf','manage_organization_users','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"permission_title":"Управление пользователями компании"}'),
    ('d33cd63e-4e6d-4eba-87a7-742bb140a6c5','manage_organization_roles','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"permission_title":"Управление ролями компании"}'),
    ('31290212-1813-4b07-ba75-775431e6503a','manage_organization_permissions','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"permission_title":"Управление правами компании"}'),
    ('639cf2e8-87dc-417e-adf7-328d9e2c170f','view_organization_users','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"permission_title":"Просмотр сотрудников компании"}'),
    ('fadf7e07-4b7a-4b4b-80d6-58f006f2eed3','change_own_password','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', '{"permission_title":"Изменение своего пароля"}');

CREATE TABLE IF NOT EXISTS system_roles_permissions_relations (
                       uuid uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                       system_role_uuid uuid REFERENCES system_roles(uuid),
                       permission_uuid uuid REFERENCES system_roles_permissions(uuid),
                       organization_uuid uuid NOT NULL,
                       created_at timestamptz DEFAULT now(),
                       updated_at timestamptz DEFAULT now(),
                       deleted_at timestamptz
);
CREATE INDEX IF NOT EXISTS system_roles_permissions_relations_id_idx ON system_roles_permissions_relations (uuid);
CREATE UNIQUE INDEX IF NOT EXISTS system_roles_permissions_relations_role_description ON system_roles_permissions_relations (system_role_uuid, permission_uuid, organization_uuid);
INSERT INTO system_roles_permissions_relations (system_role_uuid, permission_uuid, organization_uuid)
VALUES
    ('c901adc3-f360-4db8-992d-bf630c2386fa','deaadff5-7983-4ac0-bcc3-00de18da0ae1','5549f99b-b05a-4604-834f-2147e1687cc4'),
    ('c901adc3-f360-4db8-992d-bf630c2386fa','41b56cd0-6b65-4873-8235-b117bff11f19','5549f99b-b05a-4604-834f-2147e1687cc4'),
    ('c901adc3-f360-4db8-992d-bf630c2386fa','5a92be63-9463-4c24-a436-04233f89ef90','5549f99b-b05a-4604-834f-2147e1687cc4'),
    ('e3d62285-fbaa-46f4-8077-83dfb5b17227','18e971a5-4cd7-4dd5-b7d8-261ee63a8ccf','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01'),
    ('e3d62285-fbaa-46f4-8077-83dfb5b17227','d33cd63e-4e6d-4eba-87a7-742bb140a6c5','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01'),
    ('e3d62285-fbaa-46f4-8077-83dfb5b17227','31290212-1813-4b07-ba75-775431e6503a','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01'),
    ('253beae8-7be4-4d60-a6ca-2d3ed92a2f8b','639cf2e8-87dc-417e-adf7-328d9e2c170f','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01'),
    ('253beae8-7be4-4d60-a6ca-2d3ed92a2f8b','fadf7e07-4b7a-4b4b-80d6-58f006f2eed3','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01'),
    ('0516441f-0477-4ffb-a2a9-2477f831e60d','fadf7e07-4b7a-4b4b-80d6-58f006f2eed3','5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01');

CREATE TABLE IF NOT EXISTS users (
    uuid uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    email varchar(128) NOT NULL,
    password_hash varchar(256) NOT NULL,
    system_role_uuid uuid REFERENCES system_roles(uuid),
    is_active boolean DEFAULT false,
    organization_uuid varchar(36) NOT NULL,
    metadata jsonb,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
    );
CREATE INDEX IF NOT EXISTS users_id_idx ON users (uuid);
CREATE UNIQUE INDEX IF NOT EXISTS users_organization_email ON users (email, organization_uuid);
INSERT INTO users (email, password_hash, system_role_uuid, is_active, organization_uuid, metadata)
VALUES
    ('admin@example.com', '$2a$10$FuHpGfpjv2BjKFhDiut45uimN7o5qqDSvOu1uKVTGZhJuidjqE.26', 'c901adc3-f360-4db8-992d-bf630c2386fa', true, '5549f99b-b05a-4604-834f-2147e1687cc4',null),
    ('organization_owner@example.com', '$2a$10$FuHpGfpjv2BjKFhDiut45uimN7o5qqDSvOu1uKVTGZhJuidjqE.26', 'e3d62285-fbaa-46f4-8077-83dfb5b17227', true, '5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', null),
    ('manager@example.com', '$2a$10$FuHpGfpjv2BjKFhDiut45uimN7o5qqDSvOu1uKVTGZhJuidjqE.26', '253beae8-7be4-4d60-a6ca-2d3ed92a2f8b', true, '5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', null),
    ('user@example.com', '$2a$10$FuHpGfpjv2BjKFhDiut45uimN7o5qqDSvOu1uKVTGZhJuidjqE.26', '0516441f-0477-4ffb-a2a9-2477f831e60d', true, '5fad25aa-0e2c-4ceb-8e08-0d4e4f9e4d01', null);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP INDEX IF EXISTS users_id_idx;
DROP INDEX IF EXISTS users_organization_email;

DROP TABLE IF EXISTS system_roles CASCADE;
DROP INDEX IF EXISTS system_roles_id_idx;
DROP INDEX IF EXISTS system_roles_organization_description;

DROP TABLE IF EXISTS system_roles_permissions CASCADE;
DROP INDEX IF EXISTS system_roles_permissions_id_idx;
DROP INDEX IF EXISTS system_roles_permissions_organization_description;

DROP TABLE IF EXISTS system_roles_permissions_relations CASCADE;
DROP INDEX IF EXISTS system_roles_permissions_relations_id_idx;
DROP INDEX IF EXISTS system_roles_permissions_relations_role_description;
-- +goose StatementEnd
