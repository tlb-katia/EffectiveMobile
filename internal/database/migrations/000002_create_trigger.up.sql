CREATE OR REPLACE FUNCTION increment_verse_number()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.verse_number := (SELECT COALESCE(MAX(verse_number), 0) + 1 FROM verses WHERE song_id = NEW.song_id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER set_verse_number
    BEFORE INSERT ON verses
    FOR EACH ROW
EXECUTE FUNCTION increment_verse_number();
