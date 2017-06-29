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

## Caveats

- `one-inv --host=foo` not implemented yet
- any non-flag CLI args must be passed with an equals sign

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


## Ansible

See

- http://docs.ansible.com/ansible/intro_dynamic_inventory.html
- http://docs.ansible.com/ansible/dev_guide/developing_inventory.html
