name: CreatePersonsRequest
type: CreatePersonsRequest
description: request message for CreatePersons
__proto:
    package: person
    targetfile: reqmsgs.proto
    imports:
        - person/person.proto
    options:
        go_package: github.com/theNorstroesm/todo-specs/dist/pb/person;personpb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.furo.baseperson
fields:
    body:
        type: .person.Person
        description: Body with person.Person
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.CreatePersonsRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
