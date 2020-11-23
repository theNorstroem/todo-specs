# todo-specs

This is the spec project for the [todo-server](https://github.com/theNorstroem/todo-server) and [todo-client](https://github.com/theNorstroem/todo-client).

For simplicity we do not make use of a build pipeline. Just check out this spec to the same folder you checked out the todo-client and todo-server repository. 

We write some artifacts (environment.js) directly to ../todo-client/src/config/... for demo purposes only. 
 
Read more about the [furo specs here](https://fidl.furo.pro/) and about the [spectools](https://spectools.furo.pro/) here.

### Spectools config
This project just have a default flow configured in the .spectools config.

#### Commands
You can find the following commands are in the *scripts* directory of this project.

- *gen_transcoder* : "./scripts/grpcgateway/gateway.sh"
- *exec_protoc_messages* : "./scripts/proto2DomainTypes.sh"
- *exec_protoc_services* : "./scripts/proto2DomainServices.sh"
- *copy_file_templates* : "./scripts/copyFileTemplates/copy.sh" # copy additional files to dist
- *update_client_project* : "scripts/updateClientSpec.sh" # we do not use npm for this simple demo


## Overview


```

    :10000
   +-------------------------+
   |                         |
   |       todo-client       |  <-----------------+
   |       (webserver)       |                    |
   +-------------------------+                    |
                |                         +----------------+
                |                         |   todo-specs   |
    :7001       v                         |   ----------   |
   +-------------------------+            +----------------+
   |                         |                    |    \
   |       grpc gateway      |  <-----------------+     \
   |                         |                    |      \
   +-------------------------+                    |     You are here.
                |                                 |
    :7000       v                                 |
   +-------------------------+                    |
   |                         |                    |
   |       todo-server       |  <-----------------+
   |                         |
   +-------------------------+



```