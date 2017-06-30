# ansible-one-inventory

Dynamic Ansible inventory via OpenNebula API.

[![Travis CI Build Status](https://travis-ci.org/marthjod/ansible-one-inventory.svg?branch=master)](https://travis-ci.org/marthjod/ansible-one-inventory)

```bash
go get -u
go build -o one-inv main.go
```

## Configure

See [opennebula-inventory.example.yaml](https://github.com/marthjod/ansible-one-inventory/blob/master/opennebula-inventory.example.yaml).

## Run

```bash
one-inv --list
one-inv --debug --list
```

## Caveats, assumptions, limitations

`one-inv --host=foo` not implemented yet.

Any non-flag CLI args must be passed with an equals sign.

The inventory tool assumes that any role or group membership is sufficiently encoded in (and can thus be deduced from) the VM name (or another, explicitly configured hostname field).
It does not inspect additional attributes for determining group association (although this may change in the future).

## Test

```
go test ./...
```

### Coverage

```
# for each package
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Example

Given VMs with names _db-staging-west_, _web-staging-west_, _db-production-east_ and the following config,

```yaml
static_group_filters:
  web: "^web"
  database: "^db"
  east: "-east$"
  west: "-west$"
dynamic_group_filters:
  pattern: ".+-(staging|production)-(east|west)"
  prefix: ""
  infix: "-"
  suffix: ""
  pattern_replace: "-(we|ea)st"
```

the output will look like this:

```json
{
  "database": [
    "db-production-east",
    "db-staging-east"
  ],
  "east": [
    "db-production-east",
    "db-staging-east"
  ],
  "production": [
    "db-production-east"
  ],
  "staging": [
    "db-staging-east"
  ],
  "web": [
    "web-staging-west"
  ],
  "west": [
    "web-staging-west"
  ]
}
```

## Documentation

```
godoc -http=:6060
```

## Ansible

See

- http://docs.ansible.com/ansible/intro_dynamic_inventory.html
- http://docs.ansible.com/ansible/dev_guide/developing_inventory.html
