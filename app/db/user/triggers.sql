CREATE TRIGGER after_team_insert
AFTER INSERT ON teams
FOR EACH ROW
EXECUTE FUNCTION insert_into_user_teams();

