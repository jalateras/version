## A Skelton Micro Seervice in Go

A simple layout for a micro service written in `golang` 


### Configuration
We are using the [codegansta cli](https://github.com/codegangsta/cli) package for all configuration management.

### Server  
We are using the [echo](https://github.com/labstack/echo) micro web framework for the foundation of our micro services.
It is super fast and provides the ability to plug in middleware and a pretty neat router.
 

### Logging
The [logrus](https://github.com/Sirupsen/logrus) package is used for logging and is API compatible wit hthe standard
library logger.


### TODO

* Add some simple metrics (in-memory)
* Add some health checks
* Add tests
