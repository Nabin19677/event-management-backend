-- Users Table
CREATE TABLE users (
    user_id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(16)
);

-- Events Table
CREATE TABLE events (
    event_id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    start_date DATE,
    end_date DATE,
    location VARCHAR(100),
    description TEXT,
    admin_user_id INT REFERENCES users(user_id) -- Admin user for the event
);

-- Event Roles Table
CREATE TABLE event_roles (
    role_id serial PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);

-- Event Organizers Table (Associates users with roles for an event)
CREATE TABLE event_organizers (
    event_organizer_id serial PRIMARY KEY,
    event_id INT REFERENCES events(event_id),
    user_id INT REFERENCES users(user_id),
    role_id INT REFERENCES event_roles(role_id)
);

-- Event Attendees Table (Associates users with events)
CREATE TABLE event_attendees (
    event_attendee_id serial PRIMARY KEY,
    event_id INT REFERENCES events(event_id),
    user_id INT REFERENCES users(user_id)
);

-- Event Sessions/Activities Table
CREATE TABLE event_sessions (
    session_id serial PRIMARY KEY,
    event_id INT REFERENCES events(event_id),
    name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP
);

-- Expenses Table
CREATE TABLE expenses (
    expense_id serial PRIMARY KEY,
    event_id INT REFERENCES events(event_id),
    item_name VARCHAR(100) NOT NULL,
    cost DECIMAL(10, 2) NOT NULL,
    description TEXT,
    category VARCHAR(50)
);

INSERT INTO event_roles (role_name) VALUES
    ('Admin'),
    ('Contributor'),
    ('Attendee');