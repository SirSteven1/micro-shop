# micro-shop
Recently, I was studying microservices and learned about go-zero. Unfortunately, there seems to be no relatively complete open source project based on go-zero in the community. Therefore, I decided to work with everyone to build a complete open source project based on go-zero from scratch.

The design concept of this series is to build a basic high-concurrency microservice mall system based on go-zero. Why did you choose this design concept?
First, the business operation of the e-commerce platform is complex and changeable, and there are many resource sections. Microservices can achieve independent business operation and business decoupling.
2. Better handle high-concurrency scenarios.

This series is a practical series of tutorials for go-zero. Basic knowledge will be less covered, that is, you need to have the following basics:

* Understand the basic grammar and basic usage of GO language
* Understand the use of MySQL
* Understand the basic use of go-zero

### Why use go-zero?
go-zero is a web and rpc framework that integrates various engineering practices. The elastic design ensures the stability of the large-concurrency server and has withstood sufficient actual combat testing.

go-zero includes the minimalist API definition and generation tool goctl, which can generate Go, iOS, Android, Kotlin, Dart, TypeScript, and JavaScript code with one click based on the defined api file, and can run it directly.

Benefits of using go-zero:

* Easily obtain the stability to support tens of millions of daily active services
* Built-in microservice management capabilities such as cascading timeout control, current limiting, adaptive fusing, and adaptive load shedding, without the need for configuration and additional code
* Microservice governance middleware can be seamlessly integrated into other existing frameworks.
* Minimalist API description, one-click generation of code for each end
* Automatically verify the legality of client request parameters
* Extensive microservice governance and concurrency toolkit

I won’t go into details about other benefits of go-zero here. For details, please refer to [go-zero official documentation](https://go-zero.dev/en/docs/tasks).

### Business needs
In order to approach project development in real scenarios, the project is constructed based on real product business requirements.
The business of the mall is relatively large and complex, so not all functions are listed here. Although not fully listed, we will try our best to realize most of the business needs of the mall. In this RC Build version, go-zero The core functions and problems encountered are demonstrated one by one.

The following mind map lists the main functions of the system:
![图片一](https://scarlet-impressive-bonobo-517.mypinata.cloud/ipfs/QmdQJ9Zku25XauLhxB6DzZBwZPHLDMcbeEcUvFN3mEvp3K)

### commands and tools needed
Building the go-zero project requires the use of some tools and commands. Here we prepare in advance:

* Goctl tool (see go-zero official documentation for installation instructions)

>   goctl is a code generation tool under the go-zero microservice framework. Using goctl can significantly improve development efficiency and allow developers to focus their time on business development. Its functions include: api service generation, rpc service generation, model code generation, template management, etc.

* [protoc](https://grpc.io/docs/protoc-installation/) & [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/) (for installation methods, please see go-zero official documentation)

>   protoc is a tool written in C++ that can translate proto files into code in a specified language. In go-zero's microservices, we use grpc for communication between services, and writing grpc requires the use of protoc and the plug-in protoc-gen-go that is translated into go language rpc stub code.
