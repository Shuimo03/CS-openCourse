pub mod coin;

enum IpAddressKind {
    IpV4,
    IpV6
}

struct IpAddress{
    kind:IpAddressKind,
    address:String,
}


fn main(){
    let home = IpAddress{
        kind:IpAddressKind::IpV4,
        address:String::from("127.0.0.1"),
    };
    
    let loopback = IpAddress{
        kind:IpAddressKind::IpV6,
        address:String::from("::1"),
    };

}