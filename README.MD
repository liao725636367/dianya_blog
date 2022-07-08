### 实现的接口

![image-20220708204720517](README.assets/image-20220708204720517.png)

![image-20220708204731915](README.assets/image-20220708204731915.png)

### 运行方式

1. 安装 go1.17+mysql+redis 环境 其中mysql 版本要高于 5.7.22
2. 创建 blog 数据库 编码 utf8mb4 导入项目目录下 blog.sql  
3. 修改 manifest/config.yaml 中 数据库和redis 配置
4. 项目目录下运行 go run main.go
