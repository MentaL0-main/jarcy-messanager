cmake_minimum_required(VERSION 3.16)

project(
  jarcy
  VERSION 0.1
  LANGUAGES CXX
)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

find_package(SDL3 3.0 QUIET)

if (NOT SDL3_FOUND)
  find_package(PkgConfig QUIET)
  if (PKG_CONFIG_FOUND)
    pkg_check_modules(SDL3_PKG QUIET sdl3)
    if (SDL3_PKG_FOUND)
      add_library(SDL3::SDL3 UNKNOWN IMPORTED)
      set_target_properties(SDL3::SDL3 PROPERTIES
        INTERFACE_INCLUDE_DIRECTORIES "${SDL3_PKG_INCLUDEDIR}"
        INTERFACE_LINK_LIBRARIES "${SDL3_PKG_LIBRARIES}"
      )
      set(SDL3_FOUND TRUE)
    endif()
  endif()
endif()

if (NOT SDL3_FOUND)
  message(FATAL_ERROR "SDL3 not found. Install SDL3 development package or set SDL3_DIR (CMake config).")
endif()

add_executable(${PROJECT_NAME}
  src/main.cpp
  src/core/Application.cpp
  src/auth/Auth.cpp)

target_link_libraries(${PROJECT_NAME} PRIVATE SDL3::SDL3)

if (TARGET SDL3::SDL3main)
  target_link_libraries(${PROJECT_NAME} PRIVATE SDL3::SDL3main)
endif()

install(TARGETS ${PROJECT_NAME} RUNTIME DESTINATION bin)

