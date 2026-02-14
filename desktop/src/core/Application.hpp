#pragma once

#include <SDL3/SDL.h>

namespace jr {

class Application {
public:
  void run();
  Application() = default;
  ~Application();

private:
  SDL_Window* window_{};
  SDL_Renderer* renderer_{};
  SDL_Event event_{};
  
  bool running_ = true;
  
  unsigned int window_width_ = 500;
  unsigned int window_height_ = 900;

  void init();
  void init_sdl3();
  void init_window();
  void init_renderer();

  void main_loop();

  void destroy();
  void destroy_renderer();
  void destroy_window();
};

} // namespace jr
