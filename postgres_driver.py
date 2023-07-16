import psycopg2
import json

def fetch_user_data():
    # Connect to the PostgreSQL database
    conn = psycopg2.connect(
        host="localhost",
        database="postgresql",
        user="nekiarie",
        password="Kiariew@njiiri94"
    )

    # Create a cursor object to interact with the database
    cursor = conn.cursor()

    # Execute the SELECT query
    cursor.execute("SELECT * FROM public.user_table")

    # Fetch all rows from the result set
    rows = cursor.fetchall()

    # Convert the rows to a list of dictionaries
    data = []
    for row in rows:
        d = {
            'user_id': row[0],
            'name': row[1],
            'age': row[2],
            'phone': row[3]
        }
        data.append(d)

    # Convert the data to a JSON string
    json_data = json.dumps(data)

    # Close the cursor and connection
    cursor.close()
    conn.close()

    return json_data

if __name__ == '__main__':
    json_data = fetch_user_data()
    print(json_data)
