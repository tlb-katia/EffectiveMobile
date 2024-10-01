CREATE TABLE IF NOT EXISTS songs (
   id SERIAL PRIMARY KEY,
   group_name VARCHAR(255),
   song_name VARCHAR(255) UNIQUE,
   release_date DATE,
   link TEXT
);

CREATE TABLE IF NOT EXISTS verses
(
    id           SERIAL PRIMARY KEY,
    song_id      INT REFERENCES songs (id) ON DELETE CASCADE,
    verse_number INT,
    verse_text   TEXT

);










--
-- INSERT INTO songs (group_name, song_name, release_date, link)
-- VALUES
-- ('The Beatles', 'Hey Jude', '1968-08-26', 'https://en.wikipedia.org/wiki/Hey_Jude'),
-- ('Queen', 'Bohemian Rhapsody', '1975-10-31', 'https://en.wikipedia.org/wiki/Bohemian_Rhapsody'),
-- ('Eagles', 'Hotel California', '1976-12-08', 'https://en.wikipedia.org/wiki/Hotel_California'),
-- ('Led Zeppelin', 'Stairway to Heaven', '1971-11-08', 'https://en.wikipedia.org/wiki/Stairway_to_Heaven'),
-- ('The Rolling Stones', 'Paint It Black', '1966-05-07', 'https://en.wikipedia.org/wiki/Paint_It_Black'),
-- ('Loreen', 'Euphoria', '2012-02-26', 'https://en.wikipedia.org/wiki/Euphoria_(Loreen_song)');
--
-- INSERT INTO verses (song_id, verse_number, verse_text)
-- VALUES
--     ((SELECT id FROM songs
--       WHERE group_name = 'The Beatles' AND song_name = 'Hey Jude'),
--
--      1, 'Hey Jude, don''t make it bad.
-- Take a sad song and make it better.
-- Remember to let her into your heart,
-- Then you can start to make it better.'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'The Beatles' AND song_name = 'Hey Jude'),
--
--      2, 'Hey Jude, don''t be afraid.
-- You were made to go out and get her.
-- The minute you let her under your skin,
-- Then you begin to make it better.'),
--
-- ((SELECT id FROM songs
--   WHERE group_name = 'The Beatles' AND song_name = 'Hey Jude'),
--
--     3, 'And anytime you feel the pain, hey Jude, refrain,
-- Don''t carry the world upon your shoulders.
-- For well you know that it''s a fool who plays it cool
-- By making his world a little colder.');
--
--
-- INSERT INTO verses (song_id, verse_number, verse_text)
-- VALUES
--     ((SELECT id FROM songs
--       WHERE group_name = 'Queen' AND song_name = 'Bohemian Rhapsody'),
--
--      1, 'Is this the real life? Is this just fantasy?
-- Caught in a landslide, no escape from reality
-- Open your eyes, look up to the skies and see
-- I''m just a poor boy, I need no sympathy
-- Because I''m easy come, easy go
-- Little high, little low
-- Any way the wind blows doesn''t really matter to me, to me'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'Queen' AND song_name = 'Bohemian Rhapsody'),
--
--      2, 'Mama, just killed a man
-- Put a gun against his head, pulled my trigger, now he''s dead
-- Mama, life had just begun
-- But now I''ve gone and thrown it all away
-- Mama, ooh, didn''t mean to make you cry
-- If I''m not back again this time tomorrow
-- Carry on, carry on as if nothing really matters'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'Queen' AND song_name = 'Bohemian Rhapsody'),
--
--      3, 'Too late, my time has come
-- Sends shivers down my spine, body''s aching all the time
-- Goodbye, everybody, I''ve got to go
-- Gotta leave you all behind and face the truth
-- Mama, ooh (any way the wind blows)
-- I don''t wanna die
-- I sometimes wish I''d never been born at all');
--
--
--
-- INSERT INTO verses (song_id, verse_number, verse_text)
-- VALUES
--     ((SELECT id FROM songs
--       WHERE group_name = 'Eagles' AND song_name = 'Hotel California'),
--      1,
--      'On a dark desert highway, cool wind in my hair
-- Warm smell of colitas rising up through the air
-- Up ahead in the distance, I saw a shimmering light
-- My head grew heavy and my sight grew dim, I had to stop for the night
-- There she stood in the doorway, I heard the mission bell
-- And I was thinkin'' to myself, "This could be heaven or this could be hell"
-- Then she lit up a candle and she showed me the way
-- There were voices down the corridor, I thought I heard them say'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'Eagles' AND song_name = 'Hotel California'),
--      2, 'Welcome to the Hotel California
-- Such a lovely place (such a lovely place)
-- Such a lovely face
-- Plenty of room at the Hotel California
-- Any time of year (any time of year)
-- You can find it here'),
--
-- ((SELECT id FROM songs
--   WHERE group_name = 'Eagles' AND song_name = 'Hotel California'),
--     3, 'Her mind is Tiffany-twisted, she got the Mercedes-Benz, uh
-- She got a lot of pretty, pretty boys that she calls friends
-- How they dance in the courtyard, sweet summer sweat
-- Some dance to remember, some dance to forget
-- So I called up the Captain, "Please bring me my wine"
-- He said, "We haven''t had that spirit here since 1969"
-- And still, those voices are calling from far away
-- Wake you up in the middle of the night just to hear them say');
--
--
--
--
-- INSERT INTO verses (song_id, verse_number, verse_text)
-- VALUES
--     ((SELECT id FROM songs
--        WHERE group_name = 'Led Zeppelin' AND song_name = 'Stairway to Heaven'),
--      1,
--      'There''s a lady who''s sure all that glitters is gold
-- And she''s buying a stairway to Heaven'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'Led Zeppelin' AND song_name = 'Stairway to Heaven'),
--      2,
--      'When she gets there she knows, if the stores are all closed
-- With a word she can get what she came for'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'Led Zeppelin' AND song_name = 'Stairway to Heaven'),
--      3,
--      'Ooh, ooh, and she''s buying a stairway to Heaven');
--
--
--
-- INSERT INTO verses (song_id, verse_number, verse_text)
-- VALUES
--     ((SELECT id FROM songs
--       WHERE group_name = 'The Rolling Stones' AND song_name = 'Paint It Black'),
--      1,
--      'I see a red door
-- And I want it painted black
-- No colors anymore
-- I want them to turn black'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'The Rolling Stones' AND song_name = 'Paint It Black'),
--      2,
--     'I see the girls walk by
--     Dressed in their summer clothes
--     I have to turn my head
--     Until my darkness goes'),
--
--     ((SELECT id FROM songs
--       WHERE group_name = 'The Rolling Stones' AND song_name = 'Paint It Black'),
--      3,
--      'I see a line of cars
-- And they''re all painted black
-- With flowers and my love
-- Both never to come back');
--
--
--
-- INSERT INTO verses (song_id, verse_number, verse_text)
-- VALUES
--     ((SELECT id FROM songs
--     WHERE group_name = 'Loreen' AND song_name = 'Euphoria'),
--
--     1, 'Why, why can''t this moment last forevermore?
--     Tonight, tonight eternity''s an open door
--     No, don''t ever stop doing the things you do
--     Don''t go, in every breath I take, I''m breathing yo'),
--
--     ((SELECT id FROM songs
--     WHERE group_name = 'Loreen' AND song_name = 'Euphoria'),
--
--     2, 'Euphoria
--     Forever, ''til the end of time
--     From now on, only you and I
--     We''re going u-u-u-u-u-u-up
--     Euphoria
--     An everlasting piece of art
--     A beating love within my heart
--     We''re going u-u-u-u-u-u-up'),
--
--     ((SELECT id FROM songs
--     WHERE group_name = 'Loreen' AND song_name = 'Euphoria'),
--
--     3, 'We are here, we''re all alone in our own universe
--     We are free, where everything''s allowed and love comes first
--     Forever and ever together, we sail into infinity
--     We''re higher and higher and higher, we''re reaching for divinity');


