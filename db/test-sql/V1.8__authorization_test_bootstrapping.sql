/* Token is stored in _resources/config/testuser.token */
UPDATE users SET api_token='2d6745c21b39aff06d2a662639ffbc707a6c011a34362b58371def48ad517c28' WHERE email='testuser@vulcan.com';

/* Token is stored in _resources/config/testuser1.token */
INSERT INTO users (firstname, lastname, email, disabled, admin, api_token) VALUES ('User1', 'AuthTest', 'testuser1@vulcan.com', false, false, '678daf8f4462f5515e03ea3b00ba256589c674e3751a17211349761c37794b10');

/* Token is stored in _resources/config/testuser2.token */
INSERT INTO users (firstname, lastname, email, disabled, admin, api_token) VALUES ('User2', 'AuthTest', 'testuser2@vulcan.com', false, false, 'cf510b54466be2fba4d5684b8a6f00b44af135a0aecb2bcc09a39112768487c2');

INSERT INTO teams (name, description) VALUES ('Team1', 'Testing Team for User1');
INSERT INTO teams (name, description) VALUES ('Team2', 'Testing Team for User2');

/* Add User1 in Team1 as owner */
INSERT INTO user_team (user_id, team_id, role, is_default) SELECT u.id as user_id, t.id as team_id, 'owner' as role, false as is_default FROM users u, teams t WHERE u.email = 'testuser1@vulcan.com' AND t.name = 'Team1';
/* Add User2 in Team2 as owner */
INSERT INTO user_team (user_id, team_id, role, is_default) SELECT u.id as user_id, t.id as team_id, 'owner' as role, false as is_default FROM users u, teams t WHERE u.email = 'testuser2@vulcan.com' AND t.name = 'Team2';

INSERT INTO groups (team_id, name) SELECT t.id as team_id, 'Group1' as name FROM teams t WHERE t.name = 'Team1';
INSERT INTO groups (team_id, name) SELECT t.id as team_id, 'Group2' as name FROM teams t WHERE t.name = 'Team2';

INSERT INTO assets (team_id, asset_type_id, identifier) SELECT t.id as team_id, a.id as asset_type_id, 'Asset1' as identifier FROM teams t, asset_types a WHERE t.name = 'Team1' and a.name = 'IP';
INSERT INTO assets (team_id, asset_type_id, identifier) SELECT t.id as team_id, a.id as asset_type_id, 'Asset2' as identifier FROM teams t, asset_types a WHERE t.name = 'Team2' and a.name = 'IP';

INSERT INTO recipients (team_id, email) SELECT t.id as team_id, 'testuser2@vulcan.com' as email FROM teams t WHERE t.name = 'Team2';

INSERT INTO asset_group (asset_id, group_id) SELECT a.id as asset_id, g.id as group_id FROM assets a, groups g WHERE a.identifier = 'Asset2' and g.name='Group2';
