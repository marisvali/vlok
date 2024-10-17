package main

import (
	"embed"
	"fmt"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	. "github.com/marisvali/vlok/ai"
	. "github.com/marisvali/vlok/gamelib"
	. "github.com/marisvali/vlok/world"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	_ "image/png"
	"slices"
)

var BlockSize = I(80)

//go:embed data/*
var embeddedFiles embed.FS

type Gui struct {
	defaultFont        font.Face
	imgFood            *ebiten.Image
	imgCharacter       *ebiten.Image
	imgWall            *ebiten.Image
	imgTextBackground  *ebiten.Image
	imgTextColor       *ebiten.Image
	world              World
	frameIdx           Int
	folderWatcher      FolderWatcher
	textHeight         Int
	guiMargin          Int
	useEmbedded        bool
	buttonRegionWidth  Int
	playSize           Pt
	buttonPause        Rectangle
	buttonNewLevel     Rectangle
	buttonRestartLevel Rectangle
	justPressedKeys    []ebiten.Key // keys pressed in this frame
	mousePt            Pt           // mouse position in this frame
	username           string
	ai                 AI
}

type uploadData struct {
	user        string
	version     int64
	id          uuid.UUID
	playthrough []byte
}

func (g *Gui) JustPressed(key ebiten.Key) bool {
	return slices.Contains(g.justPressedKeys, key)
}

func (g *Gui) JustClicked(button Rectangle) bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		return false
	}
	return button.ContainsPt(g.mousePt)
}

func (g *Gui) UserRequestedPause() bool {
	return g.JustPressed(ebiten.KeyEscape) || g.JustClicked(g.buttonPause)
}

func (g *Gui) UserRequestedNewLevel() bool {
	return g.JustPressed(ebiten.KeyN) || g.JustClicked(g.buttonNewLevel)
}

func (g *Gui) UserRequestedRestartLevel() bool {
	return g.JustPressed(ebiten.KeyR) || g.JustClicked(g.buttonRestartLevel)
}

func (g *Gui) Update() error {
	if g.JustPressed(ebiten.KeyX) {
		return ebiten.Termination
	}

	var input PlayerInput
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		input.Move = true
		input.MovePt = g.ScreenToWorld(g.mousePt)
	}

	// input = g.ai.Step(&g.world)
	g.world.Step(input)

	if g.folderWatcher.FolderContentsChanged() {
		g.loadGuiData()
	}

	g.frameIdx.Inc()
	return nil
}

func (g *Gui) ScreenToWorld(screenPos Pt) (worldPos Pt) {
	// worldPos = (screenPos - guiMargin) * (world.Size / playSize)
	playPos := screenPos.Minus(Pt{g.guiMargin, g.guiMargin})
	x := playPos.X.Times(g.world.Size.X).DivBy(g.playSize.X)
	y := playPos.Y.Times(g.world.Size.Y).DivBy(g.playSize.Y)
	worldPos = Pt{x, y}
	return
}

func (g *Gui) WorldToScreen(worldPos Pt) (screenPos Pt) {
	// screenPos = worldPos * (playSize / world.Size) + guiMargin
	x := worldPos.X.Times(g.playSize.X).DivBy(g.world.Size.X)
	y := worldPos.Y.Times(g.playSize.Y).DivBy(g.world.Size.Y)
	playPos := Pt{x, y}
	screenPos = playPos.Plus(Pt{g.guiMargin, g.guiMargin})
	return
}

func (g *Gui) DrawPlayRegion(screen *ebiten.Image) {
	DrawSprite(screen, g.imgFood, 150, 100, 30, 30)
	DrawSprite(screen, g.imgCharacter, 100, 100, 30, 30)
	DrawSprite(screen, g.imgWall, 50, 50, 30, 30)
}

func (g *Gui) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	{
		upperLeft := Pt{g.guiMargin, g.guiMargin}
		lowerRight := upperLeft.Plus(g.playSize)
		playRegion := SubImage(screen, Rectangle{upperLeft, lowerRight})
		g.DrawPlayRegion(playRegion)
	}

	// buttonRegionX := I(screen.Bounds().Dx()).Minus(g.buttonRegionWidth)
	screenSize := IPt(screen.Bounds().Dx(), screen.Bounds().Dy())
	{
		upperLeft := Pt{ZERO, screenSize.Y.Minus(g.textHeight)}
		// lowerRight := upperLeft.Plus(Pt{buttonRegionX, g.textHeight.DivBy(TWO)})
		lowerRight := Pt{screenSize.X, screenSize.Y.Minus(g.textHeight.DivBy(TWO))}
		textRegion := SubImage(screen, Rectangle{upperLeft, lowerRight})
		textRegion.Fill(color.RGBA{215, 215, 15, 255})
		g.DrawInstructionalText(textRegion)
	}

	{
		// upperLeft := Pt{buttonRegionX, I(screen.Bounds().Dy()).Minus(g.textHeight)}
		// lowerRight := upperLeft.Plus(Pt{I(screen.Bounds().Dx()), g.textHeight})
		upperLeft := Pt{ZERO, screenSize.Y.Minus(g.textHeight.DivBy(TWO))}
		lowerRight := Pt{screenSize.X, screenSize.Y}
		buttonRegion := SubImage(screen, Rectangle{upperLeft, lowerRight})
		buttonRegion.Fill(color.RGBA{5, 215, 215, 255})
		g.DrawButtons(buttonRegion)
	}

	// Output TPS (ticks per second, which is like frames per second).
	ebitenutil.DebugPrint(screen, fmt.Sprintf("ActualTPS: %f", ebiten.ActualTPS()))
}

