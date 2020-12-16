name: Filter
type: Filter
description: Filter root object
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
    clause:
        type: filter.Condition
        description: Root bracket with ors inside, this is the most complex but most flexible way to define a filter
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: filter.Filter.clause.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    flat:
        type: map<string,filter.Condition>
        description: |-
            Shortcut to set filter conditions without nesting.
            It is up to you how the server handles the request.
            Examples for a flat filter a,b,c:
            - all active conditions *must* match (a && b && c).
            - all conditions are handled as or (a || b || c).
            - you build your own logic like (a && b) || c.
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: filter.Filter.must.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
