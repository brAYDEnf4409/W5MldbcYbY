// 代码生成时间: 2025-10-04 02:23:22
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// AnimationEffect defines the structure for an animation effect.
type AnimationEffect struct {
    Name    string `json:"name"`
    Duration int    `json:"duration"`
    Easing   string `json:"easing"`
}

// AnimationService manages the animation effects.
type AnimationService struct {
    // Store to hold animation effects.
    effects map[string]AnimationEffect
}

// NewAnimationService creates a new AnimationService.
func NewAnimationService() *AnimationService {
    return &AnimationService{
        effects: make(map[string]AnimationEffect),
    }
}

// AddEffect adds a new animation effect to the service.
func (s *AnimationService) AddEffect(name string, effect AnimationEffect) error {
    if _, exists := s.effects[name]; exists {
        return fmt.Errorf("animation effect with name '%s' already exists", name)
    }
    s.effects[name] = effect
    return nil
}

// GetEffect retrieves an animation effect by name.
func (s *AnimationService) GetEffect(name string) (AnimationEffect, error) {
    effect, exists := s.effects[name]
    if !exists {
        return AnimationEffect{}, fmt.Errorf("animation effect '%s' not found", name)
    }
    return effect, nil
}

func main() {
    app := iris.New()
    
    service := NewAnimationService()
    
    // Add some predefined animation effects.
    if err := service.AddEffect("fade-in", AnimationEffect{
        Name:    "fade-in",
        Duration: 500,
        Easing:   "ease-in",
    }); err != nil {
        fmt.Println("Error adding animation effect: ", err)
        return
    }
    
    // Set up routes.
    app.Get("/effects", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "List of animation effects",
            "effects": service.effects,
        })
    })
    
    app.Get("/effects/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        effect, err := service.GetEffect(name)
        if err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.Map{
            "name":    effect.Name,
            "duration": effect.Duration,
            "easing":  effect.Easing,
        })
    })
    
    // Start the server.
    app.Listen(":8080")
}