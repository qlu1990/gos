# 狗子
 - 不错这个webframework 就是叫狗子，我的golang 成长之路
# start

```golang
   package main
   import "github.com/qlu1990/gos"
   
   func hello(c *gos.Context) {
	fmt.Fprintln(c.ResponseWriter, "hello world")
    }


   func main(){
       r := gos.NewGos()
       r.AddGet("/hello", hello)
       r.Run(":8000") //runing listen port 8000
   }
  



```



# 路由
  - ~~路由使用map 可以优化 可以参考gin 使用属树~~(已经实现)
  - ~~路由还不能支持路径中带参数~~ (已实现)

  
## 路由函数

```golang
    type HandlerFunc func(*Context)
    r := gos.NewGos()
    r.AddGet(url string ,f HandlerFunc) // 路由get
    r.AddPost(url string ,f HandlerFunc) // 路由post
    r.AddHead(url string ,f HandlerFunc) // 路由head
    r.AddDelete(url string ,f HandlerFunc) // 路由delete

```
 - 请求路径中带参数
  ```golang
    func GetPerson(c *gos.Context) {
	name := c.Param("name")
	gos.Response(c.ResponseWriter,"hello "+name, http.StatusOK)
    }
    func main(){
    r := gos.NewGos()
    r.AddGet("/hello/:name",GetPerson)
    }
    

  ```

# Middleware

## 实现数据结构 

 - `Name` 中间件名称不能重名
 - `HandlerFunc`  被调用调用函数 
```golang
    type Middleware struct {  
	Name        string
	HandlerFunc HandlerFunc      
    }
    
    type HandlerFunc func(*Context)
```

## 使用方法
以glog 为例

```golang
    r := gos.NewGos()
   var Mlog = Middleware{
	Name: "log",
	HandlerFunc: func(c *Context) {
		Glog.Info("Request : ", c.Request.Method, " ", c.Request.RequestURI)
	},
    }
   r.Use(Mlog)

```

#日志模块
## 使用

```golang

   r := gos.NewGos()
   r.Info("info log")
   r.Debug("debug log")
   r.Error("error log")
   r.Fatal（"fatal log")
   r.Warn("warn log")

```

# 案例
[案例 gos-example](https://github.com/qlu1990/gos-example)

email： 876392131@qq.com


