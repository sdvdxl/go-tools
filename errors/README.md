# 错误
包括不可变错误 `ConstError` 和 用用于http api 返回的 `Auth`,`Dumplicated`,`IllegalArgument`和`NotFound` error。

## ConstError
该error不可变，`errors.ConstError("错误")` 即可得到一个不可变的error对象。

## CodeError
该error 含有 `code` 和 `msg` 两个参数，code 用于返回具体的错误代码，msg  其他错误的结构体组合了 CodeError结构体，同时提供了NewXXX的方法，比如 要得到一个Auth error，`errors.NewAuth(401,"密码错误")` 即可。