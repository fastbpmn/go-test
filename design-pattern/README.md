## design-pattern
#### 试验go的语言特性
* 本地还没有提交到github的模块可以被其他模块引用吗？
```
可以，在godoc模块引入routine模块，2个模块都没有提交到github，也是可以引用的。
```
需要注意的是，此时godoc模块不能使用go mod tidy，因为go mod tidy需要从github下载routine模块，而routine模块还属于本地。