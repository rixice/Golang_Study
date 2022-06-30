package main

// WithValue 能够将请求作用域的数据与Context对象建立关系
// WithValue(parent context.Context, key any, val any) context.Context
// 所提供的key，应该是可比较的，不该是string或其他内置类型，避免上下文在包之间产生冲突
// WithValue的用户应该为key定义自己的类型，避免在分配给interface{}时进行分配
