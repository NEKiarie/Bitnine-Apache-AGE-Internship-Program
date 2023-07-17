import psycopg2
import json

def get_user_table_as_json():
    try:
        # Establish a connection to the PostgreSQL database
        conn = psycopg2.connect(
            host="localhost",
            database="postgresql",
            user="nekiarie",
            password="Kiariew@njiiri94"
        )

        # Create a cursor object to interact with the database
        cursor = conn.cursor()

        # Execute a SELECT query to retrieve data from the table
        cursor.execute("SELECT * FROM public.user_table")

        # Fetch all rows from the resultset
        rows = cursor.fetchall()

        # Convert the rows to a list of dictionaries
        data = []
        for row in rows:
            user_id, name, age, phone = row
            data.append({
                "user_id": int(user_id),
                "name": name,
                "age": int(age),
                "phone": phone
            })

        # Convert the data to JSON format
        json_data = json.dumps(data)

        # Close the cursor and connection
        cursor.close()
        conn.close()

        # Return the JSON string
        return json_data

    except (Exception, psycopg2.DatabaseError) as error:
        print("Error while fetching data from PostgreSQL:", error)

# Call the function to retrieve the user table as a JSON string
user_table_json = get_user_table_as_json()

# Print the JSON string
print(user_table_json)
