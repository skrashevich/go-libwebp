# the name of the target operating system
set(CMAKE_SYSTEM_NAME Darwin)

# which compilers to use for C and C++
set(CMAKE_C_COMPILER   o64-clang)
set(CMAKE_CXX_COMPILER o64-clang++)

# adjust the default behavior of the FIND_XXX() commands: # search programs in the host environment
set(CMAKE_FIND_ROOT_PATH_MODE_PROGRAM NEVER)
