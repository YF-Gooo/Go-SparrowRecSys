# Go-SparrowRecSys

Go-SparrowRecSys 参考了王喆老师的 https://github.com/wzhe06/SparrowRecSys 电影推荐系统项目。使用 golang+python 作为主要开发语言, 同时包含了 tensorFlow，pyflink,gozero 等技术栈。同时前端采用 ts+elementplus 开发

## 项目计划

## 环境要求

- golang1.17+
- Python 3.7+
- TensorFlow 2.0+

## 项目数据

项目数据来源于开源电影数据集[MovieLens](https://grouplens.org/datasets/movielens/)，项目自带数据集对 MovieLens 数据集进行了精简，仅保留 1000 部电影和相关评论、用户数据。全量数据集请到 MovieLens 官方网站进行下载，推荐使用 MovieLens 20M Dataset。


## 生成项目

    生成api文件: goctl api -o rec.api
    生成业务代码: goctl api go -api rec.api -dir . -style gozero