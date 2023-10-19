ALTER TABLE event_attendees
ADD CONSTRAINT unique_event_attendees UNIQUE (event_id, user_id);
