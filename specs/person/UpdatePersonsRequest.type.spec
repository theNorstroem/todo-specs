name: UpdatePersonsRequest
type: UpdatePersonsRequest
description: request message for UpdatePersons
__proto:
    package: person
    targetfile: reqmsgs.proto
    imports:
        - google/protobuf/field_mask.proto
        - person/person.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/person;personpb
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
            label: person.UpdatePersonsRequest.body.label
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
            label: person.UpdatePersonsRequest.pn.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    update_mask:
        type: google.protobuf.FieldMask
        description: Needed to patch a record
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.UpdatePersonsRequest.update_mask.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
