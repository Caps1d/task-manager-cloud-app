CREATE OR REPLACE FUNCTION insert_into_user_teams()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO user_teams (user_id, team_id, user_role)
    VALUES (NEW.manager_id, NEW.id, 'Manager');
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;
