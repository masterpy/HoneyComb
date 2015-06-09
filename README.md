#Honeycomb

##动量影视项目管理及流程开发项目
###代号：蜂巢计划

###一、 开发内容

####1 团队管理（人员管理）
简单的人员管理系统

    添加团队成员
    修改团队成员信息
    删除团队成员（授权取消，信息保留）
    
####2 项目管理
最终以开发文档为准，但至少包含以下几部分

    项目成员管理
    添加项目成员
    分配角色
    移除项目成员
    任务管理（资产管理）
    任务分解
    制定计划
    任务分配
    跟踪反馈
    API（for pipeline）
    
####3 Pipeline
主要包含以下三类工具

    文件检查工具
    模型层级检查工具
    绑定后材质检查工具
    缓存后动画检查工具
    资产发布工具
    模型发布工具
    uv发布工具
    材质发布工具
    动画发布与自动缓存工具
    资产导入工具
    项目资产列表（for layout & animation）
    镜头资产列表（for lighting & rendering）
    
###二、 技术路线

####1 项目管理
服务端采用go语言进行开发，最终部署在linux系统下（CentOS）

数据库采用MySQL

客户端采用Html5 + CSS + Javascript，部分工具使用Python + PySide，开发初期可能会采用QML或Python + PySide来创建原型。

####2 Pipeline
统一采用Python进行开发，UI部分采用PySide

###三、 开发周期
整个项目周期为3个月（到8月30日），根据开发计划会阶段性的放出相应的工具。在开发的过程中，开发的内容可能有所调整。

项目开发周期表：

|开发内容            |截止时间|备注                                      |
|--------------------|--------|------------------------------------------|
|资产拷贝工具（临时）|6月5日  |由于权限问题，需要将发布资产放到Z:/Project|
|资产检查工具|6月30日|  |
|任务分解|7月10日|任务等同于资产|
|资产发布工具|7月10日|半自动|
|团队管理及任务分配|7月30日| |
|资产发布工具|7月30日|最终版|
|跟踪反馈与版本管理|8月30日|一期最终版|


