
##通信机制

前后端通过http协议进行通信，客户端发送post请求，服务端返回json数据

客户端向服务器端的请求中必须包含下面两个属性

api-key : ""                 用于验证用户信息
command : "command name"     api命令，根据command将分发到不同的处理函数

服务器端向客户端返回的格式如下

{
    "error": {
        "errorCode" : 0,
        "errorMessage": ""
    },
    "command": "command name",
    "UserCode": string,      //发起该操作的user
    "result": {}             //不同的接口不同的结果
}

客户端首先解析错误信息，如果没有错误则根据将result的内容转发到对应的函数


添加项目
{
    "command" : "addProject",
    "project_name" : "",
    "project_detial" : ""
}

{
    "error": {
        "errorCode" : 0,
        "errorMessage": ""
    },
    "command": "addProject",
    "UserCode": string,      //发起该操作的user
    "result": {}             //不同的接口不同的结果
}

获取所有项目
{
    "command" : "getProjects",
}

{
    "error": {
        "errorCode" : 0,
        "errorMessage": ""
    },
    "command": "addProject",
    "UserCode": string,      //发起该操作的user
    "result": {}             //不同的接口不同的结果
}

添加mission
{
    "command" : "addMission",
    "mission_name" : "",
    "mission_type" : "",
    "mission_detial" : "",
    "project_code" : "",
    "has_child" : 0,
    "parent_code" : "",
    "child_index" : 0,
    "status" : 0
}

{
    "error": {
        "errorCode" : 0,
        "errorMessage": ""
    },
    "command": "addMission",
    "UserCode": string,      //发起该操作的user
    "result": {}             //不同的接口不同的结果
}

获取所有的Mission
{
    "command" : "getMissions",
    "project_code" : "",
}

{
    "error": {
        "errorCode" : 0,
        "errorMessage": ""
    },
    "command": "addProject",
    "UserCode": string,      //发起该操作的user
    "result": {}             //不同的接口不同的结果
}


修改mission

获取所有mission

删除mission

复制mission

改变mission层级



