import psycopg2
import os


def main():
    try:
        connection = psycopg2.connect(user=os.environ["DB_USERNAME"],
                                    password=os.environ["DB_PASSWORD"],
                                    host=os.environ["DB_HOST"],
                                    port=os.environ["DB_PORT"],
                                    database=os.environ["DB_DATABASE"])
        cursor = connection.cursor()

        # Refresh the view
        print("Refreshing materialized view " + os.environ["INPUT_VIEW"])
        cursor.execute(f'refresh materialized view {os.environ["INPUT_VIEW"]};')
        print("DONE")


    except (Exception, psycopg2.Error) as error:
        print("Failed to refresh materialized view: ", error)

    finally:
        # Close connection.
        if connection:
            cursor.close()
            connection.close()

if __name__ == "__main__":
    main()