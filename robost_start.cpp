
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

#include <iostream>
#include <string>

#include <boost/asio.hpp>

int main(int argc, char** argv) {
	boost::asio::io_service io;
	while (true) {
		io.run();
	}
	return 0;
}
