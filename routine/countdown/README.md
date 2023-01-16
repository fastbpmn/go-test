## countdown
#### 通过countdown模块演示如何使用模块子包
* 模块子包的package名字要和所属文件夹名字保持一致吗？
```
不需要保持一致，但建议保持一致，因为模块子包的package，是其他模块引入时会自动带出来的别名，为了项目的规范、整洁，强烈建议保持一致。
```
举个例子来说，比如说：模块子包的文件夹名叫：countdown，go文件的package就应该叫：countdown，你叫countdown2也不是不可以，但是显然不符合项目规范，当然了，特殊情况除外，我是强烈建议保持一致的。