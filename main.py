import psycopg2
import os


def main():
    try:
        with psycopg2.connect(user=os.environ["DB_USERNAME"], 
            password=os.environ["DB_PASSWORD"], 
            host=os.environ["DB_HOST"], 
            port=os.environ["DB_PORT"], 
            database=os.environ["DB_DATABASE"]) as connection:
            with connection.cursor() as cursor:
                print("Refreshing materialized view " + os.environ["INPUT_VIEW"])
                cursor.execute(f'refresh materialized view {os.environ["INPUT_VIEW"]};')
                connection.commit()

    except (Exception, psycopg2.Error) as error:
        print("Failed to refresh materialized view: ", error)

if __name__ == "__main__":
    main()