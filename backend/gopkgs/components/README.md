# components

Components repo is used to load components for a web app. Such as MySQL, Redis, MQ etc.

In `./components`, there are a file:

- `components.go`: It defined the `Component` struct used to load different components. The `Component` has some methods:
    - `Load`
    - `Start`
    - `GetConfig`

In the child directories, it's the implementation of the components. For example, in `./components/redisx`, it's the implementation of Redis.

In each different implementation, it should have a global map used to store all source of this component. For example, we can use 3 Redis service and named with 'default', 'lock', 'others'. We can get the specified variable/pointer by the source name.

Surely, there must `Init` and `GetXXX` methods for each component to get init or get an instance. The `Init` method will be used to start the web app automatically. So it should receive a fixed parameter.
