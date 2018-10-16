FROM scratch
WORKDIR $GOPATH/src/ServiceTest
ADD . $GOPATH/src/ServiceTest
ENTRYPOINT ["./ServiceTest_linux_amd64"]
#http://note.youdao.com/noteshare?id=667520ea387bf2df2872499a4a5139cb&sub=814570E78A7544D8AF178486BDEBCAC0