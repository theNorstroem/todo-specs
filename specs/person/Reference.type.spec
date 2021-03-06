name: Reference
type: Reference
description: Use this type to set a reference to a person
__proto:
    package: person
    targetfile: person.proto
    imports:
        - furo/furo.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/person;personpb
        java_multiple_files: "true"
        java_outer_classname: PersonProto
        java_package: com.furo.baseperson
fields:
    id:
        type: string
        description: Id of the referenced person.
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.Reference.id.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    display_name:
        type: string
        description: Label of the referenced person
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: person.Reference.display_name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    link:
        type: furo.Link
        description: HTS for the initial search (default works on root resources only)
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: |
                {
                    "rel": "list",
                    "href": "/persons",
                    "method": "GET",
                    "type": "person.Person",
                    "service": "Persons"
                }
            hint: ""
            label: person.Reference.link.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
