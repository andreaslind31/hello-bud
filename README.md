Trying out

# Bud - The Full-Stack Web Framework for Go

## how to
### - create stuff
project:
```bash
bud create hello-bud
```

controllers with views:
* Scaffold a new post controller with only the index and show actions
```bash
bud new controller post index show
```

*Scaffold a new post controller at the root route
```bash
bud new controller post:/
```

### - run
```bash
bud run
```

For more info: [bud documentation](https://github.com/livebud/bud)
