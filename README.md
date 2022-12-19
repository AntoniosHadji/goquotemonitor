Go Quote Monitor
=================

Monitoring spread width in bps for various tokens and trade sizes.

Consists of:
* Custom Go code requests quotes and inserts data to DB
* Postgres DB running in docker container
* Grafana front end for charting data (also running in docker)

Dockerfile for containerizing executable so that all pieces can run via Docker

TODO:
[ ] Current quote sources are hard coded into the executable. Move to DB.


