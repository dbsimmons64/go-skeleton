-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    txn_date DATE NOT NULL,
    who TEXT NOT NULL,
    description TEXT,
    payee TEXT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    category TEXT,
    inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

    INSERT INTO transactions (txn_date, who, description, payee, amount, category)
    VALUES
    ('2024-10-01', 'John Doe', 'Grocery shopping', 'Supermarket', 45.50, 'Groceries'),
    ('2024-10-05', 'Jane Doe', 'Electricity bill payment', 'Utility Company', 120.00, 'Utilities'),
    ('2024-10-10', 'John Doe', 'Gas station fill-up', 'Gas Station', 60.00, 'Transportation'),
    ('2024-10-15', 'Jane Doe', 'Dinner with family', 'Restaurant', 85.75, 'Dining'),
    ('2024-10-20', 'John Doe', 'Online course purchase', 'Online Education', 150.00, 'Education'),
    ('2024-11-01', 'Jane Doe', 'New shoes', 'Shoe Store', 75.00, 'Shopping'),
    ('2024-11-07', 'John Doe', 'Car maintenance', 'Auto Repair Shop', 300.00, 'Transportation'),
    ('2024-11-12', 'Jane Doe', 'Grocery shopping', 'Supermarket', 50.25, 'Groceries'),
    ('2024-11-18', 'John Doe', 'Gym membership renewal', 'Fitness Center', 99.99, 'Health'),
    ('2024-11-25', 'Jane Doe', 'Holiday gift purchase', 'Gift Store', 125.00, 'Gifts'),
    ('2024-12-03', 'John Doe', 'Internet bill payment', 'Internet Provider', 65.00, 'Utilities'),
    ('2024-12-08', 'Jane Doe', 'New jacket', 'Clothing Store', 120.00, 'Shopping'),
    ('2024-12-15', 'John Doe', 'Groceries for holiday dinner', 'Supermarket', 200.00, 'Groceries'),
    ('2024-12-20', 'Jane Doe', 'Concert tickets', 'Ticket Vendor', 140.00, 'Entertainment'),
    ('2024-12-27', 'John Doe', 'Hotel booking', 'Hotel Chain', 500.00, 'Travel'),
    ('2025-01-05', 'Jane Doe', 'Books for the new semester', 'Bookstore', 60.00, 'Education'),
    ('2025-01-10', 'John Doe', 'Grocery shopping', 'Supermarket', 45.75, 'Groceries'),
    ('2025-01-15', 'Jane Doe', 'Day spa visit', 'Spa', 130.00, 'Health'),
    ('2025-01-20', 'John Doe', 'New headphones', 'Electronics Store', 89.99, 'Shopping'),
    ('2025-01-28', 'Jane Doe', 'Charity donation', 'Charity Organization', 200.00, 'Donations');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
