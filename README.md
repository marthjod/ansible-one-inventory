# ansible-one-inventory

Dynamic Ansible inventory via OpenNebula API.

## Build

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

## Ansible

See

- http://docs.ansible.com/ansible/intro_dynamic_inventory.html
- http://docs.ansible.com/ansible/dev_guide/developing_inventory.html