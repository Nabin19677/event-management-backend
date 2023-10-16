-- Event Expense Category
CREATE TABLE event_expense_categories (
    category_id serial PRIMARY KEY,
    category_name VARCHAR(50) NOT NULL
);

INSERT INTO event_expense_categories (category_name) VALUES
    ('venue'),
    ('catering'),
    ('decorations');

ALTER TABLE expenses
ADD COLUMN category_id INT REFERENCES event_expense_categories(category_id);

ALTER TABLE expenses
DROP COLUMN category;


