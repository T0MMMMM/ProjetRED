package gameEngine

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *EngineStruct) drawScene() {
	rl.BeginDrawing()
	rl.ClearBackground(engine.bgColor)

	rl.BeginMode2D(engine.cam)


	engine.bord.colisionList = [][]float32{}
	for i := 0; i < len(engine.bord.tileMap); i++ {
		if engine.bord.tileMap[i] != 0 {
			engine.bord.tileDest.X = engine.bord.tileDest.Width * float32(i%engine.bord.mapW)
			engine.bord.tileDest.Y = engine.bord.tileDest.Height * float32(i/engine.bord.mapW)

			if engine.bord.srcMap[i] == "g" {
				engine.sprite.texture = engine.sprite.grass
				engine.bord.tileSrc.X = 0
				engine.bord.tileSrc.Y = 0
				rl.DrawTexturePro(engine.sprite.grass, rl.NewRectangle(16*11, 32, 16, 16), engine.bord.tileDest, rl.NewVector2(engine.bord.tileDest.Width, engine.bord.tileDest.Height), 0, rl.White)
			}
			if engine.bord.srcMap[i] == "d" {
				engine.sprite.texture = engine.sprite.donjon
			}
			if !engine.doorOpen {
				if engine.bord.srcMap[i] == "f" {
					engine.sprite.texture = engine.sprite.donjon2
					engine.bord.colisionList = append(engine.bord.colisionList, []float32{engine.bord.tileDest.X, engine.bord.tileDest.Y})
				}
			} else {
				if engine.bord.srcMap[i] == "f" {
					engine.sprite.texture = engine.sprite.house
				}
			}
			if engine.bord.srcMap[i] == "h" {
				engine.sprite.texture = engine.sprite.house
			}
			if engine.bord.srcMap[i] == "w" {
				engine.sprite.texture = engine.sprite.water
				engine.bord.colisionList = append(engine.bord.colisionList, []float32{engine.bord.tileDest.X, engine.bord.tileDest.Y})
			}
			if engine.bord.srcMap[i] == "t" {
				engine.sprite.texture = engine.sprite.tilled
				engine.bord.tileSrc.X = 0
				engine.bord.tileSrc.Y = 0
				rl.DrawTexturePro(engine.sprite.tilled, rl.NewRectangle(16*11, 32, 16, 16), engine.bord.tileDest, rl.NewVector2(engine.bord.tileDest.Width, engine.bord.tileDest.Height), 0, rl.White)
			}
			
			if engine.bord.srcMap[i] == "h" || engine.bord.srcMap[i] == "f" { // si il y a une barrière ou une maison on met de l'herbe en dessous
                engine.bord.tileSrc.X = 0
                engine.bord.tileSrc.Y = 0
                rl.DrawTexturePro(engine.sprite.donjon, rl.NewRectangle(16, 64, 16, 16), engine.bord.tileDest, rl.NewVector2(engine.bord.tileDest.Width, engine.bord.tileDest.Height), 0, rl.White)
            }
			engine.bord.tileSrc.X = engine.bord.tileSrc.Width * float32((engine.bord.tileMap[i]-1)%int(engine.sprite.texture.Width/int32(engine.bord.tileSrc.Width)))
			engine.bord.tileSrc.Y = engine.bord.tileSrc.Height * float32((engine.bord.tileMap[i]-1)/int(engine.sprite.texture.Width/int32(engine.bord.tileSrc.Width)))
			rl.DrawTexturePro(engine.sprite.texture, engine.bord.tileSrc, engine.bord.tileDest, rl.NewVector2(engine.bord.tileDest.Width, engine.bord.tileDest.Height), 0, rl.White)
		}
	}

	for i := 0; i < len(engine.monster); i++ {
		if engine.monster[i].alive {
			rl.DrawTexturePro(engine.monster[i].sprite, rl.NewRectangle(engine.monster[i].Src.X+float32(engine.monster[i].Xstart), engine.monster[i].Src.Y, engine.monster[i].Src.Width, engine.monster[i].Src.Height), engine.monster[i].Dest, rl.NewVector2(engine.monster[i].Dest.Width, engine.monster[i].Dest.Height), 0, rl.White)
		}
	}

	rl.DrawTexturePro(engine.shopKeeper.sprite, engine.shopKeeper.Src, engine.shopKeeper.Dest, rl.NewVector2(engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Height), 0, rl.White)
	rl.DrawTextEx(engine.fontNum, "Shop", rl.NewVector2(float32(engine.shopKeeper.Dest.X-32), float32(engine.shopKeeper.Dest.Y-40)), 8, 0, rl.Black)
	rl.DrawTexturePro(engine.shopKeeper.sprite, engine.shopKeeper.Src, rl.NewRectangle(686, 1153, 32, 32), rl.NewVector2(engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Height), 0, rl.White)
	rl.DrawTextEx(engine.fontNum, "Shop", rl.NewVector2(float32(686-32), float32(1153-40)), 8, 0, rl.Black)
	rl.DrawTexturePro(engine.shopKeeper.sprite, engine.shopKeeper.Src, rl.NewRectangle(594, 780, 32, 32), rl.NewVector2(engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Height), 0, rl.White)
	rl.DrawTextEx(engine.fontNum, "Shop", rl.NewVector2(float32(594-32), float32(780-40)), 8, 0, rl.Black)
	rl.DrawTexturePro(engine.shopKeeper.sprite, engine.shopKeeper.Src, rl.NewRectangle(1030, 340, 32, 32), rl.NewVector2(engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Height), 0, rl.White)
	rl.DrawTextEx(engine.fontNum, "Shop", rl.NewVector2(float32(1030-32), float32(340-40)), 8, 0, rl.Black)

	rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.character.hp)), rl.NewVector2(float32(engine.player.Dest.X)-183, float32(engine.player.Dest.Y)-113), 10, 0, rl.Black)
	if engine.character.gold >= 0 && engine.character.gold < 10 {
		rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.character.gold)), rl.NewVector2(float32(engine.player.Dest.X)+155, float32(engine.player.Dest.Y)-113), 10, 0, rl.Black)
	}
	if engine.character.gold >= 10 && engine.character.gold < 100 {
		rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.character.gold)), rl.NewVector2(float32(engine.player.Dest.X)+145, float32(engine.player.Dest.Y)-113), 10, 0, rl.Black)
	}
	if engine.character.gold < 1000 && engine.character.gold >= 100{
		rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.character.gold)), rl.NewVector2(float32(engine.player.Dest.X)+135, float32(engine.player.Dest.Y)-113), 10, 0, rl.Black)
	}
	if engine.character.gold > 999 {
		rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.character.gold)), rl.NewVector2(float32(engine.player.Dest.X)+125, float32(engine.player.Dest.Y)-113), 10, 0, rl.Black)
	}
	rl.DrawTexture(engine.sprite.money, int32(engine.player.Dest.X)+165, int32(engine.player.Dest.Y)-118, rl.White)
	rl.DrawTexture(engine.sprite.redHeart, int32(engine.player.Dest.X)-202, int32(engine.player.Dest.Y)-116, rl.White)

	rl.DrawTexturePro(engine.player.Sprite, engine.player.Src, engine.player.Dest, rl.NewVector2(engine.player.Dest.Width, engine.player.Dest.Height), 0, rl.White)
	if engine.shop && engine.character.showInventory {
		rl.DrawTexture(engine.sprite.layer, int32(engine.player.Dest.X)-300, int32(engine.player.Dest.Y)-300, rl.White)
		rl.DrawTexture(engine.shopKeeper.shopSprite, int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2 - int32(engine.player.Dest.Width / 2) + 12, int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2- int32(engine.player.Dest.Height / 2) - 100+50, rl.White)
		t := 0
		slot2 := 0
		for i := 0; i < len(engine.shopKeeper.items); i++ {
			if i == 3 || i == 6 {
				slot2 += 1
				t = 0
			}
			rl.DrawTexture(engine.shopKeeper.items[i].sprite, (int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2)+10+int32(23*t) - int32(engine.player.Dest.Width / 2) + 12, (int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2)+10+int32(23*slot2)- int32(engine.player.Dest.Height / 2) - 100+50, rl.White)
			t++
		}
		rl.DrawTextEx(engine.fontText, "Shop", rl.NewVector2(float32(engine.player.Dest.X-170), float32(engine.player.Dest.Y-110)+50), 30, 0, rl.Black)
		//rl.DrawTextEx(engine.fontNum, strconv.Itoa(engine.character.gold), rl.NewVector2(float32(engine.player.Dest.X)+178, float32(engine.player.Dest.Y)-133+50), 10, 0, rl.Black)
		rl.DrawTexture(engine.sprite.money, int32(engine.player.Dest.X)+165, int32(engine.player.Dest.Y)-118, rl.White)
		if engine.shopKeeper.showPrice[0] == 1 && len(engine.shopKeeper.items) > engine.shopKeeper.showPrice[1] {
			
			if engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price >= 10 && engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price < 100 {
				rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price)), rl.NewVector2(float32(engine.player.Dest.X-47+20), float32(engine.player.Dest.Y-50)+50), 10, 0, rl.Black)
			}
			if engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price < 1000 && engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price >= 100{
				rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price)), rl.NewVector2(float32(engine.player.Dest.X-47+10), float32(engine.player.Dest.Y-50)+50), 10, 0, rl.Black)
			}
			if engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price > 999 {
				rl.DrawTextEx(engine.fontNum, strconv.Itoa(int(engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].price)), rl.NewVector2(float32(engine.player.Dest.X-47), float32(engine.player.Dest.Y-50)+50), 10, 0, rl.Black)
			}
			rl.DrawTexture(engine.sprite.money, int32(engine.player.Dest.X)-7, int32(engine.player.Dest.Y)-5, rl.White)
			rl.DrawTextEx(engine.fontText, engine.shopKeeper.items[engine.shopKeeper.showPrice[1]].description, rl.NewVector2(float32(engine.player.Dest.X-48)-50, float32(engine.player.Dest.Y+25)+50), 10, 0, rl.Black)
		}
		//rl.DrawTexture(engine.sprite.layer, int32(engine.player.Dest.X)-300, int32(engine.player.Dest.Y)-300+50, rl.White)
		rl.DrawTexture(engine.sprite.invBar, int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2 - int32(engine.player.Dest.Width / 2), int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2- int32(engine.player.Dest.Height / 2)+50, rl.White)
		t = 0
		slot2 = 0
		for i := 0; i < len(engine.character.inventory); i++ {
			if i == 4 {
				slot2 = 1
				t = 0
			}
			rl.DrawTexture(engine.character.inventory[i].sprite, (int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2)+10+int32(23*t)- int32(engine.player.Dest.Width / 2), (int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2)+10+int32(23*slot2)- int32(engine.player.Dest.Height / 2)+50, rl.White)
			t++
		}
	} else if engine.character.showInventory {
		rl.DrawTexture(engine.sprite.layer, int32(engine.player.Dest.X)-300, int32(engine.player.Dest.Y)-300, rl.White)
		rl.DrawTexture(engine.sprite.invBar, int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2 - int32(engine.player.Dest.Width / 2), int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2- int32(engine.player.Dest.Height / 2), rl.White)
		t := 0
		slot2 := 0
		for i := 0; i < len(engine.character.inventory); i++ {
			if i == 4 {
				slot2 = 1
				t = 0
			}
			rl.DrawTexture(engine.character.inventory[i].sprite, (int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2)+10+int32(23*t)- int32(engine.player.Dest.Width / 2), (int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2)+10+int32(23*slot2)- int32(engine.player.Dest.Height / 2), rl.White)
			t++
		}
	}

	if engine.character.showText {
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*4, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X+48-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X+48*2-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X+48*3-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X+48*4-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X+48*5-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*6, 0, 16, 16*2), rl.NewRectangle(engine.player.Dest.X+48*6-130, engine.player.Dest.Y+100, 48, 32*2), rl.NewVector2(48, 32*2), 0, rl.White)
		if engine.textBox.frameCountSpace % 2 == 0 {
			rl.DrawTexturePro(engine.textBox.space, rl.NewRectangle(12*4, 12*5, 12, 12), rl.NewRectangle(engine.player.Dest.X+48*6-166, engine.player.Dest.Y+80, 12*2, 12*2), rl.NewVector2(12, 12), 0, rl.White)
		} else {
			rl.DrawTexturePro(engine.textBox.space, rl.NewRectangle(12*5, 12*5, 12, 12), rl.NewRectangle(engine.player.Dest.X+48*6-166, engine.player.Dest.Y+80, 12*2, 12*2), rl.NewVector2(12, 12), 0, rl.White)
		}
		if engine.textBox.frameCountText == len(engine.textBox.textPrint) && len(engine.textBox.textToPrint) > len(engine.textBox.textPrint) {
			engine.textBox.textPrint += string(engine.textBox.textToPrint[engine.textBox.frameCountText])
		}
		if engine.textBox.textPrint == engine.textBox.textToPrint {
			engine.textBox.textWriting = false
		}
		if len(engine.textBox.textToPrint) > len(engine.textBox.textPrint) || !engine.textBox.textWriting {
			rl.DrawTextEx(engine.fontText, engine.textBox.textPrint, rl.NewVector2(float32(engine.player.Dest.X-150), float32(engine.player.Dest.Y+55)), 11, 0, rl.Black)
		}
	}
	

	rl.EndMode2D()
	rl.EndDrawing()
}
