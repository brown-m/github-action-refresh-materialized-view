# GitHub Action: Refresh materialized view

A GitHub action to refresh a materialized view in a Postgres or compatible database (e.g. redshift or panoply).

## Configuration

The action requires the following environment variables to be present to form a connection to the database and refresh the specified view:

```
DB_HOST         -- The host to connect e.g. db.panoply.io
DB_PORT         -- The port to connect on e.g. 5432
DB_USERNAME     -- Database username to connect as e.g. admin_user
DB_PASSWORD     -- Password for the specified database user e.g. password123
DB_DATABASE     -- The database to connect to e.g. demo
```

It is best to configure the connection parameters as secrets in your repo.

## Usage

Combined with GitHub workflows it provides a lightweight scheduler for refreshing materialized views and can accommodate dependencies on other materialized views being refreshed by virtue of the the workflow syntax in GitHub actions. 

To require a view to be refreshed before another can be refreshed, specify the job's `depends_on` parameter. Otherwise views will be refreshed concurrently provided they are defined as there own individual job.

## Sample Workflow

In the below example, the `customer-dimension` and `product-dimension` views are refreshed concurrently. When they have both refreshed the `sales-dataset` view will then refresh as it depends on the former two views for its data. You can specify an arbitrary number of jobs to be completed before this job is run.

```
name: cloudrun-deploy-development
on:
  schedule:
    # Run every 30 minutes
    - cron:  '*/30 * * * *'
env:
  DB_HOST: ${{ secrets.DB_HOST }}
  DB_USERNAME: ${{ secrets.DB_USERNAME }}
  DB_PASSWORD: ${{ secrets.DB_PASSWORD }}
  DB_DATABASE: ${{ secrets.DB_DATABASE }}
  DB_PORT: ${{ secrets.DB_PORT }}

jobs:
  customer-dimension:
    runs-on: ubuntu-latest
    steps:
    - name: Refresh materialized view
      uses: brown-m/github-action-refresh-materialized-view@master
      with:
        view: "public"."customer_dimension"

  product-dimension:
    runs-on: ubuntu-latest
    steps:
    - name: Refresh materialized view
      uses: brown-m/github-action-refresh-materialized-view@master
      with:
        view: "public"."product_dimension"

  sales-dataset:
    needs: [customer-dimension, product-dimension]
    runs-on: ubuntu-latest
    steps:
    - name: Refresh materialized view
      uses: brown-m/github-action-refresh-materialized-view@master
      with:
        view: "public"."sales_dataset"
```

## Notes

The user specified should be the owner of the materialized view, otherwise the user will likely lack the required privileges to perform the `refresh materialized view` SQL query.
