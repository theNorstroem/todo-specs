name: ListDemossRequest
type: ListDemossRequest
description: request message for ListDemoss
__proto:
    package: demo
    targetfile: reqmsgs.proto
    imports:
        - google/protobuf/empty.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/demo;demopb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.furo.basedemo
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
            label: demo.ListDemossRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    q:
        type: string
        description: Use this to search for a person by text.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.ListDemossRequest.q.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    filter:
        type: string
        description: Use this field to filter the persons, this is not searching.
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.ListDemossRequest.filter.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    order_by:
        type: string
        description: Use this field to specify the ordering.
        __proto:
            number: 4
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.ListDemossRequest.order_by.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    page:
        type: string
        description: Use this field to specify page to display.
        __proto:
            number: 5
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.ListDemossRequest.page.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
