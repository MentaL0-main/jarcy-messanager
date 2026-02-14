#include "core/Application.hpp"

#include <iostream>
#include <cstdlib>
#include <exception>

int main() {
  jr::Application app;

  try {
    app.run();
  } catch (std::exception& error_) {
    std::cerr << error_.what() << std::endl;
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
