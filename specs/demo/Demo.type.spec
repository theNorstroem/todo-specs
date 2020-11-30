name: Demo
type: Demo
description: a demo type
__proto:
    package: demo
    targetfile: demo.proto
    imports: []
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/demo;demopb
        java_multiple_files: "true"
        java_outer_classname: DemoProto
        java_package: com.furo.basedemo
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
            label: demo.Demo.id.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: id is required
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
            label: demo.Demo.display_name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: display_name is required
    something:
        type: string
        description: something
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.Demo.something.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: something is required
