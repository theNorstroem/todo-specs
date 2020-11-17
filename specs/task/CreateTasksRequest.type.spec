name: CreateTasksRequest
type: CreateTasksRequest
description: request message for CreateTasks
__proto:
    package: task
    targetfile: reqmsgs.proto
    imports:
        - task/task.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/task;taskpb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: pro.furo.todo
fields:
    body:
        type: .task.Task
        description: Body with task.Task
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: task.CreateTasksRequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
