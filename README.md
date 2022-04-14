#data-Collection-Tree

Steps to run

1. clone the repository in $go Path/src
2. cd ../path/src
3. go run main.go

Run On Postman.

1. select POST Api option
2. select Body
    To Set Data
        a. http://localhost:8081/v1/insert
        b. select Raw -> Json. 
        c. {
            "dim": [{
                "key": "device",
                "val": "desktop"
                },
                {
                    "key": "country",
                    "val": "IN"
                }
            ],
            "metrics": [{
                        "key": "webreq",
                        "val": 50
                    },
                    {
                        "key": "timespent",
                        "val": 40
                    }
                ]
            }
        d. click on Send

    To Get Data
        a. http://localhost:8081/v1/query
        b. select Raw -> Json. 
        c. {
                "dim": [{
                        "key": "country",
                        "val": "IN"
                }]
            }
        d. click on send

** all names are case sensitive

Output for Set Data:
{
    "Res":"200 OK",
    "Output":{
        "dim":null,
        "metrics":null
    }
}


Output for Get Data:
{
    "Res":"200 OK",
    "Output":{
        "dim":[
            {
                    "key":"country",
                    "val":"IN"
            }
            ],
        "metrics":[
            {
                "key":"webreq",
                "val":50
            },
            {
                "key":"timespent",
                "val":40
            }
        ]
    }
}

++ Data may not apear in alignment