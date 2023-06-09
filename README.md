# protobuf-plugin-template
基于golang构建 protobuf plugin 的基本脚手架

## 模块划分

- 分析器
- 内容生成器

## 功能描述

`generator`调用`分析器`返回`内容生成器`构造文件内容，并交由protoc创建`文件`  