name: CreateDemosRequest
type: CreateDemosRequest
description: request message for CreateDemos
__proto:
    package: demo
    targetfile: reqmsgs.proto
    imports:
        - demo/demo.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/demo;demopb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.furo.basedemo
fields:
    body:
        type: .demo.Demo
        description: Body with demo.Demo
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.CreateDemosRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
