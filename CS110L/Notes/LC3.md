## Ownership (continued) and Error Handling

可以参考Rust权威指南的第四章和第九章

### C中的所有权

在调查结果中，许多人对Rust的所有权模型( ownership model)感到疑惑，I would like to make the claim that Rust code actually does the same thing as *good* C code! If you read production C code, you’ll find notions of ownership embedded in the comments.

这里有一些来至 [Open vSwitch](https://github.com/openvswitch/ovs/blob/134e6831acca48f10df3d59b8e1567c24dd925d2/ofproto/ofproto-provider.h#L1094) 的代码：

```c
/* Get status of the virtual port (ex. tunnel, patch).
 *
 * Returns '0' if 'port' is not a virtual port or has no errors.
 * Otherwise, stores the error string in '*errp' and returns positive errno
 * value. The caller is responsible for freeing '*errp' (with free()).
 *
 * This function may be a null pointer if the ofproto implementation does
 * not support any virtual ports or their states.
 */
int (*vport_get_status)(const struct ofport *port, char **errp);
```

还有一份来自 [ffmpeg](https://github.com/FFmpeg/FFmpeg/blob/f1894c206eec463832eef851a5388949a68a050f/libavutil/opt.h#L695)的例子，一个受欢迎的媒体转二进制库

```c
/**
 * @note Any old dictionary present is discarded and replaced with a copy of the new one. The
 * caller still owns val is and responsible for freeing it.
 */
int av_opt_set_dict_val(void *obj, const char *name, const AVDictionary *val, int search_flags);
```

### 通过例子来理解所有权

```rust
fn om_nom_nom(s:String){
    println!("{}",s);
}

fn main() {
    
    let s = String::from("hello");
    //s.push_str(" world");
    om_nom_nom(s);
    om_nom_nom(s);
}
```

这是一个很简单的例子，在main调用两次om_nom_nom函数，其中第一个om_nom_nom拿走了s的所有权，并且调用完函数后，s的所有权被删除了，它的内存也没了，所以之后对s的操作就无效了。

这里再来看一个u32类型的om_nom_nom，同样是调用两次，结果却不一样，两次都可以打印。

```rust
fn om_nom_nom(param: u32) {
    println!("{}", param);
}

fn main() {
    let x = 1;
    om_nom_nom(x);
    om_nom_nom(x);
```

原因在于u32实现了copy trait，

### copy trait

### References

### read-only pointers

## Error handling