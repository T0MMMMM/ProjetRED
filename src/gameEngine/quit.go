package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *EngineStruct) quit() {// on unload toute les textures et ferme la fenetre de jeu
	rl.UnloadTexture(engine.sprite.grass)
	rl.UnloadTexture(engine.sprite.donjon)
	rl.UnloadTexture(engine.sprite.donjon2)
	rl.UnloadTexture(engine.sprite.house)
	rl.UnloadTexture(engine.sprite.water)
	rl.UnloadTexture(engine.sprite.tilled)
	rl.UnloadTexture(engine.sprite.invBar)

	rl.UnloadTexture(engine.player.Sprite)

	rl.UnloadMusicStream(engine.music)
	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}
