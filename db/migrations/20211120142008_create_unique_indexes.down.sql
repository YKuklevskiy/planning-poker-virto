DROP INDEX email_unique_idx;
ALTER TABLE api_keys DROP CONSTRAINT api_keys_warrior_id_name_key;
ALTER TABLE organization_team DROP CONSTRAINT organization_team_team_id_key;
ALTER TABLE organization_department DROP CONSTRAINT organization_department_organization_id_name_key;
ALTER TABLE department_team DROP CONSTRAINT department_team_team_id_key;
