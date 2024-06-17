# DouTok 文档站

This website is built using [Docusaurus](https://docusaurus.io/), a modern static website generator.

## 本地启动

```shell
yarn install
```

```shell
yarn start
```

## 如何部署

提交一个PR，在PR被合并后将被自动部署到Github Pages上。

## 如何增加文档

1. `./docs`目录用于存放DouTok相关的文档，在其中添加目录和Markdown文件将会直接显示在文档站中。图片请统一存放在`./static/img`目录下。
2. `./blog`目录用于存放DouTok的博客文章，在新增博客前请确认`author.yml`中有您的信息。添加以日期开头的目录将能显示博客的写作日期。
3. `./src/pages/community.md`对应文档站中的社区页面，您可以在其中添加社区相关的内容。

在编写文档按照Markdown文件格式编写即可。
