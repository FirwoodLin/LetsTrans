# 目录

- [目录](#目录)
- [名词释义](#名词释义)
- [翻译项目管理 LetsTrans-UR-PM](#翻译项目管理-letstrans-ur-pm)
  - [文件分类与导引](#文件分类与导引)
  - [文件增添与删除](#文件增添与删除)
- [辅助翻译界面基本功能 LetsTrans-UR-BF](#辅助翻译界面基本功能-letstrans-ur-bf)
  - [源文本处理](#源文本处理)
  - [目标文本操作](#目标文本操作)
  - [文本检索](#文本检索)
  - [段落管理](#段落管理)
- [翻译记忆 LetsTrans-UR-TM](#翻译记忆-letstrans-ur-tm)
  - [翻译记忆库管理](#翻译记忆库管理)
  - [使用记忆库](#使用记忆库)
- [术语库 LetsTrans-UR-GM](#术语库-letstrans-ur-gm)
  - [添加术语](#添加术语)
  - [显示术语](#显示术语)
  - [使用术语](#使用术语)
- [翻译服务集成 LetsTrans-UR-TI](#翻译服务集成-letstrans-ur-ti)
  - [全局翻译](#全局翻译)
  - [局部翻译](#局部翻译)


# 名词释义

> - 源文本：需要进行翻译的原始文本
> - 目标文本：由用户产生的翻译结果
> - 段落：将文本按照一定的长度进行分割，便于管理

# 翻译项目管理 LetsTrans-UR-PM

## 文件分类与导引

- 用户应能在项目管理中创建子文件夹对文件进行分类
- 用户应能对文件名进行检索

## 文件增添与删除

- 用户应能对文件进行批量删除操作
- 用户应能批量导入合法的外部文件
- 用户应能导入pdf、docx、ppt、TXT、html、md等格式的文件
- 用户应能将翻译后的文件导出成原始格式或.xlf格式(双语对照)

# 辅助翻译界面基本功能 LetsTrans-UR-BF

## 源文本处理

- 用户导入的文本应能自动分段并作为源文本进行展示。
- 用户导入的文本分段后应能在左侧有序列表展示
- 用户应能对现有的分段进行合并与删除
- 用户应能按段修改源文本的内容

## 目标文本操作

- 用户所需求的目标文本应与原文本一一对应分段列表展示
- 用户应能自由修改目标的内容

## 文本检索

- 用户应能在源文本或者目标文本中检索指定的关键词
- 用户检索的结果应以分段列表的形式展示，并在搜索结果中高亮检索关键词

## 段落管理

- 用户应能标记段落的完成状态（已完成/未完成）

# 翻译记忆 LetsTrans-UR-TM

## 翻译记忆库管理

- 用户翻译完成某个段落后，应自动把当前段落加入翻译记忆库
- 用户正在翻译某个段落时，应在翻译记忆库检索相似的记忆单元

## 使用记忆库

- 用户每翻译一个段落的时候，应首先在翻译记忆库中检索与当前段落匹配或相似的记忆单元
- 若搜索到高度相似的文本，应在下方参考栏显示来自翻译记忆库的翻译文本，并显示匹配度，供用户选择

# 术语库 LetsTrans-UR-GM

## 添加术语

- 用户可通过输入自定义的源文本和目标文本向术语库中添加术语
- 用户应能以固定格式导入现有的术语库

## 显示术语

- 用户翻译段落时，会在源文本高亮显示该段落用户已添加的所有术语

## 使用术语

- 用户应能选择术语并在目标文本中插入术语翻译
- 用户应能从多重语义的术语中选择一种插入

# 翻译服务集成 LetsTrans-UR-TI

## 全局翻译

- 用户应能调用现有的翻译 API 对全部文本进行预翻译
- 用户应能更改所使用的 API

## 局部翻译

- 用户应能批量选择段落进行重新预翻译