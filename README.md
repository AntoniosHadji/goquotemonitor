Go Quote Monitor
=================

Monitoring spread width in bps for various tokens and trade sizes.

Consists of:
* Custom Go code requests quotes and inserts data to DB
* Postgres DB running in docker container
* Grafana front end for charting data (also running in docker)

Dockerfile for containerizing executable so that all pieces can run via Docker


Production quotes require access to a production account.
using https://cloud.primetrust.com/accounts/0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b



TODO:  

[ ] Current quote parameters are hard coded into the executable.  
[ ] Create mechanism to shutdown goroutines instead of just CTRL-C to exit.  
[ ] Remove other hard coded configuration  
    - account-id

To run:
-------

Requires ENV vars for:

* DATABASE_URL
* TOKEN

Create database using `setup/primetrust.sql`

Start executable


Via Docker (assuming docker image is named `quotes:ubuntu`) :
`docker run -it --env-file ./setup/env quotes:ubuntu`

File format for env-file is list of VAR/VALUE pairs:

`DATABASE_URL=postgres://postgres:password@172.17.0.2/primetrust
TOKEN=eyJhbGciOiJIUzI1NiJ9...
`

To exit:
--------

CTRL-C
