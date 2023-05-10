// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "arp"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  bool add_interface(string iface_name)")
  fmt.Fprintln(os.Stderr, "  bool del_interface(string iface_name)")
  fmt.Fprintln(os.Stderr, "  bool add_ip(string iface_name, vlan_tag_t stag, vlan_tag_t ctag, ip4_t ip)")
  fmt.Fprintln(os.Stderr, "  bool del_ip(string iface_name, vlan_tag_t stag, vlan_tag_t ctag)")
  fmt.Fprintln(os.Stderr, "   request_mac( requests)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := arp.NewArpResponderClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "add_interface":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "AddInterface requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.AddInterface(value0))
    fmt.Print("\n")
    break
  case "del_interface":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "DelInterface requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.DelInterface(value0))
    fmt.Print("\n")
    break
  case "add_ip":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "AddIP requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err18 := (strconv.ParseInt(flag.Arg(2), 10, 16))
    if err18 != nil {
      Usage()
      return
    }
    argvalue1 := int16(tmp1)
    value1 := arp.VlanTagT(argvalue1)
    tmp2, err19 := (strconv.ParseInt(flag.Arg(3), 10, 16))
    if err19 != nil {
      Usage()
      return
    }
    argvalue2 := int16(tmp2)
    value2 := arp.VlanTagT(argvalue2)
    tmp3, err20 := (strconv.ParseInt(flag.Arg(4), 10, 32))
    if err20 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := arp.Ip4T(argvalue3)
    fmt.Print(client.AddIP(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "del_ip":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "DelIP requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err22 := (strconv.ParseInt(flag.Arg(2), 10, 16))
    if err22 != nil {
      Usage()
      return
    }
    argvalue1 := int16(tmp1)
    value1 := arp.VlanTagT(argvalue1)
    tmp2, err23 := (strconv.ParseInt(flag.Arg(3), 10, 16))
    if err23 != nil {
      Usage()
      return
    }
    argvalue2 := int16(tmp2)
    value2 := arp.VlanTagT(argvalue2)
    fmt.Print(client.DelIP(value0, value1, value2))
    fmt.Print("\n")
    break
  case "request_mac":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RequestMac requires 1 args")
      flag.Usage()
    }
    arg24 := flag.Arg(1)
    mbTrans25 := thrift.NewTMemoryBufferLen(len(arg24))
    defer mbTrans25.Close()
    _, err26 := mbTrans25.WriteString(arg24)
    if err26 != nil { 
      Usage()
      return
    }
    factory27 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt28 := factory27.GetProtocol(mbTrans25)
    containerStruct0 := arp.NewArpResponderRequestMacArgs()
    err29 := containerStruct0.ReadField1(jsProt28)
    if err29 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Requests
    value0 := argvalue0
    fmt.Print(client.RequestMac(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
