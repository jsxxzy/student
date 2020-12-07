# 极域电子教室管理

主要是用来杀掉`极域电子教室`软件的客户端,方便摸鱼. 原理很简单的, 调用 `ntsd.exe` 来杀死进程

```bash
ntsd -c q -pn studentmain.exe
```

使用

```bash
go build -o kill.exe .

# 运行
kill run

# 杀掉进程
kill kill
```