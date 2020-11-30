name: DemoEntity
type: DemoEntity
description: Entitycontainer which holds a demo.Demo
__proto:
    package: demo
    targetfile: demo.proto
    imports:
        - furo/furo.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/demo;demopb
        java_multiple_files: "true"
        java_outer_classname: DemoProto
        java_package: com.furo.basedemo
fields:
    data:
        type: demo.Demo
        description: the data contains a demo.Demo
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.DemoEntity.data.label
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
            label: demo.DemoEntity.links.label
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
            label: demo.DemoEntity.meta.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
