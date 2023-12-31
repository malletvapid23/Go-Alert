-- +migrate Up
ALTER TABLE user_favorites
    ADD COLUMN tgt_user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    ADD CONSTRAINT user_favorites_user_id_tgt_user_id_key UNIQUE(user_id, tgt_user_id);

-- +migrate Down
ALTER TABLE user_favorites
    DROP CONSTRAINT user_favorites_user_id_tgt_user_id_key,
    DROP COLUMN tgt_user_id;
