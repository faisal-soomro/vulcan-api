UPDATE users SET admin = false where admin IS NULL;
ALTER TABLE users ALTER COLUMN admin SET NOT NULL;
ALTER TABLE users ALTER COLUMN admin SET DEFAULT FALSE;
UPDATE users SET active = true where active IS NULL;
ALTER TABLE users ALTER COLUMN active SET NOT NULL ;
ALTER TABLE users ALTER COLUMN active SET DEFAULT TRUE;
