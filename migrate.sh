# Run sql file on sqlite3
#!/bin/bash

# Check if sqlite3 is installed
if ! command -v sqlite3 &> /dev/null; then
    echo "sqlite3 could not be found. Please install it first."
    exit 1
fi

# Check if the SQL file exists
if [ ! -f "tables.sql" ]; then
    echo "SQL file 'tables.sql' does not exist."
    exit 1
fi

# Create the database file if it does not exist
if [ ! -f "data.db" ]; then
    echo "Creating database file 'data.db'."
    touch data.db
fi

# Execute the SQL file on the database
echo "Executing SQL file on 'data.db'."
sqlite3 data.db < tables.sql

# Check if the command was successful
if [ $? -eq 0 ]; then
    echo "Database migration completed successfully."
else
    echo "An error occurred during the migration."
    exit 1
fi