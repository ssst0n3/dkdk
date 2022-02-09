## backend
* [start](./doc/backend/start.md)
* [write documents for swagger](./doc/backend/swagger.md)

### add different registry driver support
There already are many projects can be used as registry v2 client:
* [genuinetools/reg](https://github.com/genuinetools/reg)
* [deislabs/oras]("https://github.com/deislabs/oras")
* ...

We can use different projects for different features. 
For example, genuinetools/reg supports many apis of registry v2, but it's too low level. 
deislabs/oras has many high level modules we can use directly, but it can not be used to download a specific layer of a repository.

And we will also implement pure http client for registry v2.