func (g *Gui) DrawButtons(screen *ebiten.Image) {
	height := I(screen.Bounds().Dy())
	buttonWidth := I(200)

	buttonCols := []color.RGBA{{35, 115, 115, 255}, {65, 215, 115, 255}, {225, 115, 215, 255}}

	buttons := []*ebiten.Image{}
	for i := I(0); i.Lt(I(3)); i.Inc() {
		upperLeft := Pt{buttonWidth.Times(i), ZERO}
		lowerRight := Pt{buttonWidth.Times(i.Plus(ONE)), height}
		button := SubImage(screen, Rectangle{upperLeft, lowerRight})
		button.Fill(buttonCols[i.ToInt()])
		buttons = append(buttons, button)
	}

	textCol := color.RGBA{0, 0, 0, 255}
	g.DrawText(buttons[0], "[ESC] Pause", true, textCol)
	g.DrawText(buttons[1], "[R] Restart level", true, textCol)
	g.DrawText(buttons[2], "[N] New level", true, textCol)

	// Remember the regions so that Update() can react when they're clicked.
	g.buttonPause = FromImageRectangle(buttons[0].Bounds())
	g.buttonRestartLevel = FromImageRectangle(buttons[1].Bounds())
	g.buttonNewLevel = FromImageRectangle(buttons[2].Bounds())
}

func (g *Gui) DrawInstructionalText(screen *ebiten.Image) {
	var message string
	message = "Please, let me eat."

	DrawSprite(screen, g.imgTextBackground, 0, 0,
		float64(screen.Bounds().Dx()),
		float64(screen.Bounds().Dy()))

	var r image.Rectangle
	r.Min = screen.Bounds().Min
	r.Max = image.Point{screen.Bounds().Max.X, r.Min.Y + screen.Bounds().Dy()}
	textBox := screen.SubImage(r).(*ebiten.Image)
	g.DrawText(textBox, message, true, g.imgTextColor.At(0, 0))
}

func (g *Gui) DrawText(screen *ebiten.Image, message string, centerX bool, color color.Color) {
	// Remember that with text there is an origin point for the text.
	// That origin point is kind of the lower-left corner of the bounds of the
	// text. Kind of. Read the BoundString docs to understand, particularly this
	// image:
	// https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/Art/glyphterms_2x.png
	// This means that if you do text.Draw at (x, y), most of the text will
	// appear above y, and a little bit under y. If you want all the pixels in
	// your text to be above y, you should do text.Draw at
	// (x, y - text.BoundString().Max.Y).
	textSize := text.BoundString(g.defaultFont, message)
	var offsetX int
	if centerX {
		offsetX = (screen.Bounds().Dx() - textSize.Dx()) / 2
	} else {
		offsetX = 0
	}
	offsetY := (screen.Bounds().Dy() - textSize.Dy()) / 2
	textX := screen.Bounds().Min.X + offsetX
	textY := screen.Bounds().Max.Y - offsetY - textSize.Max.Y
	text.Draw(screen, message, g.defaultFont, textX, textY, color)
}

func (g *Gui) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Gui) LoadImage(filename string) *ebiten.Image {
	if g.useEmbedded {
		return LoadImageEmbedded(filename, &embeddedFiles)
	} else {
		return LoadImage(filename)
	}
}

func (g *Gui) loadGuiData() {
	// Read from the disk over and over until a full read is possible.
	// This repetition is meant to avoid crashes due to reading files
	// while they are still being written.
	// It's a hack but possibly a quick and very useful one.
	CheckCrashes = false
	for {
		CheckFailed = nil
		g.imgFood = g.LoadImage("data/food.png")
		g.imgCharacter = g.LoadImage("data/character.png")
		g.imgWall = g.LoadImage("data/wall.png")
		g.imgTextBackground = g.LoadImage("data/text-background.png")
		g.imgTextColor = g.LoadImage("data/text-color.png")
		if CheckFailed == nil {
			break
		}
	}
	CheckCrashes = true
}

func main() {
	var g Gui
	g.username = getUsername()

	g.textHeight = I(75)
	g.guiMargin = I(50)
	g.buttonRegionWidth = I(200)
	g.playSize = Pt{I(800), I(800)}
	windowSize := g.playSize
	windowSize.Add(Pt{g.guiMargin.Times(TWO), g.guiMargin})
	windowSize.Y.Add(g.textHeight)
	ebiten.SetWindowSize(windowSize.X.ToInt(), windowSize.Y.ToInt())
	ebiten.SetWindowTitle("Miln")
	ebiten.SetWindowPosition(100, 100)

	g.useEmbedded = !FileExists("data")
	if !g.useEmbedded {
		g.folderWatcher.Folder = "data"
	}
	g.loadGuiData()

	// font
	var err error
	// Load the Arial font
	fontData, err := opentype.Parse(goregular.TTF)
	Check(err)

	g.defaultFont, err = opentype.NewFace(fontData, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	Check(err)

	// Start the game.
	err = ebiten.RunGame(&g)
	Check(err)
}
