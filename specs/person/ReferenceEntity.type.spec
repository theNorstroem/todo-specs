name: ReferenceEntity
type: ReferenceEntity
description: Entitycontainer which holds a person.Reference
__proto:
    package: person
    targetfile: person.proto
    imports:
        - furo/furo.proto
    options:
        go_package: github.com/theNorstroesm/todo-specs/dist/pb/person;personpb
        java_multiple_files: "true"
        java_outer_classname: PersonProto
        java_package: com.furo.baseperson
fields:
    data:
        type: person.Reference
        description: the data contains a person.Reference
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.ReferenceEntity.data.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    links:
        type: furo.Link
        description: the Hateoas links
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.ReferenceEntity.links.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    meta:
        type: furo.Meta
        description: Meta for the response
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.ReferenceEntity.meta.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}