# Simple Dependency Injection Container

- `ALERT`: not ready for usage on production

`simple di` is a R&D project on usage of reflection in golang, and also a minimal development of a simple dependency injection container
who aimed for learning manner.

simple di contains of following key items.

* Provider
  a constructor function who provide some attributes for us.
  
* Invoker
  the final function who return an optional error, and consumed attributes who are provided by constructor functions.
  
* Container
  an abstraction over management of providers and invoker, inorder to wireup dependency tree.

## TODOs:

- [x] Implement provider option
- [x] Implement invoker option
- [x] Implement container
- [ ] Manager errors in proper way
- [ ] Add logging option inorder to trace container behavior