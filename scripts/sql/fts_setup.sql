-- Adding document field
ALTER TABLE movies
    ADD COLUMN IF NOT EXISTS document tsvector;

-- Compute tsvector for each movie
UPDATE movies
SET document = setweight(to_tsvector(title), 'A') ||
               setweight(to_tsvector(overview), 'C') ||
               setweight(to_tsvector(coalesce(tagline, '')), 'D');

-- Creating index for tsvector document
CREATE INDEX IF NOT EXISTS movies_document_idx
    ON movies
        USING gin (document);

-- Trigger for recomputing row tsvector after each insert or update
CREATE FUNCTION movies_tsvector_trigger() RETURNS TRIGGER AS
$$
BEGIN
    new.document = setweight(to_tsvector(title), 'A') ||
                   setweight(to_tsvector(overview), 'C') ||
                   setweight(to_tsvector(coalesce(tagline, '')), 'D');
END

$$ LANGUAGE 'plpgsql';

CREATE TRIGGER movies_document_trigger
    BEFORE INSERT OR UPDATE
    ON movies
    FOR EACH ROW
EXECUTE PROCEDURE movies_tsvector_trigger();
