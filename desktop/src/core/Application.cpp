#include "Application.hpp"
#include <SDL3/SDL_init.h>
#include <SDL3/SDL_mutex.h>
#include <SDL3/SDL_render.h>
#include <SDL3/SDL_video.h>
#include <stdexcept>

namespace jr {

void Application::run() {
  init();
  main_loop();
}

void Application::init() {
  init_sdl3();
  init_window();
  init_renderer();
}

void Application::init_sdl3() {
  if (!SDL_Init(SDL_INIT_VIDEO)) {
    throw std::runtime_error("Failed to init sdl3");
  }
}

void Application::init_window() {
  window_ = SDL_CreateWindow("Jarcy", window_width_, window_height_, SDL_WINDOW_RESIZABLE);
  if (!window_) {
    throw std::runtime_error("Failed to init window");
  }
}

void Application::init_renderer() {
  renderer_ = SDL_CreateRenderer(window_, nullptr);
  if (!renderer_) {
    throw std::runtime_error("Failed to init renderer");
  }
}

void Application::main_loop() {
  while (running_) {
    while (SDL_PollEvent(&event_)) {
      if (event_.type == SDL_EVENT_QUIT) {
        running_ = false;
      }
    }
  
    SDL_SetRenderDrawColor(renderer_, 255, 255, 255, 255);
    SDL_RenderClear(renderer_);
  

    SDL_RenderPresent(renderer_);
  }
}

void Application::destroy_renderer() {
  if (renderer_) {
    SDL_DestroyRenderer(renderer_);
  }
}

void Application::destroy_window() {
  if (window_) {
    SDL_DestroyWindow(window_);
  }
}

void Application::destroy() {
  destroy_renderer();
  destroy_window();
  SDL_Quit();
}

Application::~Application() {
  destroy();
}

} // namespace jr
