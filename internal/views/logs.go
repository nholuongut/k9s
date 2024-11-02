package views

import (
	"context"
	"fmt"
	"time"

	"github.com/nholuongut/k9s/internal/resource"
	"github.com/nholuongut/k9s/internal/ui"
	"github.com/nholuongut/tview"
	"github.com/gdamore/tcell"
	"github.com/rs/zerolog/log"
)

const (
	logBuffSize  = 100
	flushTimeout = 200 * time.Millisecond

	logCoFmt = " Logs([fg:bg:]%s:[hilite:bg:b]%s[-:bg:-]) "
	logFmt   = " Logs([fg:bg:]%s) "
)

type (
	masterView interface {
		backFn() ui.ActionHandler
		appView() *appView
	}

	logsView struct {
		*tview.Pages

		app        *appView
		parent     loggable
		actions    ui.KeyActions
		cancelFunc context.CancelFunc
	}
)

func newLogsView(title string, app *appView, parent loggable) *logsView {
	v := logsView{
		app:    app,
		Pages:  tview.NewPages(),
		parent: parent,
	}

	return &v
}

// Protocol...

func (v *logsView) reload(co string, parent loggable, prevLogs bool) {
	v.parent = parent
	v.deletePage()
	v.AddPage("logs", newLogView(co, v.app, v.backCmd), true, true)
	v.load(co, prevLogs)
}

// SetActions to handle keyboard events.
func (v *logsView) setActions(aa ui.KeyActions) {
	v.actions = aa
}

// Hints show action hints
func (v *logsView) Hints() ui.Hints {
	l := v.CurrentPage().Item.(*logView)
	return l.actions.Hints()
}

func (v *logsView) backFn() ui.ActionHandler {
	return v.backCmd
}

func (v *logsView) deletePage() {
	v.RemovePage("logs")
}

func (v *logsView) stop() {
	if v.cancelFunc == nil {
		return
	}
	v.cancelFunc()
	log.Debug().Msgf("Canceling logs...")
	v.cancelFunc = nil
}

func (v *logsView) load(container string, prevLogs bool) {
	if err := v.doLoad(v.parent.getSelection(), container, prevLogs); err != nil {
		v.app.Flash().Err(err)
		l := v.CurrentPage().Item.(*logView)
		l.log("😂 Doh! No logs are available at this time. Check again later on...")
		return
	}
	v.app.SetFocus(v)
}

func (v *logsView) doLoad(path, co string, prevLogs bool) error {
	v.stop()

	l := v.CurrentPage().Item.(*logView)
	l.logs.Clear()
	l.setTitle(path, co)

	var ctx context.Context
	ctx = context.WithValue(context.Background(), resource.IKey("informer"), v.app.informer)
	ctx, v.cancelFunc = context.WithCancel(ctx)

	c := make(chan string, 10)
	go updateLogs(ctx, c, l, logBuffSize)

	res, ok := v.parent.getList().Resource().(resource.Tailable)
	if !ok {
		close(c)
		return fmt.Errorf("Resource %T is not tailable", v.parent.getList().Resource())
	}

	if err := res.Logs(ctx, c, v.logOpts(path, co, prevLogs)); err != nil {
		v.cancelFunc()
		close(c)
		return err
	}

	return nil
}

func (v *logsView) logOpts(path, co string, prevLogs bool) resource.LogOptions {
	ns, po := namespaced(path)
	return resource.LogOptions{
		Fqn: resource.Fqn{
			Namespace: ns,
			Name:      po,
			Container: co,
		},
		Lines:    int64(v.app.Config.K9s.LogRequestSize),
		Previous: prevLogs,
	}
}

func updateLogs(ctx context.Context, c <-chan string, l *logView, buffSize int) {
	defer func() {
		log.Debug().Msgf("updateLogs view bailing out!")
	}()
	buff, index := make([]string, buffSize), 0
	for {
		select {
		case line, ok := <-c:
			if !ok {
				log.Debug().Msgf("Closed channel detected. Bailing out...")
				l.flush(index, buff)
				return
			}
			if index < buffSize {
				buff[index] = line
				index++
				continue
			}
			l.flush(index, buff)
			index = 0
			buff[index] = line
			index++
		case <-time.After(flushTimeout):
			l.flush(index, buff)
			index = 0
		case <-ctx.Done():
			return
		}
	}
}

// ----------------------------------------------------------------------------
// Actions...

func (v *logsView) backCmd(evt *tcell.EventKey) *tcell.EventKey {
	v.stop()
	v.parent.switchPage("master")

	return evt
}
