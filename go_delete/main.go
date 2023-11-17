package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

//	func MouseDragNode(n *cdp.Node, cxt context.Context) error {
//		boxes, err := dom.GetContentQuads().WithNodeID(n.NodeID).Do(cxt)
//		if err != nil {
//			return err
//		}
//		if len(boxes) == 0 {
//			return chromedp.ErrInvalidDimensions
//		}
//		content := boxes[0]
//		c := len(content)
//		if c%2 != 0 || c < 1 {
//			return chromedp.ErrInvalidDimensions
//		}
//		var x, y float64
//		for i := 0; i < c; i += 2 {
//			x += content[i]
//			y += content[i+1]
//		}
//		x /= float64(c / 2)
//		y /= float64(c / 2)
//		p := &input.DispatchMouseEventParams{
//			Type:       input.MousePressed,
//			X:          x,
//			Y:          y,
//			Button:     input.Left,
//			ClickCount: 1,
//		}
//		// 鼠标左键按下
//		if err := p.Do(cxt); err != nil {
//			return err
//		}
//		// 拖动
//		p.Type = input.MouseMoved
//		max := 380.0
//		for {
//			if p.X > max {
//				break
//			}
//			rt := rand.Intn(20) + 20
//			chromedp.Run(cxt, chromedp.Sleep(time.Millisecond*time.Duration(rt)))
//			x := rand.Intn(2) + 15
//			y := rand.Intn(2)
//			p.X = p.X + float64(x)
//			p.Y = p.Y + float64(y)
//			//fmt.Println("X坐标：",p.X)
//			if err := p.Do(cxt); err != nil {
//				return err
//			}
//		}
//		// 鼠标松开
//		p.Type = input.MouseReleased
//		return p.Do(cxt)
//	}
func main() {
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
	path := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	conf := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(path),
		chromedp.Flag("headless", false),                       //选择是否隐藏浏览器
		chromedp.Flag("hide-scrollbars", false),                //浏览器中的滚动条是否隐藏
		chromedp.Flag("mute-audio", false),                     //是否静音
		chromedp.Flag("blink-settings", "imagesEnabled=false"), //不显示图片
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.UserAgent(ua),
		//chromedp.NoSandbox, //禁用浏览器的沙箱功能
	)
	Exec_instances, cancerl := chromedp.NewExecAllocator(context.Background(), conf...)
	defer cancerl()
	instance, cancerl := chromedp.NewContext(Exec_instances)
	defer cancerl()
	err := chromedp.Run(instance,
		chromedp.ActionFunc(func(cxt context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(cxt)
			return err
		}),
		chromedp.Navigate("https://www.aliyundrive.com/sign/in?spm=aliyundrive.index.0.0.1b026f60isxakp"),
		chromedp.Click("#login > div.login-content.nc-outer-box > div > div.login-blocks.block0 > a"),
		chromedp.Sleep(time.Second*3),
		chromedp.Click("#fm-login-id"),
		chromedp.Sleep(time.Second*3),
		//在引号内写你的账号/手机号
		chromedp.WaitVisible("#fm-login-id"),
		chromedp.SendKeys("#fm-login-id", "13206391063"),
		chromedp.Sleep(time.Second*5),
		chromedp.Click("#fm-login-password"),
		//在引号里写你的密码
		chromedp.SendKeys("#fm-login-password", "20040115bnh"),
		chromedp.Click(".fm-submit.password-login"),
		chromedp.Sleep(time.Second*5),
		chromedp.Click("#adrive-nav-sub-tab-container > ul > li:nth-child(2) > div > span"),
		chromedp.Click("#root > div.layout--kfnv1 > div > div.content--xjOPq > div.file-drop-zone---q-el > div > div.node-list--lS4pi > div.ant-dropdown-trigger.trigger-wrapper--H1k2P > div > div > div > div:nth-child(1) > div > div:nth-child(1) > div > div > div > div > div > div > div > div.node-card--wp9KL > div.cover--5Y7I- > div > img"),
		chromedp.Sleep(time.Second*5),
		chromedp.Click("#root > div.layout--kfnv1 > div > div.content--xjOPq > div.file-drop-zone---q-el > div > div.header-wrapper--sCjdX > div.left-wrapper--N4ml8 > div > div > div > input"),
		chromedp.Sleep(time.Second*3),
		chromedp.Click("#root > div.layout--kfnv1 > div > div.content--xjOPq > div.file-drop-zone---q-el > div > div.outer-wrapper--ANyIY.show--lU1m1.static--mEYfj.enter-done > div > div:nth-child(5) > span > svg > use"),
		chromedp.Click("body > div:nth-child(13) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div.ant-modal-confirm-btns > button.ant-btn.ant-btn-primary.ant-btn-dangerous"),
		chromedp.Click("#root > div.layout--kfnv1 > div > div.nav-tab--IxVGu > div.nav-tab-bottom--XrpFb > div.ant-dropdown-trigger.nav-tab-bottom-more--29ARY"),
		chromedp.Sleep(1*time.Second),
		chromedp.Click("body > div:nth-child(14) > div > div > ul > li:nth-child(13) > div > div"),
		chromedp.Sleep(1*time.Second),
		chromedp.Click("body > div:nth-child(16) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div.ant-modal-confirm-btns > button.ant-btn.ant-btn-primary.ant-btn-dangerous"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(5 * time.Second)

}
