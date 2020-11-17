name: DeletePersonsRequest
type: DeletePersonsRequest
description: request message for DeletePersons
__proto:
    package: person
    targetfile: reqmsgs.proto
    imports:
        - google/protobuf/empty.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/person;personpb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.furo.baseperson
fields:
    body:
        type: .google.protobuf.Empty
        description: Body with google.protobuf.Empty
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.DeletePersonsRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    pn:
        type: string
        description: The query param pn stands for the id of a person.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.DeletePersonsRequest.pn.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
