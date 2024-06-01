# Whatsday

A simple dockerized application to get rid of the task of congratulating your friends and family for birthday.
Just setup your YAML config with the docker container once and never ever think about writing messages anymore!

## Installation

Just create a config file like this:

```yaml
messages:
  - Happy birthday!
birthday: 
  03-28:
    - +0111111111
```
And then just use the following `docker-compose.yml` to init your container.
After starting, you will have to scan the QR code in the container output.