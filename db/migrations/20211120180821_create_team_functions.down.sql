DROP FUNCTION team_get_by_id(IN teamId UUID);
DROP FUNCTION team_get_user_role(IN userId UUID, IN teamId UUID);
DROP FUNCTION team_list(IN l_limit INTEGER, IN l_offset INTEGER);
DROP FUNCTION team_list_by_user(IN userId UUID, IN l_limit INTEGER, IN l_offset INTEGER);
DROP FUNCTION team_create(IN userId UUID, IN teamName VARCHAR(256));
DROP FUNCTION team_user_list(IN teamId UUID, IN l_limit INTEGER, IN l_offset INTEGER);
DROP FUNCTION team_user_add(IN teamId UUID, IN userId UUID, IN userRole VARCHAR(16));
DROP PROCEDURE team_user_remove(teamId UUID, userId UUID);
DROP FUNCTION team_battle_list(IN teamId UUID, IN l_limit INTEGER, IN l_offset INTEGER);
DROP FUNCTION team_battle_add(IN teamId UUID, IN battleId UUID);
DROP FUNCTION team_battle_remove(IN teamId UUID, IN battleId UUID);
DROP PROCEDURE team_delete(teamId UUID);