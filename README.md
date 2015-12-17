rslog介绍
====    
   rslog是在[seelog](https://github.com/cihub/seelog)项目的基础上加了支持把日志文件传送到rsyslog中的一个lib.    

rslog使用说明    
====    
只需要在自己的项目中导入rslog模块,并注册一次下面的函数,后期按照直接记录日志即可(使用方法和原[seelog](https://github.com/cihub/seelog)项目一样)     

``` c     
    if ok := log.RegisterRsyslogHandler("./rsyslog.xml"); ok == false {
        fmt.Printf("register rsyslog reveiver handler fail")
    }
```

rsyslog.xml配置文件文件内容:

```c     
<seelog type="sync">
<outputs>
    <custom name="rsyslog" formatid="syslog" data-net="tcp" data-addr="127.0.0.1:5514" data-tag="myrsyslog" />
    <rollingfile formatid="syslog" type="date" filename="./log/test" archivetype="zip" datepattern="2006-01-02" />
    <console formatid="syslog" />
</outputs>
<formats>
    <format id="syslog" format="%File %Func %Line %LEV %Msg%n"/>
</formats>
</seelog>

```    

上面配置中的下面字段就是向rsyslog写日志的配置:      


```    c
<custom name="rsyslog" formatid="syslog" data-net="tcp" data-addr="127.0.0.1:5514" data-tag="myrsyslog" />
```

其中上面参数含义:           

 * data-net:  连接rsyslog使用的套接字类型(tcp,udp)default tcp         
 * data-addr: rsyslog的ip和port (default 127.0.0.1:514)        
 * data-tag:  每一行记录的日志信息的tag标志,default "rslog"      

日志输出格式配置:  

```c    
<format id="syslog" format="%File %Func %Line %LEV %Msg%n"/>
```  
 * %File:  在输出日志段加文件名            
 * %Func:  在输出日志段加调用函数名           
 * %Line:  在输出日志段加调用行数     
 
更多详细配置请看[Format](https://github.com/cihub/seelog/wiki/Format-reference).       







