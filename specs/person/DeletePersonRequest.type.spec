name: DeletePersonRequest
type: DeletePersonRequest
description: request message for DeletePerson
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
            label: person.DeletePersonRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    prs:
        type: string
        description: The query param prs stands for the id of a person.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.DeletePersonRequest.pn.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
