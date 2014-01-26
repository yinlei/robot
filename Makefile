BOOST_INCLUDE = ../../boost_1_51_0/
BOOST_LIB = ../../boost_1_51_0/stage/lib/
robost-start : robost_start.cpp
	g++ robost_start.cpp -I$(BOOST_INCLUDE)  -L$(BOOST_LIB) -o robost-start -lboost_system

clean :
	rm robost-start
