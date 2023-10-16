ALTER TABLE expenses
ADD COLUMN category VARCHAR(50);

ALTER TABLE expenses
DROP COLUMN category_id;

DROP TABLE event_expense_categories;