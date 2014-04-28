NRPE in Go
=======

This project is an alternate implementation of NRPE in golang. There are a few
advantages to this project over the main implemenation of NRPE:

1. Statically linked binary. A single build will work across distros.
2. A configuration format that it not completely custom and is machine parseable.
3. Not maintained by Nagios Enterprises.

### Some notes

* This project is not a drop-in replacement for NRPE. It solely implements the
wire protocol, the per-node configuration is different. This is because the
NRPE configuration format is generally obnoxious to automate.
