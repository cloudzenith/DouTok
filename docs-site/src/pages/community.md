# 社区/如何参与

首先，感谢DouTok的全体贡献者。DouTok因为有了每位贡献者的付出，才能变的越来越好。

![DouTok贡献者](https://contrib.rocks/image?repo=tremblingv5/DouTok)

## 线下交流群

QQ：622383022

## 如何参与DouTok项目设计与开发

### 交流过程中的语言

DouTok欢迎任何语言的交流，建议使用中文/英文中的一种。DouTok目前所有开发者均来自中国，所以请不要介意直接使用中文。

另外，很多其他开源项目建议使用英文，所以在交流过程中使用英文也是一个不错的选择。

总而言之，无论中文还是英文都是可以的。

### 提交Issue或发起Discussion

DouTok欢迎任何Issue和Discussion，请随意提交Issue或发起Discussion，不必感觉拘谨，请千万不要觉得自己的想法可能是错误的，任何Issue和Discussion都将会让DouTok变得更好。当然，DouTok建议在提交Issue或发起Discussion前先搜索有没有类似的Issue或Discussion，这样可以尽量避免一个问题重复讨论和产生没有对齐想法的情况。

如果您想提交一个Issue，我们建议使用给定的Issue模板。当然，您也可以不使用任何模板或使用任何其他模板提交一个Issue，只要您尽可能让其他人理解您的想法即可。

### 分支管理与提交PR

在开发过程中，DouTok使用git-flow来组织所有分支。简而言之，您应该遵照如下流程去管理分支和提交PR

1. Fork本仓库，即DouTok仓库
2. 创建一个新的分支，并为其起一个恰当的名字
3. 在新建的分支中编写代码、新增文档或做任何事
4. 提交并推送您的代码到Fork仓库中
5. 使用`git rebase`解决与DouTok主分支的冲突
6. 提交PR

当然，DouTok在提交PR方面也有一些建议：

1. 使用给定的PR模板
2. 使用`rebase`来解决冲突
3. 让一个PR尽可能少的包含commit，如果可以的话，只包含一个commit
   > 这样做的好处在于：
   > 1. 使PR提交历史简洁明了
   > 2. 如果产生冲突，过多的commit将会使冲突解决变得复杂
4. 在提交PR前使用linter来检查代码，在提交代码后请确保所有的Github Check通过
