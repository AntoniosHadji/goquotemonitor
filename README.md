Go Quote Monitor
=================

Monitoring spread width in bps for various tokens and trade sizes.

Consists of:
* Custom Go code requests quotes and inserts data to DB
* Postgres DB running in docker container
* Grafana front end for charting data (currently using free cloud grafana)

Dockerfile for containerizing executable so that all pieces can run via Docker


Production quotes require access to a production account.  
using https://cloud.primetrust.com/accounts/0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b



TODO:  

[*] Add http server for UI to manage config data  (Work in Progress)
[ ] Clean up tests to work with data in config table  
[ ] Monitor asset balance in hot and warm  
[ ] Intelligently switch between hot and warm quotes as needed  
[ ] Create mechanism to shutdown goroutines instead of just CTRL-C to exit.  

To run:
-------

Requires ENV vars for:

* DATABASE_URL
* TOKEN

Create database using `./setup/sql/primetrust.sql`
Populate work table with data from `./setup/sql/work.data.sql`
Start executable


Via Docker:  
`docker run -d --env-file ./setup/env ghcr.io/primetrust/goquotemonitor:latest`

File format for env-file is list of VAR/VALUE pairs:

```
DATABASE_URL=postgres://postgres:password@172.17.0.2/primetrust
TOKEN=eyJhbGciOiJIUzI1NiJ9...
```

To exit:
--------

CTRL-C


History:
--------

* 2023-05-10 Added work in progress web UI
* 2023-01-06 added data from work table
* 2022-12-16 Added Coinbase worker
* 2022-12-20 No time delay between initial requests + refactor updates
* 2022-12-21 Data for quote parameters comes from table maintained in DB

Notes:
------

github actions
* replaced `${{ github.repository }}` with `antonioshadji/goquotemonitor` because tag must be lowercase
