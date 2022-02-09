# Write documents for swagger

You've already seen the swagger-ui at http://127.0.0.1:14000/swagger/index.html. 
This feature is provided by [swaggo](https://github.com/swaggo).

You can see an example [here](https://github.com/swaggo/swag/tree/master/example/celler/controller). 

If you are not familiar with swagger, there's a [OpenAPI Specification](https://swagger.io/docs/specification/2-0/basic-structure/).

When you update some document, you can rebuild the project

```
cd dkdk
docker-compose down && docker-compose build && docker-compose up -d
```

All the routers are supposed to add docs.