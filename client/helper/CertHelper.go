package helper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/credentials"
)

// 获取客户端证书配置
func GetClientCreds() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	if err != nil {
		log.Fatal("LoadX509KeyPair: ", err)
	}
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)
	return credentials.NewTLS(&tls.Config{
		// InsecureSkipVerify: true,                    // 出现错误rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate signed by unknown authority"  暂时设置为true来取消对服务端证书的校验
		Certificates: []tls.Certificate{cert}, // 客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
}
