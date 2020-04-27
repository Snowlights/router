package main

import (
	"errors"
	"fmt"
	"github.com/Snowlights/router/processor"
	"github.com/julienschmidt/httprouter"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
	"time"
)

func main(){
	addr, router := processor.InitRouter()

	tcpAddr, err := net.ResolveTCPAddr("tcp",addr)
	if err != nil{
		panic(err)
		return
	}

	netListen, err := net.Listen(tcpAddr.Network(),tcpAddr.String())
	if err != nil{
		panic(err)
		return
	}

	laddr, err := GetServAddr(netListen.Addr())
	if err != nil{
		netListen.Close()
		panic(err)
		return
	}

	fmt.Printf("listen addr[%s]",laddr)

	mw := nethttp.Middleware(opentracing.GlobalTracer(), httpTrafficLogMiddleware(router.(*httprouter.Router)),
		nethttp.OperationNameFunc(func(r *http.Request) string {
			return "HTTP " + r.Method + ": " + r.URL.Path
		}))

	go func(){
		err := http.Serve(netListen,mw)
		if err != nil{
			panic(err)
		}
	}()

	for true{
		time.Sleep(time.Second)
	}

}

func httpTrafficLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: log before handling business logic
		next.ServeHTTP(w, r)
	})
}

func GetServAddr(a net.Addr) (string,error){
	addr := a.String()
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}
	if len(host) == 0 {
		host = "0.0.0.0"
	}

	ip := net.ParseIP(host)

	if ip == nil {
		return "", fmt.Errorf("ParseIP error:%s", host)
	}
	/*
		fmt.Println("ADDR TYPE", ip,
			"IsGlobalUnicast",
			ip.IsGlobalUnicast(),
			"IsInterfaceLocalMulticast",
			ip.IsInterfaceLocalMulticast(),
			"IsLinkLocalMulticast",
			ip.IsLinkLocalMulticast(),
			"IsLinkLocalUnicast",
			ip.IsLinkLocalUnicast(),
			"IsLoopback",
			ip.IsLoopback(),
			"IsMulticast",
			ip.IsMulticast(),
			"IsUnspecified",
			ip.IsUnspecified(),
		)
	*/

	raddr := addr
	if ip.IsUnspecified() {
		// 没有指定ip的情况下，使用内网地址
		inerip, err := GetInterIp()
		if err != nil {
			return "", err
		}

		raddr = net.JoinHostPort(inerip, port)
	}

	//slog.Tracef("ServAddr --> addr:[%s] ip:[%s] host:[%s] port:[%s] raddr[%s]", addr, ip, host, port, raddr)

	return raddr, nil
}

func GetInterIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println(ipnet.IP.String())
				return ipnet.IP.String(), nil
			}
		}
	}

	/*
		for _, addr := range addrs {
			//fmt.Printf("Inter %v\n", addr)
			ip := addr.String()
			if "10." == ip[:3] {
				return strings.Split(ip, "/")[0], nil
			} else if "172." == ip[:4] {
				return strings.Split(ip, "/")[0], nil
			} else if "196." == ip[:4] {
				return strings.Split(ip, "/")[0], nil
			} else if "192." == ip[:4] {
				return strings.Split(ip, "/")[0], nil
			}

		}
	*/

	return "", errors.New("no inter ip")
}