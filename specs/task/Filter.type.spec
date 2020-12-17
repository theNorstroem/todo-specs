name: Filter
type: Filter
description: Filterable fields for the task type This type should be used in the filter argument for the List method. Use the --filterChanged(*._base64) to connect to the filter of the collection agent
__proto:
    package: task
    targetfile: task.proto
    imports:
        - furo/filter/filter.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/task;taskpb
        java_multiple_files: "true"
        java_outer_classname: TaskProto
        java_package: com.furo.basetask
fields:
    display_name:
        type: furo.filter.Condition
        description: Filter for text in the task label
        __proto:
            number: 1
        __ui: null
        meta:
            default: |-
                {
                  "fld": "display_name",
                  "is": "*"
                }
            hint: ""
            label: task.Filter.display_name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific:
                comparators: '*'
                default_comparator: '*'
                hide_comparator: true
        constraints: {}
    note:
        type: furo.filter.Condition
        description: Filter for text in the notes
        __proto:
            number: 2
        __ui: null
        meta:
            default: |-
                {
                  "fld": "note",
                  "is": "*",
                  "val": ""
                }
            hint: ""
            label: task.Filter.note.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific:
                comparators: '*, sw'
                default_comparator: ""
                hide_comparator: false
        constraints: {}
    due_date:
        type: furo.filter.Condition
        description: Filter for the due date
        __proto:
            number: 3
        __ui: null
        meta:
            default: |-
                {
                  "fld": "due_date",
                  "is": "lt",
                  "val": ""
                }
            hint: ""
            label: task.Filter.due_date.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific:
                comparators: lt, gt, lte, gte
                default_comparator: lt
        constraints: {}
    done:
        type: furo.filter.Condition
        description: Filter for completed tasks
        __proto:
            number: 5
        __ui: null
        meta:
            default: |-
                {
                  "fld": "done",
                  "is": "eq"
                  "val": "true"
                }
            hint: ""
            label: task.Filter.done.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific:
                comparators: eq, not
                default_comparator: eq
                hide_comparator: true
        constraints: {}
    person:
        type: furo.filter.Condition
        description: Filter for the person with the id of the person
        __proto:
            number: 6
        __ui: null
        meta:
            default: |-
                {
                  "fld": "responsible_person",
                  "is": "eq",
                  "val": ""
                }
            hint: ""
            label: task.Filter.person.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific:
                comparators: eq
                default_comparator: eq
                hide_comparator: true
        constraints: {}
