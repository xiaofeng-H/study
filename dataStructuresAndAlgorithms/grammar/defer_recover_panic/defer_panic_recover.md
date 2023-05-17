# defer

## 一、定义

defer 能够让我们推迟执行某些函数调用，推迟到当前函数返回前才实际执行。defer与panic和recover结合，形成了Go语言风格的异常与捕获机制。

## 二、使用场景

defer 语句经常被用于处理成对的操作，如文件句柄关闭、连接关闭、释放锁
1、资源释放（文件、数据库连接）

2、异常捕获和处理

## 三、优点

方便开发者使用

## 四、缺点

有性能损耗

## 五、实现原理：

Go1.14中编译器会将defer函数直接插入到函数的尾部，无需链表和栈上参数拷贝，性能大幅提升。把defer函数在当前函数内展开并直接调用，这种方式被称为open
coded defer

## 六、defer、panic、recover特性（代码演示）

1、defer方法运行在当前方法return之后，调用方拿到返回值之前

2、函数退出前，按照先进后出的顺序，执行defer函数（Test2）

3、panic后的defer函数不会被执行（遇到panic，如果没有捕获错误，函数会立刻终止）（Test3）

4、panic没有被recover时，抛出的panic到当前goroutine最上层函数时，最上层程序直接异常终止（Test4）

5、panic有被recover时，当前goroutine最上层函数正常执行（Test5）

6、defer参数预处理
