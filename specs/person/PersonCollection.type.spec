name: PersonCollection
type: PersonCollection
description: Collectioncontainer which holds a person.Person
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
    entities:
        type: person.PersonEntity
        description: the data contains a person.Person
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.PersonCollection.entities.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
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
            label: person.PersonCollection.links.label
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
            label: person.PersonCollection.meta.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}