## Go中string与[]byte的两种转换方式

- 标准转换

```go
s1 := "hello"
b := []byte(s1)

// []byte to string
s2 := string(b)
```

* 强转换（通过unsafe和reflect包，可以实现另外一种转换方式，我们将之称为强转换（也常常被人称作黑魔法））

```go
func String2Bytes(s string) []byte {
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh := reflect.SliceHeader{
        Data: sh.Data,
        Len:  sh.Len,
        Cap:  sh.Len,
    }
    return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}
```

**强转换方式的性能会明显优于标准转换**（为什么）



## String和[]byte在go中到底是什么？

* []byte

在go中，byte是uint8的别名，在go标准库builtin中有如下说明：

```go
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8
```

在go的源码中`src/runtime/slice.go`，slice的定义如下：

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

array是底层数组的指针，len表示长度，cap表示容量。对于[]byte来说，array指向的就是byte数组。

* string

关于string类型，在go标准库builtin中有如下说明：

```go
// string is the set of all strings of 8-bit bytes, conventionally but not
// necessarily representing UTF-8-encoded text. A string may be empty, but
// not nil. Values of string type are immutable.
type string string
```

翻译过来就是：string是8位字节的集合，通常但不一定代表UTF-8编码的文本。string可以为空，但是不能为nil。**string的值是不能改变的。**

在go的源码中`src/runtime/string.go`，string的定义如下：

```go
type stringStruct struct {
    str unsafe.Pointer
    len int
}
```

stringStruct代表的就是一个string对象，str指针指向的是某个数组的首地址，len代表的数组长度。那么这个数组是什么呢？我们可以在实例化stringStruct对象时找到答案。

```go
//go:nosplit
func gostringnocopy(str *byte) string {
    ss := stringStruct{str: unsafe.Pointer(str), len: findnull(str)}
    s := *(*string)(unsafe.Pointer(&ss))
    return s
}
```

可以看到，入参str指针就是指向byte的指针，那么我们可以确定string的底层数据结构就是byte数组。

综上，string与[]byte在底层结构上是非常的相近（后者的底层表达仅多了一个cap属性，因此它们在内存布局上是可对齐的），这也就是为何builtin中内置函数copy会有一种特殊情况`copy(dst []byte, src string) int`的原因了。

```go
// The copy built-in function copies elements from a source slice into a
// destination slice. (As a special case, it also will copy bytes from a
// string to a slice of bytes.) The source and destination may overlap. Copy
// returns the number of elements copied, which will be the minimum of
// len(src) and len(dst).
func copy(dst, src []Type) int
```



- 区别

对于[]byte与string而言，两者之间最大的区别就是string的值不能改变。这该如何理解呢？下面通过两个例子来说明。

对于[]byte来说，以下操作是可行的：

```go
b := []byte("Hello Gopher!")
    b [1] = 'T'
```

string，修改操作是被禁止的：

```go
s := "Hello Gopher!"
    s[1] = 'T'
```

而string能支持这样的操作：

```go
s := "Hello Gopher!"
    s = "Tello Gopher!"
```

字符串的值不能被更改，但可以被替换。 string在底层都是结构体`stringStruct{str: str_point, len: str_len}`，string结构体的str指针指向的是一个字符常量的地址， 这个地址里面的内容是不可以被改变的，因为它是只读的，但是这个指针可以指向不同的地址。

那么，以下操作的含义是不同的：

```go
s := "S1" // 分配存储"S1"的内存空间，s结构体里的str指针指向这块内存
s = "S2"  // 分配存储"S2"的内存空间，s结构体里的str指针转为指向这块内存

b := []byte{1} // 分配存储'1'数组的内存空间，b结构体的array指针指向这个数组。
b = []byte{2}  // 将array的内容改为'2'
```

因为string的指针指向的内容是不可以更改的，所以每更改一次字符串，就得重新分配一次内存，之前分配的空间还需要gc回收，这是导致string相较于[]byte操作低效的根本原因。

## Q&A

**Q1. 为啥强转换性能会比标准转换好？**

对于标准转换，无论是从[]byte转string还是string转[]byte都会涉及底层数组的拷贝。而强转换是直接替换指针的指向，从而使得string和[]byte指向同一个底层数组。这样，当然后者的性能会更好。

**Q2. 为啥在上述测试中，当x的数据较大时，标准转换方式会有一次分配内存的操作，从而导致其性能更差，而强转换方式却不受影响？**

标准转换时，当数据长度大于32个字节时，需要通过mallocgc申请新的内存，之后再进行数据拷贝工作。而强转换只是更改指针指向。所以，当转换数据较大时，两者性能差距会愈加明显。

**Q3. 既然强转换方式性能这么好，为啥go语言提供给我们使用的是标准转换方式？**

首先，我们需要知道Go是一门类型安全的语言，而安全的代价就是性能的妥协。但是，性能的对比是相对的，这点性能的妥协对于现在的机器而言微乎其微。另外强转换的方式，会给我们的程序带来极大的安全隐患。

如下示例

```go
a := "hello"
b := String2Bytes(a)
b[0] = 'H'
```

a是string类型，前面我们讲到它的值是不可修改的。通过强转换将a的底层数组赋给b，而b是一个[]byte类型，它的值是可以修改的，所以这时对底层数组的值进行修改，将会造成严重的错误（通过defer+recover也不能捕获）。

**Q4. 为啥string要设计为不可修改的？**

我认为有必要思考一下该问题。string不可修改，意味它是只读属性，这样的好处就是：在并发场景下，我们可以在不加锁的控制下，多次使用同一字符串，在保证高效共享的情况下而不用担心安全问题