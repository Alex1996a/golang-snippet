### 想看的有关 map 的文章
1、https://blog.golang.org/maps (基础官方文章肯定要过一遍)
2、https://www.jianshu.com/p/f2e7650da938
3、delete key in map: https://stackoverflow.com/questions/1736014/delete-mapkey-in-go
4、https://halfrost.com/go_map_chapter_one/ [设计一个线程安全的 map ]

### 实现 map 
1. 如何设计并实现一个线程安全的 Map ？(下篇)
https://halfrost.com/go_map_chapter_two/

### map 的相关 issue 
1. https://github.com/golang/go/issues/20135 []
runtime: maps do not shrink after elements removal (delete) 。内存泄露
这里可以参考 go-zero 的解决方案。 

2.

