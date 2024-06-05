CREATE INDEX movies_title_idx ON movies(title text_pattern_ops);

CREATE INDEX user_id_idx on ratings(user_id);