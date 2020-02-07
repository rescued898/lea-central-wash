package gui

import (
	"log"
	"runtime"
	"time"

	"github.com/DiaElectronics/lea-central-wash/cmd/storage/internal/app"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"
)

// Demo ...

func init() {
	runtime.LockOSThread()
}

const (
	// Easy is very easy ...
	Easy Option = 0
	// Hard is nearly impossible ...
	Hard Option = 1
)

func b(v int32) bool {
	return v == 1
}

func flag(v bool) int32 {
	if v {
		return 1
	}
	return 0
}

// State describes the state of the window
type State struct {
	bgColor nk.Color
	prop    int32
	opt     Option
	text    nk.TextEdit
}

// Option describes an option value
type Option uint8

// Demo describes a demo
type Demo struct {
	winWidth  int
	winHeight int

	maxVertexBuffer  int
	maxElementBuffer int
	app              app.App
	info             string
	updInfo          bool
}

// NewDemo is a constructor
func NewDemo(app app.App) *Demo {
	res := &Demo{
		winWidth:         400,
		winHeight:        500,
		maxVertexBuffer:  512 * 1024,
		maxElementBuffer: 128 * 1024,
		app:              app,
	}
	return res
}

// Run runs the demo
func (d *Demo) Run() {

	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(d.winWidth, d.winHeight, "Nuklear Demo", nil, nil)
	if err != nil {
		closer.Fatalln(err)
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	log.Printf("glfw: created window %dx%d", width, height)

	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	// sansFont := nk.NkFontAtlasAddFromBytes(atlas, MustAsset("assets/FreeSans.ttf"), 16, nil)
	// config := nk.NkFontConfig(14)
	// config.SetOversample(1, 1)
	// config.SetRange(nk.NkFontChineseGlyphRanges())
	// simsunFont := nk.NkFontAtlasAddFromFile(atlas, "/Library/Fonts/Microsoft/SimHei.ttf", 14, &config)
	nk.NkFontStashEnd()
	// if simsunFont != nil {
	// 	nk.NkStyleSetFont(ctx, simsunFont.Handle())
	// }

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	state := &State{
		bgColor: nk.NkRgba(28, 48, 62, 255),
	}

	nk.NkTexteditInitDefault(&state.text)

	go d.update()

	fpsTicker := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			d.gfxMain(win, ctx, state)
		}
	}
}

func (d *Demo) update() {
	for {
		d.info = d.app.KasseInfo()
		d.updInfo = true
		time.Sleep(10 * time.Second)
	}
}

func (d *Demo) gfxMain(win *glfw.Window, ctx *nk.Context, state *State) {
	if d.updInfo {
		nk.NkTexteditDelete(&state.text, 0, int32(len(state.text.GetGoString())))
		nk.NkTexteditText(&state.text, d.info, int32(len(d.info)))
		d.updInfo = false
	}

	nk.NkPlatformNewFrame()
	// Layout
	bounds := nk.NkRect(50, 50, 230, 250)
	update := nk.NkBegin(ctx, "Demo", bounds,
		nk.WindowBorder|nk.WindowMovable|nk.WindowScalable|nk.WindowMinimizable|nk.WindowTitle)

	if update > 0 {
		nk.NkLayoutRowStatic(ctx, 30, 80, 1)
		{
			if nk.NkButtonLabel(ctx, "button") > 0 {
				log.Println("[INFO] button pressed!")
			}
		}
		nk.NkLayoutRowDynamic(ctx, 30, 2)
		{
			if nk.NkOptionLabel(ctx, "easy", flag(state.opt == Easy)) > 0 {
				state.opt = Easy
			}
			if nk.NkOptionLabel(ctx, "hard", flag(state.opt == Hard)) > 0 {
				state.opt = Hard
			}
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditBuffer(ctx, nk.EditField, &state.text, nk.NkFilterDefault)
			if nk.NkButtonLabel(ctx, "Print Entered Text") > 0 {
				log.Println(state.text.GetGoString())
			}
		}
		nk.NkLayoutRowDynamic(ctx, 25, 1)
		{
			nk.NkPropertyInt(ctx, "Compression:", 0, &state.prop, 100, 10, 1)
		}
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "background:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 25, 1)
		{
			size := nk.NkVec2(nk.NkWidgetWidth(ctx), 400)
			if nk.NkComboBeginColor(ctx, state.bgColor, size) > 0 {
				nk.NkLayoutRowDynamic(ctx, 120, 1)
				cf := nk.NkColorCf(state.bgColor)
				cf = nk.NkColorPicker(ctx, cf, nk.ColorFormatRGBA)
				state.bgColor = nk.NkRgbCf(cf)
				nk.NkLayoutRowDynamic(ctx, 25, 1)
				r, g, b, a := state.bgColor.RGBAi()
				r = nk.NkPropertyi(ctx, "#R:", 0, r, 255, 1, 1)
				g = nk.NkPropertyi(ctx, "#G:", 0, g, 255, 1, 1)
				b = nk.NkPropertyi(ctx, "#B:", 0, b, 255, 1, 1)
				a = nk.NkPropertyi(ctx, "#A:", 0, a, 255, 1, 1)
				state.bgColor.SetRGBAi(r, g, b, a)
				nk.NkComboEnd(ctx)
			}
		}
	}
	nk.NkEnd(ctx)

	// Render
	bg := make([]float32, 4)
	nk.NkColorFv(bg, state.bgColor)
	width, height := win.GetSize()
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
	nk.NkPlatformRender(nk.AntiAliasingOn, d.maxVertexBuffer, d.maxElementBuffer)
	win.SwapBuffers()
}
