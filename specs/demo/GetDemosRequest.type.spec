name: GetDemosRequest
type: GetDemosRequest
description: request message for GetDemos
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
            label: demo.GetDemosRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    dem:
        type: string
        description: The query param prs stands for the id of a person.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: demo.GetDemosRequest.dem.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
