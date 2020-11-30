name: Demos
version: ""
description: |
    Give some good description here please
lifecycle: null
__proto:
    package: demo
    targetfile: demoservice.proto
    imports:
        - google/api/annotations.proto
        - demo/reqmsgs.proto
        - google/protobuf/empty.proto
        - furo/signatures/signatures.proto
        - demo/demo.proto
    options:
        go_package: github.com/theNorstroem/todo-specs/dist/pb/demo;demopb
        java_multiple_files: "true"
        java_outer_classname: DemoserviceProto
        java_package: com.furo.basedemo
services:
    List:
        description: List persons with pagination.
        data:
            request: google.protobuf.Empty
            response: demo.DemoCollection
            bodyfield: body
        deeplink:
            description: 'List: GET /demos google.protobuf.Empty , demo.DemoCollection #List persons with pagination.'
            href: /demos
            method: GET
            rel: list
        query:
            q:
                constraints: {}
                description: Use this to search for a person by text.
                meta: null
                type: string
            filter:
                constraints: {}
                description: Use this field to filter the persons, this is not searching.
                meta: null
                type: string
            order_by:
                constraints: {}
                description: Use this field to specify the ordering.
                meta: null
                type: string
            page:
                constraints: {}
                description: Use this field to specify page to display.
                meta: null
                type: string
        rpc_name: ListDemoss
    Get:
        description: Returns a single person.
        data:
            request: google.protobuf.Empty
            response: demo.DemoEntity
            bodyfield: body
        deeplink:
            description: 'Get: GET /demos/{dem} google.protobuf.Empty , demo.DemoEntity #Returns a single person.'
            href: /demos/{dem}
            method: GET
            rel: self
        query:
            dem:
                constraints: {}
                description: The query param prs stands for the id of a person.
                meta: null
                type: string
        rpc_name: GetDemos
    Create:
        description: Use this to create new persons.
        data:
            request: demo.Demo
            response: furo.signatures.EmptyEntity
            bodyfield: body
        deeplink:
            description: 'Create: POST /demos demo.Demo , furo.signatures.EmptyEntity #Use this to create new persons.'
            href: /demos
            method: POST
            rel: create
        query: {}
        rpc_name: CreateDemos
