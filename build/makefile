
all:test compile docker push clean

test:
	cd ../eTest/build;make test 

compile:
	cd ../eTest/build;make compile 
	cd ../eRedis/build;make compile 
	cd ../eMysql/build;make compile 
	cd ../eConf/build;make compile 
	cd ../eGrpc/build;make compile 
	cd ../eHttp/build;make compile 
	cd ../eFlag/build;make compile 
	cd ../eYaml/build;make compile 
	cd ../eModule/build;make compile 
	cd ../eDocker/build;make compile 

docker: 
	cd ../eTest/build;make docker 
	cd ../eRedis/build;make docker 
	cd ../eMysql/build;make docker 
	cd ../eConf/build;make docker 
	cd ../eGrpc/build;make docker 
	cd ../eHttp/build;make docker 
	cd ../eFlag/build;make docker 
	cd ../eYaml/build;make docker 
	cd ../eModule/build;make docker 
	cd ../eDocker/build;make docker 

	cd ../tool/test/container;make docker 

push: 
	cd ../eTest/build;make push
	cd ../eRedis/build;make push
	cd ../eMysql/build;make push
	cd ../eConf/build;make push
	cd ../eGrpc/build;make push
	cd ../eHttp/build;make push
	cd ../eFlag/build;make push
	cd ../eYaml/build;make push
	cd ../eModule/build;make push
	cd ../eDocker/build;make push

	cd ../tool/test/container;make push 

clean:
	cd ../eTest/build;make clean
	cd ../eRedis/build;make clean
	cd ../eMysql/build;make clean
	cd ../eConf/build;make clean
	cd ../eGrpc/build;make clean
	cd ../eHttp/build;make clean
	cd ../eFlag/build;make clean
	cd ../eYaml/build;make clean
	cd ../eModule/build;make clean
	cd ../eDocker/build;make clean

