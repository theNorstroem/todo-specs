name: Condition
type: Condition
description: Filter condition
__proto:
    package: filter
    targetfile: filter.proto
    imports: []
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/filter;filterpb
        java_multiple_files: "true"
        java_outer_classname: FilterProto
        java_package: com.furo.basefilter
fields:
    fld:
        type: string
        description: Field
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: fld is required
    is:
        type: string
        description: The comparator like gt, eq,...
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    val:
        type: string
        description: The value as string, parse this for your field
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    aoc:
        type: filter.Condition
        description: And bracket with ors inside
        __proto:
            number: 4
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: filter.Condition.aoc.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
